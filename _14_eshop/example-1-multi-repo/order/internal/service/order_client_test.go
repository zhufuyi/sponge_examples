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
				req := &orderV1.SubmitOrderRequest{
					UserID:       0, // 用户id
					ProductID:    0, // 商品id
					ProductCount: 0, // 商品数量
					Amount:       0, // 金额(分)
					CouponID:     0, // 优惠券id
				}

				return cli.Submit(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "SendSubmitOrderNotify",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &orderV1.SendSubmitOrderNotifyRequest{
					UserID:       0, // 用户id
					ProductID:    0, // 商品id
					ProductCount: 0, // 商品数量
					Amount:       0, // 金额(分)
					CouponID:     0, // 优惠券id
				}

				return cli.SendSubmitOrderNotify(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "Create",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &orderV1.CreateOrderRequest{
					OrderID:      "", // 订单id
					UserID:       0,  // 用户id
					ProductID:    0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount:       0,  // 金额(分)
					CouponID:     0,  // 优惠券id
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
					OrderID:      "", // 订单id
					UserID:       0,  // 用户id
					ProductID:    0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount:       0,  // 金额(分)
					CouponID:     0,  // 优惠券id
				}

				return cli.CreateRevert(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "DeleteByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &orderV1.DeleteOrderByIDRequest{
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
				req := &orderV1.UpdateOrderByIDRequest{
					Id:           0,
					OrderID:      "", // 订单id
					UserID:       0,  // 用户id
					ProductID:    0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount:       0,  // 金额(分)
					CouponID:     0,  // 优惠券id
					Status:       0,  // 订单状态: 1:待支付, 2:待发货, 3:待收货, 4:已完成, 5:已取消
				}

				return cli.UpdateByID(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "GetByID",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &orderV1.GetOrderByIDRequest{
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
				req := &orderV1.ListOrderRequest{
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
				message := &orderV1.SubmitOrderRequest{
					UserID:       0, // 用户id
					ProductID:    0, // 商品id
					ProductCount: 0, // 商品数量
					Amount:       0, // 金额(分)
					CouponID:     0, // 优惠券id
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
			name: "SendSubmitOrderNotify",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &orderV1.SendSubmitOrderNotifyRequest{
					UserID:       0, // 用户id
					ProductID:    0, // 商品id
					ProductCount: 0, // 商品数量
					Amount:       0, // 金额(分)
					CouponID:     0, // 优惠券id
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "SendSubmitOrderNotify", message, dependentProtoFilePath, total)
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
					OrderID:      "", // 订单id
					UserID:       0,  // 用户id
					ProductID:    0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount:       0,  // 金额(分)
					CouponID:     0,  // 优惠券id
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
					OrderID:      "", // 订单id
					UserID:       0,  // 用户id
					ProductID:    0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount:       0,  // 金额(分)
					CouponID:     0,  // 优惠券id
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

		{
			name: "DeleteByID",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &orderV1.DeleteOrderByIDRequest{
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
				message := &orderV1.UpdateOrderByIDRequest{
					Id:           0,
					OrderID:      "", // 订单id
					UserID:       0,  // 用户id
					ProductID:    0,  // 商品id
					ProductCount: 0,  // 商品数量
					Amount:       0,  // 金额(分)
					CouponID:     0,  // 优惠券id
					Status:       0,  // 订单状态: 1:待支付, 2:待发货, 3:待收货, 4:已完成, 5:已取消
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
				message := &orderV1.GetOrderByIDRequest{
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
				message := &orderV1.ListOrderRequest{
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
