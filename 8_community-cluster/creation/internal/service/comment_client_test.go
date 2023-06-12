// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	creationV1 "creation/api/creation/v1"
	"creation/configs"
	"creation/internal/config"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"
)

// Test each method of commentService via the rpc client
func Test_service_commentService_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := creationV1.NewCommentServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{
		{
			name: "Create",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.CreateCommentRequest{
					PostId: 0, 
					UserId: 0, 
					Content: "", 
					DeviceType: "", 
					Ip: "", 
				}
				return cli.Create(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "DeleteByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.DeleteCommentByIDRequest{
					Id: 0, 
					UserId: 0, 
					DelFlag: 0, 
				}
				return cli.DeleteByID(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "UpdateByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.UpdateCommentByIDRequest{
					Id: 0, 
					Content: "", 
				}
				return cli.UpdateByID(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "Reply",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.ReplyCommentRequest{
					CommentId: 0, 
					UserId: 0, 
					Content: "", 
					DeviceType: "", 
					Ip: "", 
				}
				return cli.Reply(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "GetByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.GetCommentByIDRequest{
					Id: 0, 
				}
				return cli.GetByID(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "ListByIDs",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.ListCommentByIDsRequest{
					Ids: nil, 
				}
				return cli.ListByIDs(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "ListLatest",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.ListCommentLatestRequest{
					PostId: 0, 
					Page: 0, 
					Limit: 0, 
				}
				return cli.ListLatest(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "ListHot",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.ListCommentHotRequest{
					PostId: 0, 
					Page: 0, 
					Limit: 0, 
				}
				return cli.ListHot(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "ListReply",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.ListCommentReplyRequest{
					CommentId: 0, 
					Page: 0, 
					Limit: 0, 
				}
				return cli.ListReply(ctx, req)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn()
			if (err != nil) != tt.wantErr {
				t.Errorf("test '%s' error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			data, _ := json.MarshalIndent(got, "", "    ")
			fmt.Println(string(data))
		})
	}
}

// Perform a stress test on commentService's method and 
// copy the press test report to your browser when you are finished.
func Test_service_commentService_benchmark(t *testing.T) {
	err := config.Init(configs.Path("creation.yml"))
	if err != nil {
		panic(err)
	}
	host := fmt.Sprintf("127.0.0.1:%d", config.Get().Grpc.Port)
	protoFile := configs.Path("../api/creation/v1/comment.proto")
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
			name: "Create",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.CreateCommentRequest{
					PostId: 0, 
					UserId: 0, 
					Content: "", 
					DeviceType: "", 
					Ip: "", 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "Create", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
		{
			name: "DeleteByID",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.DeleteCommentByIDRequest{
					Id: 0, 
					UserId: 0, 
					DelFlag: 0, 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "DeleteByID", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
		{
			name: "UpdateByID",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.UpdateCommentByIDRequest{
					Id: 0, 
					Content: "", 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "UpdateByID", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
		{
			name: "Reply",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.ReplyCommentRequest{
					CommentId: 0, 
					UserId: 0, 
					Content: "", 
					DeviceType: "", 
					Ip: "", 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "Reply", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
		{
			name: "GetByID",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.GetCommentByIDRequest{
					Id: 0, 
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
				message := &creationV1.ListCommentByIDsRequest{
					Ids: nil, 
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
			name: "ListLatest",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.ListCommentLatestRequest{
					PostId: 0, 
					Page: 0, 
					Limit: 0, 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "ListLatest", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
		{
			name: "ListHot",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.ListCommentHotRequest{
					PostId: 0, 
					Page: 0, 
					Limit: 0, 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "ListHot", message, total, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
		{
			name: "ListReply",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.ListCommentReplyRequest{
					CommentId: 0, 
					Page: 0, 
					Limit: 0, 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "ListReply", message, total, importPaths...)
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
