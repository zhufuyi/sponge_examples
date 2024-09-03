## English | [简体中文](https://juejin.cn/post/7410062139276263464)

### Environment Setup

- Copy the [user.proto](https://github.com/zhufuyi/sponge_examples/blob/main/_10_micro-grpc-http-protobuf/api/user/v1/user.proto) file.
- Install the [sponge](https://github.com/zhufuyi/sponge/blob/main/assets/install-en.md).

After installing sponge, run the following command to start the UI:

```bash
sponge run
```

<br>

### Quickly Create Microservice Project

Enter the sponge UI, click on the left menu [Protobuf] --> [Create grpc + http service], and fill in the parameters as needed to generate the microservice project code.

![micro-grpc-http-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-grpc-http-pb.png)

UnZip file, navigate to the project root directory (e.g., the user directory) and run the following commands to generate the code:

```bash
# Generate code
make proto

# Open the internal/service/user.go file and add business logic according to the generated sample code.

# Compile and start the service
make run
```

Once the service is running, you can test the APIs for both HTTP and gRPC protocols:

**1. Test gRPC API**

Open the project code using the Goland IDE, navigate to the `internal/service` directory, and locate the file with the `_client_test.go` suffix. This file contains test and performance benchmark functions for each API defined in the proto file. Before testing, ensure to fill in the request parameters.

![micro-grpc-http-pb-test](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/micro-grpc-http-pb-test.png)

If you don't have the Goland IDE, you can also test via the command line. Navigate to the `internal/service` directory, modify the request parameters in the `user_client_test.go` file, and run the following command to test a specific gRPC method, for example: `go test -run Test_service_user_methods/Register`.

**2. Test HTTP API**

Access [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html) through a browser, and use the Swagger UI to test the HTTP API.

![micro-grpc-http-pb-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/micro-grpc-http-pb-swagger.png)

<br>

---

**More detailed development grpc+http service documentation**: [https://go-sponge.com/microservice-development-mix-protobuf](https://go-sponge.com/microservice-development-mix-protobuf)

<br>
