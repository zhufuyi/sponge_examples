// Code generated by https://github.com/zhufuyi/sponge
// Test_service_stock_methods is used to test the stock api
// Test_service_stock_benchmark is used to performance test the stock api

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

// Test service stock api via grpc client
func Test_service_stock_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := stockV1.NewStockClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{

		{
			name: "Create",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &stockV1.CreateStockRequest{
					ProductID: 0, // 商品id
					Stock:     0, // 库存
				}

				return cli.Create(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "DeleteByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &stockV1.DeleteStockByIDRequest{
					Id: 0,
				}

				return cli.DeleteByID(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "UpdateByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &stockV1.UpdateStockByIDRequest{
					Id:        0,
					ProductID: 0, // 商品id
					Stock:     0, // 库存
				}

				return cli.UpdateByID(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "GetByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &stockV1.GetStockByIDRequest{
					Id: 0,
				}

				return cli.GetByID(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "List",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &stockV1.ListStockRequest{
					Params: nil,
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

// performance test service stock api, copy the report to
// the browser to view when the pressure test is finished.
func Test_service_stock_benchmark(t *testing.T) {
	err := config.Init(configs.Path("stock.yml"))
	if err != nil {
		panic(err)
	}

	grpcClientCfg := getGRPCClientCfg()
	host := fmt.Sprintf("%s:%d", grpcClientCfg.Host, grpcClientCfg.Port)
	protoFile := configs.Path("../api/stock/v1/stock.proto")
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
			name: "Create",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &stockV1.CreateStockRequest{
					ProductID: 0, // 商品id
					Stock:     0, // 库存
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "Create", message, dependentProtoFilePath, total)
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
				// todo type in the parameters before benchmark testing
				message := &stockV1.DeleteStockByIDRequest{
					Id: 0,
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "DeleteByID", message, dependentProtoFilePath, total)
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
				// todo type in the parameters before benchmark testing
				message := &stockV1.UpdateStockByIDRequest{
					Id:        0,
					ProductID: 0, // 商品id
					Stock:     0, // 库存
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "UpdateByID", message, dependentProtoFilePath, total)
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
				// todo type in the parameters before benchmark testing
				message := &stockV1.GetStockByIDRequest{
					Id: 0,
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "GetByID", message, dependentProtoFilePath, total)
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
				// todo type in the parameters before benchmark testing
				message := &stockV1.ListStockRequest{
					Params: nil,
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "List", message, dependentProtoFilePath, total)
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