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

// Test each method of collectService via the rpc client
func Test_service_collectService_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := creationV1.NewCollectServiceClient(conn)
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
				req := &creationV1.CreateCollectRequest{
					UserId: 0, 
					PostId: 0, 
				}
				return cli.Create(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "Delete",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.DeleteCollectRequest{
					Id: 0, 
					PostId: 0, 
				}
				return cli.Delete(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "List",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &creationV1.ListCollectRequest{
					UserId: 0, 
					Page: 0, 
					Limit: 0, 
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
				t.Errorf("test '%s' error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			data, _ := json.MarshalIndent(got, "", "    ")
			fmt.Println(string(data))
		})
	}
}

// Perform a stress test on collectService's method and 
// copy the press test report to your browser when you are finished.
func Test_service_collectService_benchmark(t *testing.T) {
	err := config.Init(configs.Path("creation.yml"))
	if err != nil {
		panic(err)
	}
	host := fmt.Sprintf("127.0.0.1:%d", config.Get().Grpc.Port)
	protoFile := configs.Path("../api/creation/v1/collect.proto")
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
				message := &creationV1.CreateCollectRequest{
					UserId: 0, 
					PostId: 0, 
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
			name: "Delete",
			fn: func() error {
				// todo type in the parameters to test
				message := &creationV1.DeleteCollectRequest{
					Id: 0, 
					PostId: 0, 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "Delete", message, total, importPaths...)
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
				message := &creationV1.ListCollectRequest{
					UserId: 0, 
					Page: 0, 
					Limit: 0, 
				}
				var total uint = 1000 // total number of requests
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
