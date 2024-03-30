
[**order-grpc-distributed-transaction 中文说明**](https://juejin.cn/post/7307811441666850870)

<br>

The Order-system created by using sponge is a `multi-repo` type of microservice. Although multiple service codes are placed under a directory `9_order-grpc-distributed-transaction`, the codes between different services cannot be reused, so the codes between services are completely independent.

> Tip: Sponge also supports the creation of a microservice `mono-repo` type where code can be reused between different services, which is simpler.


<br>

### Order System Introduction

The order system is the core system of a trading platform, involving various complex tasks that require careful consideration of factors such as business requirements, performance, scalability, and security.

Breaking down the order system into services like order service, stock service, coupon service, and pay service, each with its independent database, is a common practice. The order processing inevitably involves distributed transactions, such as ensuring atomicity in actions like creating an order and deducting stock. In a distributed system, ensuring the atomicity of these operations faces challenges such as process crashes, idempotence issues, rollback problems, and precise compensation problems.

In a monolithic order system, using the built-in transaction support of the database can easily solve these problems. However, when services are modularized, the issues of a distributed system must be considered. Common solutions for distributed transactions include message queue-based and state machine-based approaches, both of which are relatively heavyweight, making the order system more complex. [DTM](https://github.com/dtm-labs/dtm), as an alternative solution for distributed transactions, significantly simplifies the architecture of the order system and elegantly addresses data consistency issues in distributed transactions.

![order-system](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/order-system.png)

When a frontend requests the **gRPC gateway service order_gw** to submit an order API, the server performs the following operations:

- **order service**: Creates an order in the order table with a unique key as the order ID.
- **stock service**: Deducts stock from the stock table. If the stock is insufficient, the global transaction automatically rolls back.
- **coupon service**: Marks the coupon as used in the coupon table. If the coupon is invalid, the global transaction automatically rolls back.
- **pay service**: Creates a payment order in the pay table and informs the user to proceed to the payment page.

<br>

### Preparation

1. Prepare a MySQL service by quickly launching it using the script [docker-compose.yaml](https://github.com/zhufuyi/sponge/blob/main/test/server/mysql/docker-compose.yaml).
2. Import the prepared SQL into MySQL.

  - DTM-related SQL:
    - [dtmcli.barrier.mysql.sql](https://github.com/dtm-labs/dtm/blob/main/sqls/dtmcli.barrier.mysql.sql)
    - [dtmsvr.storage.mysql.sql](https://github.com/dtm-labs/dtm/blob/main/sqls/dtmsvr.storage.mysql.sql)
  - Order-related SQL:
    - [order.sql](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/order/test/sql/order.sql)
    - [stock.sql](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/stock/test/sql/stock.sql)
    - [coupon.sql](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/coupon/test/sql/coupon.sql)
    - [pay.sql](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/pay/test/sql/pay.sql)

3. Prepare Proto files:

    - [order.proto](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/order/api/order/v1/order.proto) for creating the order service.
    - [stock.proto](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/stock/api/stock/v1/stock.proto) for creating the stock service.
    - [coupon.proto](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/coupon/api/coupon/v1/coupon.proto) for creating the coupon service.
    - [pay.proto](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/pay/api/pay/v1/pay.proto) for creating the pay service.
    - [order_gw.proto](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/order_gw/api/order_gw/v1/order_gw.proto) for creating the gRPC gateway service.

4. Install the tool [sponge](https://github.com/zhufuyi/sponge/blob/main/assets/install-en.md). After installing sponge, run the following command to open the UI interface for generating code:

```bash
sponge run
```

5. Start the DTM service using the `docker-compose.yml` script:

```yaml
version: '3'
services:
  dtm:
    image: yedf/dtm
    container_name: dtm
    restart: always
    environment:
      STORE_DRIVER: mysql
      STORE_HOST: '192.168.3.37'
      STORE_USER: root
      STORE_PASSWORD: '123456'
      STORE_PORT: 3306
    #volumes:
    #  - /etc/localtime:/etc/localtime:ro
    #  - /etc/timezone:/etc/timezone:ro
    ports:
      - '36789:36789'
      - '36790:36790'
```

Modify the STORE_xxx related parameters and then start the DTM service:

```bash
docker-compose up -d
```

<br>

### Quickly Create Microservices for the Order System

#### Generate gRPC service codes for Order, Stock, Coupon, Pay, and Order_gw

Go to the sponge UI interface, click on the left menu bar **Protobuf** -> **create microservice project**, fill in the parameters, and generate service codes for order, stock, coupon, and pay services.

Quickly create the order service, as shown in the image below:

![order-grpc-pb-order](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_order-grpc-pb-order.png)

Quickly create the stock service, as shown in the image below:

![order-grpc-pb-stock](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_order-grpc-pb-stock.png)

Quickly create the coupon service, as shown in the image below:

![order-grpc-pb-coupon](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_order-grpc-pb-coupon.png)

Quickly create the pay service, as shown in the image below:

![order-grpc-pb-pay](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_order-grpc-pb-pay.png)

Quickly create the gRPC gateway service order_gw. Click on the left menu bar **Protobuf** -> **create grpc gateway project**, fill in the parameters, and click the download code button, as shown in the image below:

![order-http-pb-order-gw](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_order-http-pb-order-gw.png)

Rename the generated services to order, stock, coupon, pay, and order_gw. Open five terminals, each corresponding to one service.

<br>

#### Configure and Run the Stock Service

Switch to the stock service directory and perform the following steps:

1. Generate and automatically merge API-related code:

```bash
make proto
```

2. Add MySQL connection code:

```bash
make patch TYPE=mysql-init
```

3. Open the configuration file `configs/stock.yml` and modify the MySQL address and account information. Change the default gRPC server port to avoid port conflicts.

```yaml
mysql:
  dsn: "root:123456@(192.168.3.37:3306)/eshop_stock?parseTime=true&loc=Local&charset=utf8mb4"

grpc:
  port: 28282
  httpPort: 28283
```

4. Add business logic code for deducting and compensating stock in the generated template code. View the code [internal/service/stock.go](https://github.com/zhufuyi/sponge_examples/blob/main/9_order-grpc-distributed-transaction/stock/internal/service/stock.go).

5. Compile and start the stock service:

```bash
   make run
```

The [stock service source code](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/stock) is completed following the above steps.

<br>

#### Configure and Run the Coupon Service

Switch to the coupon service directory, and the steps are the same as **configuring and running the stock service**, except for the business logic code. This is the [coupon service source code](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/coupon).

Configure and fill in the specific business logic code, then compile and start the coupon service:

```bash
make run
```

<br>

#### Configure and Run the Pay Service

Switch to the pay service directory, and the steps are the same as **configuring and running the stock service**. This is the [pay service source code](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/pay).

Configure and fill in the specific business logic code, then compile and start the pay service:

```bash
make run
```

<br>

#### Configure and Run the Order Service

Switch to the order service directory, and the steps are the same as **configuring and running the stock service**. This is the [order service source code](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction/order).

As order submission requires informing the DTM service of the gRPC service addresses for the order, stock, coupon, and pay services, to manage distributed transactions, the addresses need to be configured. Open the configuration file `configs/order.yml` and add the service addresses and DMT service address configuration, as shown below:

```yaml
grpcClient:
  - name: "order"
    host: "127.0.0.1"
    port: 8282
  - name: "coupon"
    host: "127.0.0.1"
    port: 18282
    registryDiscoveryType: ""
    enableLoadBalance: false
  - name: "stock"
    host: "127.0.0.1"
    port: 28282
    registryDiscoveryType: ""
    enableLoadBalance: false
  - name: "pay"
    host: "127.0.0.1"
    port: 38282
    registryDiscoveryType: ""
    enableLoadBalance: false

dtm:
  addr: "127.0.0.1:36790"
```

Add the new fields to the configuration file and update the corresponding Go code:

```bash
make update-config
```

Add the business logic code for submitting orders, creating orders, and canceling orders to the generated template code. View the code [internal/service/order.go](https://github.com/zhufuyi/sponge_examples/blob/main/9_order-grpc-distributed-transaction/order/internal/service/order.go).

Configure and fill in the specific business logic code, then compile and start the order service:

```bash
make run
```

<br>

#### Configure and Run the gRPC Gateway Service order_gw

1. Generate gRPC service connection code.

   As the gRPC gateway service order_gw as the entry point for requests and needs to connect to the order service, generate the connection code. In the sponge UI interface, click on the left menu bar **Public** -> **generate grpc connection code**, fill in the parameters, and download the code.

   ![order-grpc-conn](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_order-grpc-conn.png)

   Extract the code, move the `internal` directory to the `order_gw` service directory.

2. Copy proto files.

   Since the gRPC gateway service order_gw needs to know which API interfaces the order service provides, copy the order service's proto file. In the terminal, navigate to the order_gw directory and execute the command:

```bash
make copy-proto SERVER=../order
```

3. Open the configuration file `configs/order_gw.yml` and configure the order service address.

```yaml
grpcClient:
  - name: "order"
    host: "127.0.0.1"
    port: 8282
    registryDiscoveryType: ""
    enableLoadBalance: false
```

4. Generate and automatically merge API-related code.

```bash
make proto
```

5. Fill in the business logic code, which involves converting HTTP requests into gRPC requests. Here, you can directly use the generated template code example. Click to view the code [internal/service/order_gw.go](https://github.com/zhufuyi/sponge_examples/blob/main/9_order-grpc-distributed-transaction/order_gw/internal/service/order_gw.go).

<br>

After configuring and filling in the business logic code, compile and start the gRPC gateway service order_gw:

```bash
make run
```

<br>

### Testing Distributed Transactions

Open the swagger interface `http://localhost:8080/apis/swagger/index.html` in the browser to test the order submission API interface.

On the DTM management interface `http://localhost:36789`, you can view the status and details of distributed transactions, and you can also view the log information of various services to understand the API interface coordinated by DTM.

<br>

#### Test Successful Order Submission Scenario

On the swagger interface, fill in the request parameters.

![order-http-pb-order-gw-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/order-http-pb-order-gw-swagger.png)

Click the Execute button to test. The order is successfully submitted, which can be seen from the DTM management interface and the logs of various services.

<br>

#### Test Failed Order Submission Scenario

1. The order fails due to an invalid coupon.

With the request parameters unchanged,

```json
{
  "userId": 1,
  "productId": 1,
  "amount": 1100,
  "productCount": 1,
  "couponId": 1
}
```

Click the Execute button directly for testing. Although an order ID is returned (this does not mean the order is successful, you actually need to get the successful order status before performing subsequent operations), you can see from the DTM management interface and the coupon service log that the order status is failed because the coupon has been used. It returned an `Aborted` error. After DTM receives the `Aborted` error message, it will compensate for the **order creation** and **coupon deduction** branch transactions to ensure data consistency.

<br>

2. The order fails due to insufficient inventory.

Fill in the request parameters, the value of the productCount field is 1000, which is definitely greater than the inventory quantity, and set the couponId parameter to 0 to indicate that no coupon is used.

```json
{
  "userId": 1,
  "productId": 1,
  "amount": 1100000,
  "productCount": 1000,
  "couponId": 0
}
```

Click the Execute button to test. Although an order ID is returned (this does not mean the order is successful), you can see from the DTM management interface and the inventory service log that the order status is failed due to insufficient inventory. It returned an `Aborted` error. After DTM receives the `Aborted` error message, it will compensate for the **order creation** branch transaction to ensure data consistency.

<br>

#### Test Simulated Process Crash, Successful Order Submission After Recovery Scenario

Stop the inventory service stock, then fill in the request parameters on the swagger interface:

```json
{
  "userId": 1,
  "productId": 1,
  "amount": 1100,
  "productCount": 1,
  "couponId": 0
}
```

Click the Execute button to test. Although an order ID is returned (this does not mean the order is successful), you can see from the DTM management interface that the order status is `submitted`. DTM will keep trying to connect to the inventory service stock. The retry is by default an exponential backoff algorithm. At this time, start the inventory service. After DTM connects to the inventory service, it completes the subsequent branch transactions and finally successfully completes the order, ensuring data consistency. Of course, you can also actively force the cancellation of the order on the business side when it times out. After DTM receives the forced cancellation of the order, it will compensate for the **order creation** branch transaction to ensure data consistency.

<br>

### Summary

This article introduces the quick construction of a simple order system from scratch. Using Sponge makes it easy to build microservices, and DTM elegantly solves distributed transaction issues related to orders. Developing an order system becomes straightforward, allowing developers to spend their efforts on business development.

Each service comes with commonly used service governance features, such as service registration and discovery, rate limiting, circuit breaking, distributed tracing, monitoring, performance analysis, resource statistics, CICD, etc. These features can be uniformly enabled or disabled in the YAML configuration file.

Of course, this is not a complete order system; it only covers the simple business logic of submitting orders. If you need to build your own order system, this can serve as a reference. Following the above steps, it's easy to add modules like product service, logistics service, user service, etc., to create an e-commerce platform.
