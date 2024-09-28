// Code generated by https://github.com/zhufuyi/sponge
// Test_service_order_methods is used to test the order api
// Test_service_order_benchmark is used to performance test the order api

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"

	orderV1 "order/api/order/v1"
	"order/configs"
	"order/internal/config"
)

// Test service order api via grpc client
func Test_service_order_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := orderV1.NewOrderClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{

		{
			name: "Submit",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &orderV1.SubmitRequest{
					UserID: 0,  // 用户id
					ProductID: 0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount: 0,  // 金额(分)
					CouponID: 0,  // 优惠券id
				}

				return cli.Submit(ctx, req)
			},
			wantErr: false,
		},


		{
			name: "Create",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &orderV1.CreateOrderRequest{
					OrderID: "",  // 订单id
					UserID: 0,  // 用户id
					ProductID: 0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount: 0,  // 金额(分)
					CouponID: 0,  // 优惠券id
				}

				return cli.Create(ctx, req)
			},
			wantErr: false,
		},


		{
			name: "CreateRevert",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &orderV1.CreateOrderRevertRequest{
					OrderID: "",  // 订单id
					UserID: 0,  // 用户id
					ProductID: 0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount: 0,  // 金额(分)
					CouponID: 0,  // 优惠券id
				}

				return cli.CreateRevert(ctx, req)
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

// performance test service order api, copy the report to
// the browser to view when the pressure test is finished.
func Test_service_order_benchmark(t *testing.T) {
	err := config.Init(configs.Path("order.yml"))
	if err != nil {
		panic(err)
	}

	grpcClientCfg := getGRPCClientCfg()
	host := fmt.Sprintf("%s:%d", grpcClientCfg.Host, grpcClientCfg.Port)
	protoFile := configs.Path("../api/order/v1/order.proto")
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
			name: "Submit",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &orderV1.SubmitRequest{
					UserID: 0,  // 用户id
					ProductID: 0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount: 0,  // 金额(分)
					CouponID: 0,  // 优惠券id
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "Submit", message, dependentProtoFilePath, total)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},


		{
			name: "Create",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &orderV1.CreateOrderRequest{
					OrderID: "",  // 订单id
					UserID: 0,  // 用户id
					ProductID: 0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount: 0,  // 金额(分)
					CouponID: 0,  // 优惠券id
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
			name: "CreateRevert",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &orderV1.CreateOrderRevertRequest{
					OrderID: "",  // 订单id
					UserID: 0,  // 用户id
					ProductID: 0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount: 0,  // 金额(分)
					CouponID: 0,  // 优惠券id
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "CreateRevert", message, dependentProtoFilePath, total)
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
