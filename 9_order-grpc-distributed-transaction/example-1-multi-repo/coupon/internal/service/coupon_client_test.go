// Code generated by https://github.com/zhufuyi/sponge
// Test_service_coupon_methods is used to test the coupon api
// Test_service_coupon_benchmark is used to performance test the coupon api

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"

	couponV1 "coupon/api/coupon/v1"
	"coupon/configs"
	"coupon/internal/config"
)

// Test service coupon api via grpc client
func Test_service_coupon_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := couponV1.NewCouponClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{

		{
			name: "CouponUse",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &couponV1.CouponUseRequest{
					CouponID: 0, 
				}

				return cli.CouponUse(ctx, req)
			},
			wantErr: false,
		},


		{
			name: "CouponUseRevert",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &couponV1.CouponUseRevertRequest{
					CouponID: 0, 
				}

				return cli.CouponUseRevert(ctx, req)
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

// performance test service coupon api, copy the report to
// the browser to view when the pressure test is finished.
func Test_service_coupon_benchmark(t *testing.T) {
	err := config.Init(configs.Path("coupon.yml"))
	if err != nil {
		panic(err)
	}

	grpcClientCfg := getGRPCClientCfg()
	host := fmt.Sprintf("%s:%d", grpcClientCfg.Host, grpcClientCfg.Port)
	protoFile := configs.Path("../api/coupon/v1/coupon.proto")
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
			name: "CouponUse",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &couponV1.CouponUseRequest{
					CouponID: 0, 
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "CouponUse", message, dependentProtoFilePath, total)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},


		{
			name: "CouponUseRevert",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &couponV1.CouponUseRevertRequest{
					CouponID: 0, 
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "CouponUseRevert", message, dependentProtoFilePath, total)
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