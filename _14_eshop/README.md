## English | [简体中文](readme-cn.md)

The microservice example built quickly by Sponge mainly realizes the flash sale and order functions, and ensures data consistency through the distributed transaction manager DTM. The system architecture is shown in the following figure:

![flash-sale-order-cache](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/flash-sale-order-cache.png)

<br>

This includes [two source code examples](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop). The two examples are identical except for slight differences in code structure. The main purpose is to demonstrate that Sponge supports creating microservice projects with different repository patterns. `example-1-multi-repo` is suitable for microservices with multiple repositories, while `example-2-mono-repo` is suitable for microservices with a single repository.

<br>

### Environment Setup

To build this e-commerce system, you need the following tools and dependencies:

- **[Sponge](https://github.com/zhufuyi/sponge)**: A tool for rapidly generating service and module code for different systems. Follow the [Sponge installation guide](https://github.com/zhufuyi/sponge/blob/main/assets/install-en.md) for more information.

Additionally, the following services are required:

- **[DTM](https://github.com/dtm-labs/dtm)**: A distributed transaction manager for ensuring data consistency across multiple service calls.
- **Redis**: Used in conjunction with DTM to handle flash sale caching.
- **Kafka**: A message queue system for asynchronous order processing.
- **MySQL**: Database for storing service data.

All services will be running on a virtual machine with the IP address `192.168.3.37`.

<br>

### Launching DTM Services

[DTM](https://github.com/dtm-labs/dtm) is a core component of this system, responsible for managing distributed transactions in the flash sale and order processes. Two DTM service instances are required—one for MySQL and one for Redis storage.

| Service Name   | Port Configuration            |
|----------------|-------------------------------|
| DTM-MySQL      | HTTP: 36789, gRPC: 36790      |
| DTM-Redis      | HTTP: 35789, gRPC: 35790      |

<br>

#### 1. Launching DTM-MySQL Service

- Import the required table structure into MySQL:
  - [dtmcli.barrier.mysql.sql](https://github.com/dtm-labs/dtm/blob/main/sqls/dtmcli.barrier.mysql.sql)
  - [dtmsvr.storage.mysql.sql](https://github.com/dtm-labs/dtm/blob/main/sqls/dtmsvr.storage.mysql.sql)

- Modify the DTM configuration file ([Sample Configuration](https://github.com/dtm-labs/dtm/blob/main/conf.sample.yml)):

   ```yaml
   Store: # specify which engine to store trans status
     Driver: 'mysql'
     Host: '192.168.3.37'
     User: 'root'
     Password: '123456'
     Port: 3306
     Db: 'dtm'
   ```

- Start the DTM service:

   ```bash
   ./dtm -c conf.yml
   ```

<br>

#### 2. Launching DTM-Redis Service

- Modify the DTM configuration file ([Sample Configuration](https://github.com/dtm-labs/dtm/blob/main/conf.sample.yml)):

   ```yaml
   Store: # specify which engine to store trans status
       Driver: 'redis'
       Host: '192.168.3.37'
       User: 'default'
       Password: '123456'
       Port: 6379
   ```

- Start the DTM service:

   ```bash
   ./dtm -c conf.yml
   ```

<br>

### Rapid Development of an E-Commerce System Using Sponge

The simplified e-commerce system consists of the following eight microservices:

- **eshop_gw**: API gateway service
- **user**: User service
- **product**: Product service
- **order**: Order service
- **stock**: Stock service
- **coupon**: Coupon service
- **pay**: Payment service
- **flashSale**: Flash sale service

<br>

#### 1. Prepare MySQL Databases and Tables for Each Service

Import the corresponding database tables for each service into MySQL:

- [User service schema](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/user.sql)
- [Product service schema](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/product.sql)
- [Order service schema](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/order_record.sql)
- [Stock service schema](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/stock.sql)
- [Coupon service schema](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/coupon.sql)
- [Payment service schema](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/pay.sql)

<br>

#### 2. Prepare Protobuf Files for Each Service

These protobuf files allow Sponge to quickly create services:

- [User service protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/user.proto)
- [Product service protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/product.proto)
- [Order service protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/order.proto)
- [Stock service protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/stock.proto)
- [Coupon service protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/coupon.proto)
- [Payment service protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/pay.proto)
- [Flash sale service protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/flashSale.proto)
- [API Gateway protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/eshop_gw.proto)

<br>

#### 3. Generate gRPC+HTTP Hybrid Service Code Based on Protobuf

Open the sponge UI page, switch to the menu `Protobuf` --> `Create grpc+http service`, fill in the parameters, and generate 7 hybrid service codes that support both grpc and http: user, product, order, stock, coupon, pay, flashSale, as shown below:

![eshop-grpc-http-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_eshop-grpc-http-pb.png)

After downloading the code, unzip each service code into the eshop directory.

> Note: If the large repository option is enabled on the code generation page, it means that the created service is suitable for a microservice mono-repo mode.

<br>

#### 4. Generate CRUD Code Based on MySQL Tables

Open the sponge UI page, switch to the menu `Public` --> `Generate service+handler CRUD code`, fill in the parameters, and generate CRUD codes for the user, product, order, stock, coupon, and pay services, as shown below:

![eshop-service-handler](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_eshop-service-handler.png)

After downloading the code, unzip the CRUD code and move the CRUD code (the `api` and `internal` directories) into the corresponding service code (if prompted with duplicate proto files, just ignore it).

> Note: If the large repository option is enabled on the code generation page, it means that the created service is suitable for a microservice mono-repo mode.

<br>

#### 5. Generate API Gateway Service Code Based on Protobuf

Open the sponge UI page, switch to the menu `Protobuf` --> `Create grpc gateway service`, fill in the parameters, and generate the API gateway service code for eshop_gw, as shown below:

![eshop-grpc-gw-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_eshop-grpc-gw-pb.png)

After downloading the code, unzip the service code into the eshop directory.

To allow the eshop_gw service to connect to the various services, you need to generate the connection code. Open the sponge UI page, switch to the menu `Public` --> `Generate grpc connection code`, fill in the parameters, and generate the connection code for eshop_gw to connect to various grpc services, as shown below:

![eshop-grpc-conn](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_eshop-grpc-conn.png)

After downloading the code, unzip it, and move the connection code (the `internal` directory) into the eshop_gw service code.

> Note: If the large repository option is enabled on the code generation page, it means that the created service is suitable for a microservice mono-repo mode.

<br>

#### 6. Write Business Logic Code

At this point, the service framework is basically set up. Next, write the actual business logic code in the `internal/service` directory of each service.

<br>

#### 7. Start the Services

Before starting the services, modify the configuration files of each service, including the port numbers, database connections, etc. The default HTTP port for each service is 8080, and the grpc port is 8282. Since they are running locally on the same machine (local test IP is 192.168.3.90), to avoid port conflicts, the ports of each service have been modified (ports can be found and modified in the `configs/xxx.yml` directory and `api/xxx/v1/xxx.proto`). Below are the modified ports:

| Service          | Protocol     | HTTP Port | gRPC Port |
|------------------|--------------|-----------|-----------|
| eshop_gw         | HTTP         | 8080      | -         |
| user             | HTTP, gRPC   | 30080     | 30082     |
| product          | HTTP, gRPC   | 30180     | 30182     |
| order            | HTTP, gRPC   | 30280     | 30282     |
| stock            | HTTP, gRPC   | 30380     | 30382     |
| coupon           | HTTP, gRPC   | 30480     | 30482     |
| pay              | HTTP, gRPC   | 30580     | 30582     |
| flashSale        | HTTP, gRPC   | 30680     | 30682     |

> Note: If running in containers or on different machines, you don’t need to modify the default ports, just change the mapped ports.


<br>

### Testing and Verification

#### Single Service Testing

After all services have successfully started, verify that each service is functioning properly by testing the APIs of the following 7 services: user, product, order, stock, coupon, pay, and flashSale.

Open your browser and navigate to `http://localhost:<service_port>/apis/swagger/index.html` to check if the APIs of each service are working correctly. Besides testing the APIs on the Swagger page, you can also run gRPC API tests by filling in parameters in the `internal/service/xxx_client_test.go` file under each service.

<br>

#### Integration Testing

Once individual services pass the tests, use the API gateway service of `eshop_gw` to test the entire system. In your browser, visit the Swagger page of the API gateway service at [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html), as shown in the figure below:

![eshop-gw-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/eshop-gw-swagger.png)

<br>

**Testing the Submit Order API**

- The submit order process uses DTM's distributed transaction model `saga`, mainly to verify the consistency of creating orders, deducting stock, creating payment orders, and coupon data.

- To avoid order failure due to insufficient stock, set the stock before testing. In the Swagger page, find the API for setting product stock, and fill in the parameters. For example, set product ID to 1 and stock to 10:

   ```json
   {
     "productID": 1,
     "stock": 10
   }
   ```

- Test the submit order API by requesting both the non-buffered queue and buffered queue APIs. In the Swagger page, find the respective API and fill in the parameters, such as user ID = 1, product ID = 1, product quantity = 1, and order amount = 100:

   ```json
   {
     "userID": 1,
     "productID": 1,
     "productCount": 1,
     "amount": 100,
     "couponID": 0
   }
   ```

> Note: If `couponID` is not set to 0, it means a coupon will be used. If the coupon is invalid or expired, the order will fail. To ensure the order succeeds, find the API for creating a new coupon on the Swagger page, create a new coupon, get the coupon ID, and fill it into the `couponID` field of the submit order API.

<br>

**Testing the Flash Sale API**

- The flash sale process uses Kafka's message queue, DTM+Redis's two-phase message, and DTM+MySQL's `saga` distributed transaction models to verify consistency in flash sales, stock deduction, order creation, and payment order creation.

- To avoid order failure due to insufficient stock, set the stock before testing. In the Swagger page, find the API for setting product stock and fill in the parameters. For example, set product ID to 1 and stock to 10:

   ```json
   {
     "productID": 1,
     "stock": 10
   }
   ```

- Test the flash sale API to verify data consistency. In the Swagger page, find the flash sale API and fill in the parameters, such as user ID = 1 and product ID = 1:

   ```json
   {
     "userID": 1,
     "productID": 1,
     "amount": 100
   }
   ```

<br>

### Stress Testing

To perform stress testing on the `eshop_gw` API gateway service and verify the system's performance under high concurrency, use the stress testing tool [K6](https://github.com/grafana/k6). Before conducting the stress test, ensure that enough stock is set to avoid order failure.

1. For stress testing the Submit Order API scenario, use the k6 script [submitOrder.js](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/submitOrder.js) and run the following command:

    ```bash
    # 1000 virtual users, running for 10 seconds
    k6 run --vus 1000 --duration 10s test/k6/submitOrder.js
   
    # Alternatively, specify the number of virtual users and the number of request iterations, for example, 1000 virtual users performing 100,000 request iterations
    k6 run -u 1000 -i 100000 submit_order.js
    ```

2. For stress testing the Flash Sale API scenario, use the k6 script [flashSale.js](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/flashSale.js) and run the following command:

   ```bash
   # 10,000 virtual users, running for 1 second
   k6 run --vus 10000 --duration 1s test/k6/flashSale.js
   ```

> Note: The results of stress testing depend on factors such as machine configuration, network environment, and database setup. Adjust accordingly based on actual conditions.

<br>
