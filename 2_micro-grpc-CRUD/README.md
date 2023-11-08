
[**micro-grpc-CRUD 中文说明**](https://juejin.cn/post/7225257817346129981)

<br>

**Import [sql](https://github.com/zhufuyi/sponge_examples/blob/main/3_micro-grpc-CRUD/test/sql/user.sql) into mysql before you start.**

<br>

### Quickly create a microservice project

Enter the Sponge UI interface, click on the left menu bar 【SQL】→【Create microservice project】, fill in some parameters to generate a complete microservice project code.

The microservice code is mainly composed of commonly used libraries such as [grpc](https://github.com/grpc/grpc-go), [gorm](https://github.com/go-gorm/gorm), [go-redis](https://github.com/go-redis/redis), and also includes grpc client CRUD test code, common service governance function code, build deployment scripts, etc.

![micro-rpc](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc.png)

Switch to the user directory and run the command:

```bash
# Generate pb.go code
make proto

# Compile and start grpc service
make run
```

Use goland IDE to open user service code, enter the internal/service directory, open the `teacher_client_test.go` file, you can test CRUD methods here, similar to testing CRUD interfaces in swagger interface. Fill in parameters before testing and click the green button to test.

![micro-grpc-test](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/micro-rpc-test.png)

<br>

### Batch add CRUD code to grpc service

Enter the Sponge UI interface, click on the left menu bar 【Public】→【Generate service CRUD code】, select any number of tables to generate code, then move the generated CRUD code to the grpc service directory to complete batch addition of CURD interfaces in microservices without changing any code.

![micro-grpc-service](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-service.png)

Switch to user service directory and run command:

```bash
# Generate pb.go code
make proto

# Compile and start user service
make run
```

Use goland IDE, enter internal/service directory, open `teach_client_test.go` and `course_client_test.go` files to test CRUD methods.

<br>

**More detailed development microservice documentation** https://go-sponge.com/microservice-development-mysql

<br>