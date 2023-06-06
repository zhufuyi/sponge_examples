package handler

import (
	"net/http"
	"testing"
	"time"

	"user/internal/cache"
	"user/internal/dao"
	"user/internal/model"
	"user/internal/types"

	"github.com/zhufuyi/sponge/pkg/gohttp"
	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/mysql/query"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

func newCourseHandler() *gotest.Handler {
	// todo additional test field information
	testData := &model.Course{}
	testData.ID = 1
	testData.CreatedAt = time.Now()
	testData.UpdatedAt = testData.CreatedAt

	// init mock cache
	c := gotest.NewCache(map[string]interface{}{utils.Uint64ToStr(testData.ID): testData})
	c.ICache = cache.NewCourseCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})

	// init mock dao
	d := gotest.NewDao(c, testData)
	d.IDao = dao.NewCourseDao(d.DB, c.ICache.(cache.CourseCache))

	// init mock handler
	h := gotest.NewHandler(d, testData)
	h.IHandler = &courseHandler{iDao: d.IDao.(dao.CourseDao)}

	testFns := []gotest.RouterInfo{
		{
			FuncName:    "Create",
			Method:      http.MethodPost,
			Path:        "/course",
			HandlerFunc: h.IHandler.(CourseHandler).Create,
		},
		{
			FuncName:    "DeleteByID",
			Method:      http.MethodDelete,
			Path:        "/course/:id",
			HandlerFunc: h.IHandler.(CourseHandler).DeleteByID,
		},
		{
			FuncName:    "DeleteByIDs",
			Method:      http.MethodPost,
			Path:        "/course/delete/ids",
			HandlerFunc: h.IHandler.(CourseHandler).DeleteByIDs,
		},
		{
			FuncName:    "UpdateByID",
			Method:      http.MethodPut,
			Path:        "/course/:id",
			HandlerFunc: h.IHandler.(CourseHandler).UpdateByID,
		},
		{
			FuncName:    "GetByID",
			Method:      http.MethodGet,
			Path:        "/course/:id",
			HandlerFunc: h.IHandler.(CourseHandler).GetByID,
		},
		{
			FuncName:    "ListByIDs",
			Method:      http.MethodPost,
			Path:        "/course/list/ids",
			HandlerFunc: h.IHandler.(CourseHandler).ListByIDs,
		},
		{
			FuncName:    "List",
			Method:      http.MethodPost,
			Path:        "/course/list",
			HandlerFunc: h.IHandler.(CourseHandler).List,
		},
	}

	h.GoRunHTTPServer(testFns)

	time.Sleep(time.Millisecond * 200)
	return h
}

func Test_courseHandler_Create(t *testing.T) {
	h := newCourseHandler()
	defer h.Close()
	testData := &types.CreateCourseRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.Course))

	h.MockDao.SQLMock.ExpectBegin()
	args := h.MockDao.GetAnyArgs(h.TestData)
	h.MockDao.SQLMock.ExpectExec("INSERT INTO .*").
		WithArgs(args[:len(args)-1]...). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(1, 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("Create"), testData)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)

}

func Test_courseHandler_DeleteByID(t *testing.T) {
	h := newCourseHandler()
	defer h.Close()
	testData := h.TestData.(*model.Course)

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Delete(result, h.GetRequestURL("DeleteByID", testData.ID))
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Delete(result, h.GetRequestURL("DeleteByID", 0))
	assert.NoError(t, err)

	// delete error test
	err = gohttp.Delete(result, h.GetRequestURL("DeleteByID", 111))
	assert.Error(t, err)
}

func Test_courseHandler_DeleteByIDs(t *testing.T) {
	h := newCourseHandler()
	defer h.Close()
	testData := h.TestData.(*model.Course)

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("DeleteByIDs"), &types.DeleteCoursesByIDsRequest{IDs: []uint64{testData.ID}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Post(result, h.GetRequestURL("DeleteByIDs"), nil)
	assert.NoError(t, err)

	// get error test
	err = gohttp.Post(result, h.GetRequestURL("DeleteByIDs"), &types.DeleteCoursesByIDsRequest{IDs: []uint64{111}})
	assert.Error(t, err)
}

func Test_courseHandler_UpdateByID(t *testing.T) {
	h := newCourseHandler()
	defer h.Close()
	testData := &types.UpdateCourseByIDRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.Course))

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Put(result, h.GetRequestURL("UpdateByID", testData.ID), testData)
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Put(result, h.GetRequestURL("UpdateByID", 0), testData)
	assert.NoError(t, err)

	// update error test
	err = gohttp.Put(result, h.GetRequestURL("UpdateByID", 111), testData)
	assert.Error(t, err)
}

func Test_courseHandler_GetByID(t *testing.T) {
	h := newCourseHandler()
	defer h.Close()
	testData := h.TestData.(*model.Course)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").
		WithArgs(testData.ID).
		WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Get(result, h.GetRequestURL("GetByID", testData.ID))
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Get(result, h.GetRequestURL("GetByID", 0))
	assert.NoError(t, err)

	// get error test
	err = gohttp.Get(result, h.GetRequestURL("GetByID", 111))
	assert.Error(t, err)
}

func Test_courseHandler_ListByIDs(t *testing.T) {
	h := newCourseHandler()
	defer h.Close()
	testData := h.TestData.(*model.Course)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("ListByIDs"), &types.GetCoursesByIDsRequest{IDs: []uint64{testData.ID}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Post(result, h.GetRequestURL("ListByIDs"), nil)
	assert.NoError(t, err)

	// get error test
	err = gohttp.Post(result, h.GetRequestURL("ListByIDs"), &types.GetCoursesByIDsRequest{IDs: []uint64{111}})
	assert.Error(t, err)
}

func Test_courseHandler_List(t *testing.T) {
	h := newCourseHandler()
	defer h.Close()
	testData := h.TestData.(*model.Course)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("List"), &types.GetCoursesRequest{query.Params{
		Page: 0,
		Size: 10,
		Sort: "ignore count", // ignore test count
	}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// nil params error test
	err = gohttp.Post(result, h.GetRequestURL("List"), nil)

	// get error test
	err = gohttp.Post(result, h.GetRequestURL("List"), &types.GetCoursesRequest{query.Params{
		Page: 0,
		Size: 10,
	}})
}

func TestNewCourseHandler(t *testing.T) {
	defer func() {
		recover()
	}()
	_ = NewCourseHandler()
}
