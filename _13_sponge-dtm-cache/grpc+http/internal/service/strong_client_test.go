// Code generated by https://github.com/zhufuyi/sponge
// Test_service_strong_methods is used to test the strong api
// Test_service_strong_benchmark is used to performance test the strong api

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"

	stockV1 "stock/api/stock/v1"
	"stock/configs"
	"stock/internal/config"
)

// Test service strong api via grpc client
func Test_service_strong_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := stockV1.NewStrongClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{

		{
			name: "Update",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &stockV1.UpdateStrongRequest{
					Id:    0,
					Stock: 0, // 库存数量
				}

				return cli.Update(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "Query",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &stockV1.QueryStrongRequest{
					Id: 0,
				}

				return cli.Query(ctx, req)
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

// performance test service strong api, copy the report to
// the browser to view when the pressure test is finished.
func Test_service_strong_benchmark(t *testing.T) {
	err := config.Init(configs.Path("stock.yml"))
	if err != nil {
		panic(err)
	}

	grpcClientCfg := getGRPCClientCfg()
	host := fmt.Sprintf("%s:%d", grpcClientCfg.Host, grpcClientCfg.Port)
	protoFile := configs.Path("../api/stock/v1/strong.proto")
	// If third-party dependencies are missing during the press test,
	// copy them to the project's third_party directory.
	dependentProtoFilePath := []string{
		configs.Path("../third_party"), // third_party directory
		configs.Path(".."),             // Previous level of third_party
	}

	tests := []struct {
		name    string
		fn      func() error
		wantErr bool
	}{

		{
			name: "Update",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &stockV1.UpdateStrongRequest{
					Id:    0,
					Stock: 0, // 库存数量
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "Update", message, dependentProtoFilePath, total)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},

		{
			name: "Query",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &stockV1.QueryStrongRequest{
					Id: 0,
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "Query", message, dependentProtoFilePath, total)
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