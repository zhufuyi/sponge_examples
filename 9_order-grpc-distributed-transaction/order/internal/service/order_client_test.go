// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	orderV1 "order/api/order/v1"
	"order/configs"
	"order/internal/config"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"
)

// Test each method of order via the rpc client
func Test_service_order_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := orderV1.NewOrderClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{
		{
			name: "Submit",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &orderV1.SubmitRequest{
					UserId:       1,    // 用户id
					ProductId:    1,    // 商品id
					Amount:       1100, // 金额(分)
					ProductCount: 1,    // 商品数量
					CouponId:     1,    // 优惠券id
				}
				return cli.Submit(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "Create",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &orderV1.CreateOrderRequest{
					OrderId:      "", // 订单id
					UserId:       0,  // 用户id
					ProductId:    0,  // 商品id
					Amount:       0,  // 金额(分)
					ProductCount: 0,  // 商品数量
					CouponId:     0,  // 优惠券id
				}
				return cli.Create(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "CreateRevert",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &orderV1.CreateOrderRevertRequest{
					OrderId:      "", // 订单id
					UserId:       0,  // 用户id
					ProductId:    0,  // 商品id
					Amount:       0,  // 金额(分)
					ProductCount: 0,  // 商品数量
					CouponId:     0,  // 优惠券id
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

// Perform a stress test on order's method and
// copy the press test report to your browser when you are finished.
func Test_service_order_benchmark(t *testing.T) {
	err := config.Init(configs.Path("order.yml"))
	if err != nil {
		panic(err)
	}
	host := fmt.Sprintf("127.0.0.1:%d", config.Get().Grpc.Port)
	protoFile := configs.Path("../api/order/v1/order.proto")
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
			name: "Submit",
			fn: func() error {
				// todo type in the parameters to test
				message := &orderV1.SubmitRequest{
					UserId:       0, // 用户id
					ProductId:    0, // 商品id
					Amount:       0, // 金额(分)
					ProductCount: 0, // 商品数量
					CouponId:     0, // 优惠券id
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "Submit", message, total, importPaths...)
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
				// todo type in the parameters to test
				message := &orderV1.CreateOrderRequest{
					OrderId:      "", // 订单id
					UserId:       0,  // 用户id
					ProductId:    0,  // 商品id
					Amount:       0,  // 金额(分)
					ProductCount: 0,  // 商品数量
					CouponId:     0,  // 优惠券id
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
			name: "CreateRevert",
			fn: func() error {
				// todo type in the parameters to test
				message := &orderV1.CreateOrderRevertRequest{
					OrderId:      "", // 订单id
					UserId:       0,  // 用户id
					ProductId:    0,  // 商品id
					Amount:       0,  // 金额(分)
					ProductCount: 0,  // 商品数量
					CouponId:     0,  // 优惠券id
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "CreateRevert", message, total, importPaths...)
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
