package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"user/api/types"
	userV1 "user/api/user/v1"
	"user/configs"
	"user/internal/config"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"
)

// Test each method of course via the rpc client
func Test_service_course_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := userV1.NewCourseServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{

		{
			name: "Create",
			fn: func() (interface{}, error) {
				// todo enter parameters before testing
				req := &userV1.CreateCourseRequest{
					Code:     "", // 课程代号
					Name:     "", // 课程名称
					Credit:   0,  // 学分
					College:  "", // 学院
					Semester: "", // 学期
					Time:     "", // 上课时间
					Place:    "", // 上课地点
				}
				return cli.Create(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "UpdateByID",
			fn: func() (interface{}, error) {
				// todo enter parameters before testing
				req := &userV1.UpdateCourseByIDRequest{
					Id:       0,
					Code:     "", // 课程代号
					Name:     "", // 课程名称
					Credit:   0,  // 学分
					College:  "", // 学院
					Semester: "", // 学期
					Time:     "", // 上课时间
					Place:    "", // 上课地点
				}
				return cli.UpdateByID(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "DeleteByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &userV1.DeleteCourseByIDRequest{
					Id: 100,
				}
				return cli.DeleteByID(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "DeleteByIDs",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &userV1.DeleteCourseByIDsRequest{
					Ids: []uint64{100},
				}
				return cli.DeleteByIDs(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "GetByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &userV1.GetCourseByIDRequest{
					Id: 1,
				}
				return cli.GetByID(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "ListByIDs",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &userV1.ListCourseByIDsRequest{
					Ids: []uint64{1, 2, 3},
				}
				return cli.ListByIDs(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "List",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &userV1.ListCourseRequest{
					Params: &types.Params{
						Page:  0,
						Limit: 10,
						Sort:  "",
						Columns: []*types.Column{
							{
								Name:  "id",
								Exp:   ">=",
								Value: "1",
								Logic: "",
							},
						},
					},
				}
				return cli.List(ctx, req)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn()
			if (err != nil) != tt.wantErr {
				t.Logf("test '%s' error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			data, _ := json.MarshalIndent(got, "", "    ")
			fmt.Println(string(data))
		})
	}
}

// Perform a stress test on {{.LowerName}}'s method and
// copy the press test report to your browser when you are finished.
func Test_service_course_benchmark(t *testing.T) {
	err := config.Init(configs.Path("user.yml"))
	if err != nil {
		panic(err)
	}
	host := fmt.Sprintf("127.0.0.1:%d", config.Get().Grpc.Port)
	protoFile := configs.Path("../api/user/v1/course.proto")
	// If third-party dependencies are missing during the press test,
	// copy them to the project's third_party directory.
	importPaths := []string{
		configs.Path("../third_party"), // third_party directory
		configs.Path(".."),             // Previous level of third_party
	}

	tests := []struct {
		name    string
		fn      func() error
		wantErr bool
	}{
		{
			name: "GetByID",
			fn: func() error {
				// todo type in the parameters to test
				message := &userV1.GetCourseByIDRequest{
					Id: 1,
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "GetByID", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},

		{
			name: "ListByIDs",
			fn: func() error {
				// todo type in the parameters to test
				message := &userV1.ListCourseByIDsRequest{
					Ids: []uint64{1, 2, 3},
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "ListByIDs", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},

		{
			name: "List",
			fn: func() error {
				// todo type in the parameters to test
				message := &userV1.ListCourseRequest{
					Params: &types.Params{
						Page:  0,
						Limit: 10,
						Sort:  "",
						Columns: []*types.Column{
							{
								Name:  "id",
								Exp:   ">=",
								Value: "1",
								Logic: "",
							},
						},
					},
				}
				var total uint = 100 // total number of requests
				b, err := benchmark.New(host, protoFile, "List", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			if (err != nil) != tt.wantErr {
				t.Errorf("test '%s' error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
		})
	}
}
