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

func newTeachHandler() *gotest.Handler {
	// todo additional test field information
	testData := &model.Teach{}
	testData.ID = 1
	testData.CreatedAt = time.Now()
	testData.UpdatedAt = testData.CreatedAt

	// init mock cache
	c := gotest.NewCache(map[string]interface{}{utils.Uint64ToStr(testData.ID): testData})
	c.ICache = cache.NewTeachCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})

	// init mock dao
	d := gotest.NewDao(c, testData)
	d.IDao = dao.NewTeachDao(d.DB, c.ICache.(cache.TeachCache))

	// init mock handler
	h := gotest.NewHandler(d, testData)
	h.IHandler = &teachHandler{iDao: d.IDao.(dao.TeachDao)}

	testFns := []gotest.RouterInfo{
		{
			FuncName:    "Create",
			Method:      http.MethodPost,
			Path:        "/teach",
			HandlerFunc: h.IHandler.(TeachHandler).Create,
		},
		{
			FuncName:    "DeleteByID",
			Method:      http.MethodDelete,
			Path:        "/teach/:id",
			HandlerFunc: h.IHandler.(TeachHandler).DeleteByID,
		},
		{
			FuncName:    "DeleteByIDs",
			Method:      http.MethodPost,
			Path:        "/teach/delete/ids",
			HandlerFunc: h.IHandler.(TeachHandler).DeleteByIDs,
		},
		{
			FuncName:    "UpdateByID",
			Method:      http.MethodPut,
			Path:        "/teach/:id",
			HandlerFunc: h.IHandler.(TeachHandler).UpdateByID,
		},
		{
			FuncName:    "GetByID",
			Method:      http.MethodGet,
			Path:        "/teach/:id",
			HandlerFunc: h.IHandler.(TeachHandler).GetByID,
		},
		{
			FuncName:    "ListByIDs",
			Method:      http.MethodPost,
			Path:        "/teach/list/ids",
			HandlerFunc: h.IHandler.(TeachHandler).ListByIDs,
		},
		{
			FuncName:    "List",
			Method:      http.MethodPost,
			Path:        "/teach/list",
			HandlerFunc: h.IHandler.(TeachHandler).List,
		},
	}

	h.GoRunHTTPServer(testFns)

	time.Sleep(time.Millisecond * 200)
	return h
}

func Test_teachHandler_Create(t *testing.T) {
	h := newTeachHandler()
	defer h.Close()
	testData := &types.CreateTeachRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.Teach))

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

func Test_teachHandler_DeleteByID(t *testing.T) {
	h := newTeachHandler()
	defer h.Close()
	testData := h.TestData.(*model.Teach)

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

func Test_teachHandler_DeleteByIDs(t *testing.T) {
	h := newTeachHandler()
	defer h.Close()
	testData := h.TestData.(*model.Teach)

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("DeleteByIDs"), &types.DeleteTeachsByIDsRequest{IDs: []uint64{testData.ID}})
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
	err = gohttp.Post(result, h.GetRequestURL("DeleteByIDs"), &types.DeleteTeachsByIDsRequest{IDs: []uint64{111}})
	assert.Error(t, err)
}

func Test_teachHandler_UpdateByID(t *testing.T) {
	h := newTeachHandler()
	defer h.Close()
	testData := &types.UpdateTeachByIDRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.Teach))

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

func Test_teachHandler_GetByID(t *testing.T) {
	h := newTeachHandler()
	defer h.Close()
	testData := h.TestData.(*model.Teach)

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

func Test_teachHandler_ListByIDs(t *testing.T) {
	h := newTeachHandler()
	defer h.Close()
	testData := h.TestData.(*model.Teach)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("ListByIDs"), &types.GetTeachsByIDsRequest{IDs: []uint64{testData.ID}})
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
	err = gohttp.Post(result, h.GetRequestURL("ListByIDs"), &types.GetTeachsByIDsRequest{IDs: []uint64{111}})
	assert.Error(t, err)
}

func Test_teachHandler_List(t *testing.T) {
	h := newTeachHandler()
	defer h.Close()
	testData := h.TestData.(*model.Teach)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("List"), &types.GetTeachsRequest{query.Params{
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
	err = gohttp.Post(result, h.GetRequestURL("List"), &types.GetTeachsRequest{query.Params{
		Page: 0,
		Size: 10,
	}})
}

func TestNewTeachHandler(t *testing.T) {
	defer func() {
		recover()
	}()
	_ = NewTeachHandler()
}
