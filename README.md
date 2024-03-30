
## Sponge Examples

[Sponge](https://github.com/zhufuyi/sponge) is a powerful development framework that integrates `code auto generation`, `gin and grpc framework`. It is easy to build a complete project from development to deployment, just fill in the business logic code on the generated template code, implementation of "low-code way" development projects.

Here are some examples of using sponge to develop go projects, the database type used in the example is mysql, also support database mongodb, postgresql, tidb, sqlite.

<br>

### Examples of creating a service using sponge

- [Create **web** service based on **sql** (including CRUD)](https://github.com/zhufuyi/sponge_examples/tree/main/1_web-gin-CRUD)
- [Create **grpc** service based on **sql** (including CRUD)](https://github.com/zhufuyi/sponge_examples/tree/main/2_micro-grpc-CRUD)
- [Create **web** service based on **protobuf**](https://github.com/zhufuyi/sponge_examples/tree/main/3_web-gin-protobuf)
- [Create **grpc** service based on **protobuf** ](https://github.com/zhufuyi/sponge_examples/tree/main/4_micro-grpc-protobuf)
- [Create **grpc gateway** service based on **protobuf**](https://github.com/zhufuyi/sponge_examples/tree/main/5_micro-gin-rpc-gateway)
- [Create **grpc+http** service based on **protobuf**](https://github.com/zhufuyi/sponge_examples/tree/main/a_micro-grpc-http-protobuf)
- [Build a microservice cluster](https://github.com/zhufuyi/sponge_examples/tree/main/6_micro-cluster)

### Examples of developing a complete project using sponge

- [Simple community web backend service](https://github.com/zhufuyi/sponge_examples/tree/main/7_community-single)
- [Simple community web service broken down into microservice](https://github.com/zhufuyi/sponge_examples/tree/main/8_community-cluster)

### Distributed transaction examples

- [Simple distributed order system](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction)
