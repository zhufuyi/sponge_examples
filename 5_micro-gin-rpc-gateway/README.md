[**micro-gin-rpc-gateway 中文说明**](https://www.bilibili.com/read/cv23189890)

<br>

### Quickly Create a GRPC Gateway Project

Before creating a grpc gateway project, prepare a proto file that it must contain **routing description information** and **Swagger description information**.

Enter the Sponge UI interface, click the left menu bar 【Protobuf】 → 【create grpc gateway project】, fill in some parameters to generate the grpc gateway project code.

The web framework uses [gin](https://github.com/gin-gonic/gin), which also includes Swagger documentation, common service governance code, build deployment scripts, etc.

![micro-rpc-gw-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-gw-pb.png)

<br>

### Calling Other Microservice API Interfaces

In order to connect the grpc service in the grpc gateway service, you need to generate code to connect the grpc service separately. Click the left menu bar 【Public】 → 【generate grpc connection code】, fill in some parameters to generate the code, and then move the generated code to the grpc gateway project code.

![micro-rpc-cli](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-cli.png)

<br>

### Configure the GRPC Service Address

Open the configuration file `configs/edusys_gw.yml` and set the address of the grpc service.

```yaml
grpcClient:
  - name: "user"   # rpc service name, used for service discovery
    host: "127.0.0.1"                    # rpc service address, used for direct connection
    port: 8282                              # rpc service port
    registryDiscoveryType: ""         # registration and discovery types: consul, etcd, nacos, if empty, connecting to server using host and port
    enableLoadBalance: false         # whether to turn on the load balancer
```

<br>

### Copy Proto Files From GRPC Service

Open a terminal in the grpc gateway service directory and execute the command to copy the proto file for the grpc service.

```bash
# ../../4_micro-grpc-protobuf is the grpc service directory.
make copy-proto SERVER=../../4_micro-grpc-protobuf
```

<br>

### Running GRPC Service and GRPC Gateway Service

**Running GRPC Service**

Switch to the directory `4_micro-grpc-protobuf`, and start the user service.

```bash
make run
```

**Running GRPC Gateway Service**

Switch to the directory `5_micro-gin-rpc-gateway/edusys-grpc-gateway`.

```bash
# Generate code
make proto  

# Open internal/service/edusys_gw.go, this is the generated API interface code, 
# with a line of panic code prompting you to fill in business logic code. 
# First import the generated code to connect the grpc server, and call the grpc service method in the business logic.

# Compile and start the web service
make run
```

Open [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html) in your browser to test the API interface.

![micro-rpc-gw-pb-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/micro-rpc-gw-pb-swagger.png)

<br>

**More detailed documentation for developing grpc gateways** https://go-sponge.com/rpc-gateway-development-protobuf

<br>
