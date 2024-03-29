// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	couponV1 "coupon/api/coupon/v1"
	"coupon/configs"
	"coupon/internal/config"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"
)

// Test each method of coupon via the rpc client
func Test_service_coupon_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := couponV1.NewCouponClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{
		{
			name: "CouponUse",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &couponV1.CouponUseRequest{
					CouponId: 0, 
				}
				return cli.CouponUse(ctx, req)
			},
			wantErr: false,
		},
		{
			name: "CouponUseRevert",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &couponV1.CouponUseRevertRequest{
					CouponId: 0, 
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

// Perform a stress test on coupon's method and 
// copy the press test report to your browser when you are finished.
func Test_service_coupon_benchmark(t *testing.T) {
	err := config.Init(configs.Path("coupon.yml"))
	if err != nil {
		panic(err)
	}
	host := fmt.Sprintf("127.0.0.1:%d", config.Get().Grpc.Port)
	protoFile := configs.Path("../api/coupon/v1/coupon.proto")
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
			name: "CouponUse",
			fn: func() error {
				// todo type in the parameters to test
				message := &couponV1.CouponUseRequest{
					CouponId: 0, 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "CouponUse", message, total, importPaths...)
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
				// todo type in the parameters to test
				message := &couponV1.CouponUseRevertRequest{
					CouponId: 0, 
				}
				var total uint = 1000 // total number of requests
				b, err := benchmark.New(host, protoFile, "CouponUseRevert", message, total, importPaths...)
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
