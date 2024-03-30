// Code generated by https://github.com/zhufuyi/sponge
// Test_service_user_methods is used to test the user api
// Test_service_user_benchmark is used to performance test the user api

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"

	userV1 "user/api/user/v1"
	"user/configs"
	"user/internal/config"
)

// Test service user api via grpc client
func Test_service_user_methods(t *testing.T) {
	conn := getRPCClientConnForTest()
	cli := userV1.NewUserClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{

		{
			name: "Register",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &userV1.RegisterRequest{
					Email:    "",
					Password: "",
				}

				return cli.Register(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "Login",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &userV1.LoginRequest{
					Email:    "",
					Password: "",
				}

				return cli.Login(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "Logout",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &userV1.LogoutRequest{
					Id:    0,
					Token: "",
				}

				return cli.Logout(ctx, req)
			},
			wantErr: false,
		},

		{
			name: "ChangePassword",
			fn: func() (interface{}, error) {
				// todo type in the parameters before testing
				req := &userV1.ChangePasswordRequest{
					Id:       0,
					Password: "",
				}

				return cli.ChangePassword(ctx, req)
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

// performance test service user api, copy the report to
// the browser to view when the pressure test is finished.
func Test_service_user_benchmark(t *testing.T) {
	err := config.Init(configs.Path("user.yml"))
	if err != nil {
		panic(err)
	}
	host := fmt.Sprintf("127.0.0.1:%d", config.Get().Grpc.Port)
	protoFile := configs.Path("../api/user/v1/user.proto")
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
			name: "Register",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &userV1.RegisterRequest{
					Email:    "",
					Password: "",
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "Register", message, dependentProtoFilePath, total)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},

		{
			name: "Login",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &userV1.LoginRequest{
					Email:    "",
					Password: "",
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "Login", message, dependentProtoFilePath, total)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},

		{
			name: "Logout",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &userV1.LogoutRequest{
					Id:    0,
					Token: "",
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "Logout", message, dependentProtoFilePath, total)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},

		{
			name: "ChangePassword",
			fn: func() error {
				// todo type in the parameters before benchmark testing
				message := &userV1.ChangePasswordRequest{
					Id:       0,
					Password: "",
				}
				total := 1000 // total number of requests

				b, err := benchmark.New(host, protoFile, "ChangePassword", message, dependentProtoFilePath, total)
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
