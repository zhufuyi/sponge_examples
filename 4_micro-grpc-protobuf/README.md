
[**micro-grpc-protobuf 中文说明**](https://juejin.cn/post/7226361209803685948)

<br>

### Quickly create a microservice project

Prepare a proto file before creating a microservice, Enter the sponge UI interface, click on 【Protobuf】→ 【create microservice project】in the left menu bar, fill in some parameters to generate common microservice project code.

The microservice framework uses [grpc](https://github.com/grpc/grpc-go), and also includes commonly used service governance function code, build deployment scripts, etc. The database used is chosen by yourself.

![micro-rpc-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-pb.png)

Switch to the user directory and run the command:

```bash
# Generate pb.go code, generate template code, generate test code
make proto

# Open internal/service/user.go, this is the generated template code. 
# There is a panic code that prompts you to fill in business logic code. Fill in business logic here.

# Compile and start user service
make run
```

Open user service code with goland IDE, go to internal/service directory, open `user_client_test.go` file, you can test grpc method here, similar to testing interface on swagger interface. Fill in parameters before testing and click green button to test.

![micro-rpc-pb-test](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/micro-rpc-pb-test.png)

<br>

**More detailed development microservice documentation** https://go-sponge.com/microservice-development-protobuf

<br>
