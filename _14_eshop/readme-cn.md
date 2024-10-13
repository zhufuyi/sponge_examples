
使用Sponge快速构建的微服务示例，主要实现抢购和订单功能，并通过分布式事务管理器DTM确保数据一致性，系统架构如下图所示：

![flash-sale-order-cache](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/flash-sale-order-cache.png)

<br>

这是[源码示例eshop](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop)，目录下包括了两个一样的代码示例，只是代码结构稍有不同，主要是为了展示说明sponge支持创建不同仓库模式的微服务项目，`example-1-multi-repo`是适合微服务多仓库(multi-repo)，`example-2-mono-repo`是适合微服务单仓库(mono-repo)。

<br>

### 准备环境

为了构建这个电商系统，准备生成代码工具：

- **[Sponge](https://github.com/zhufuyi/sponge)**: 快速创建不同类型服务和模块代码。点击 [Sponge 安装说明](https://github.com/zhufuyi/sponge/blob/main/assets/install-cn.md) 以获取更多信息。

准备以下依赖服务：

- **[DTM](https://github.com/dtm-labs/dtm)**：分布式事务管理器，用于保证系统在多服务调用中的数据一致性。
- **Redis**：结合 DTM，用作秒杀抢购的缓存。
- **Kafka**：作为订单系统的消息队列，用于处理订单异步消息。
- **MySQL**：为各个服务提供数据存储。

所有这些服务将运行在虚拟机中，虚拟机的 IP 地址为 `192.168.3.37`。

<br>

### 启动 DTM 服务

DTM 是本系统的核心组件之一，负责管理订单和秒杀抢购的分布式事务。需要启动两个 DTM 服务实例，分别用于 MySQL 和 Redis 存储。

DTM下载地址：https://github.com/dtm-labs/dtm/releases/tag/v1.18.0

| 服务名       | 端口配置                 |
|-------------|-----------------------|
| DTM-MySQL   | HTTP: 36789, gRPC: 36790 |
| DTM-Redis   | HTTP: 35789, gRPC: 35790 |

<br>

#### 1. 启动使dtm-mysql服务

- 在 MySQL 中导入 DTM 需要的表结构：
  - [dtmcli.barrier.mysql.sql](https://github.com/dtm-labs/dtm/blob/main/sqls/dtmcli.barrier.mysql.sql)
  - [dtmsvr.storage.mysql.sql](https://github.com/dtm-labs/dtm/blob/main/sqls/dtmsvr.storage.mysql.sql)

- 修改 DTM 配置文件 ([示例配置文件](https://github.com/dtm-labs/dtm/blob/main/conf.sample.yml))：

   ```yaml
   Store: # specify which engine to store trans status
     Driver: 'mysql'
     Host: '192.168.3.37'
     User: 'root'
     Password: '123456'
     Port: 3306
     Db: 'dtm'
   
   # 使用注册与发现(sponge的dtm驱动已支持etcd、consul、nacos)
   #MicroService:
      #Driver: 'dtm-driver-sponge'
      #Target: 'etcd://127.0.0.1:2379/dtmservice'
      #EndPoint: 'grpc://127.0.0.1:36790'
   ```

- 启动 DTM 服务：

   ```bash
   #dtm 服务端口 http: 36789, grpc: 36790 
   ./dtm -c conf.yml
   ```

<br>

#### 2. 启动dtm-redis服务

- 修改 DTM 配置文件 ([示例配置文件](https://github.com/dtm-labs/dtm/blob/main/conf.sample.yml))：

   ```yaml
   Store: # specify which engine to store trans status
       Driver: 'redis'
       Host: '192.168.3.37'
       User: 'default'
       Password: '123456'
       Port: 6379


   # 修改默认dtm服务端口
   HttpPort: 35789
   GrpcPort: 35790


   # 如果使用注册与发现(sponge的dtm驱动已支持etcd\consul\nacos)，则需要配置如下(使用http协议)
   #MicroService:
      #Driver: 'dtm-driver-sponge'
      #Target: 'etcd://127.0.0.1:2379/dtmservice'
      #EndPoint: 'http://127.0.0.1:35789'
   ```

- 启动 DTM 服务：

   ```bash
   #dtm 服务端口 http: 35789, grpc: 35790 
   ./dtm -c conf.yml
   ```

<br>

### 使用sponge快速构建电商系统

简易版电商系统由以下 8 个微服务组成：

- **eshop_gw**：API 网关服务
- **user**：用户服务
- **product**：商品服务
- **order**：订单服务
- **stock**：库存服务
- **coupon**：优惠券服务
- **pay**：支付服务
- **flashSale**：秒杀抢购服务

<br>

#### 1. 准备各个服务mysql数据库和表

将以下服务对应的数据库表导入 MySQL：

- [user 服务的库和表](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/user.sql)
- [product 服务的库和表](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/product.sql)
- [order 服务的库和表](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/order_record.sql)
- [stock 服务的库和表](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/stock.sql)
- [coupon 服务的库和表](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/coupon.sql)
- [pay 服务的库和表](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/sql/pay.sql)

<br>

#### 2. 准备各个服务的protobuf文件

这些protobuf文件给sponge快速创建服务：

- [user 服务的protobuf文件](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/user.proto)
- [product 服务的protobuf文件](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/product.proto)
- [order 服务的protobuf文件](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/order.proto)
- [stock 服务的protobuf文件](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/stock.proto)
- [coupon 服务的protobuf文件](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/coupon.proto)
- [pay 服务的protobuf文件](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/pay.proto)
- [flashSale 秒杀抢购服的protobuf文件](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/flashSale.proto)
- [eshop_gw api 网关服务的protobuf文件](https://github.com/zhufuyi/sponge_examples/tree/main/_14_eshop/test/protobuf/eshop_gw.proto)

<br>

#### 3. 基于protobuf生成gRPC+HTTP混合服务代码

打开sponge的UI页面，切换到菜单`Protobuf` --> `创建grpc+http服务`，填写参数，分别生成7个支持grpc+http混合服务代码user、product、order、stock、coupon、pay、flashSale，如下图所示：

![eshop-grpc-http-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/eshop-grpc-http-pb.png)

下载代码后，分别解压各个服务代码到eshop目录下。

> 注：在生成代码页面如果参数开启了大仓库类型，表示创建的服务适合微服务单仓库(mono-repo)模式。

<br>

#### 4. 基于mysql表生成CRUD代码

打开sponge的UI页面，切换到菜单`Public` --> `生成service+handler CRUD代码`，填写参数，分别生成user、product、order、stock、coupon、pay服务的CRUD代码，如下图所示：

![eshop-service-handler](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/eshop-service-handler.png)

下载代码后，分别解压CRUD代码，把CRUD代码(api和internal两个目录)移动到对应服务代码中(如果提示proto文件重复，忽略即可)。

> 注：在生成代码页面如果参数开启了大仓库类型，表示创建的服务适合微服务单仓库(mono-repo)模式。

<br>

#### 5. 基于protobuf生成api网关服务代码

打开sponge的UI页面，切换到菜单`Protobuf` --> `创建grpc网关服务`，填写参数，生成eshop_gw的api网关服务代码，如下图所示：

![eshop-grpc-gw-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/eshop-grpc-gw-pb.png)

下载代码后，解压服务代码到eshop目录下。

为了让eshop_gw服务可以连接各个服务，需要生成连接代码。打开sponge的UI页面，切换到菜单`Public` --> `生成grpc连接代码`，填写参数，生成eshop_gw的连接各个grpc代码，如下图所示：

![eshop-grpc-conn](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/eshop-grpc-conn.png)

下载代码后，解压代码，把连接代码(internal目录)移动到eshop_gw服务代码中。

> 注：在生成代码页面如果参数开启了大仓库类型，表示创建的服务适合微服务单仓库(mono-repo)模式。

<br>

#### 6. 填写业务逻辑代码

到此为止，服务框架已经基本搭建完毕。接下来分别在各服务的 `internal/service` 目录下编写实际的业务逻辑代码。

<br>

#### 7. 启动服务

启动服务前先修改各个服务的配置文件，修改端口、数据库连接等信息。各个服务默认的http端口(8080)，grpc端口(8282)，因为在同一台主机本地(本地测试ip为192.168.3.90)运行，为了防止端口冲突，所以在配置已经修改了各个服务的端口(在`configs/xxx.yml`目录下和`api/xxx/v1/xxx.proto`中修改端口)，下面是已修改的端口：

| 服务             | 协议        | HTTP 端口  | gRPC 端口 |
|------------------|-------------|------------|-----------|
| eshop_gw         | HTTP        | 8080       | -         |
| user             | HTTP, gRPC  | 30080      | 30082     |
| product          | HTTP, gRPC  | 30180      | 30182     |
| order            | HTTP, gRPC  | 30280      | 30282     |
| stock            | HTTP, gRPC  | 30380      | 30382     |
| coupon           | HTTP, gRPC  | 30480      | 30482     |
| pay              | HTTP, gRPC  | 30580      | 30582     |
| flashSale        | HTTP, gRPC  | 30680      | 30682     |

> 注：如果在容器或不同机器上运行，不需要修改默认的端口，只需修改映射端口。

<br>

### 测试与验证

#### 单个服务测试

各个服务启动成功之后，先验证单个服务是否正常运行，分别测试user、product、order、stock、coupon、pay、flashSale这7个服务的api。

在浏览器访问 `http://localhost:<服务端口>/apis/swagger/index.html`，验证各个服务的api是否正常，除了在swaggers页面测试api，也可以在各个服务下的internal/service/xxx_client_test.go 文件中填写参数后运行测试grpc api。

<br>

#### 集成测试

各个服务测试正常后，通过eshop_gw的api网关服务测试整个系统是否正常运行。在浏览器访问api网关服务的swagger页面，[http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html)，如下图所示：

![eshop-gw-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/eshop-gw-swagger.png)

<br>

**测试提交订单 API**

- 提交订单使用了dtm的分布式事务模式saga，主要验证创建订单、扣减库存、创建支付订单、优惠券数据是否一致。

- 为了避免库存不足导致订单失败，测试前先设置库存，在swagger页面找到设置产品库存api，填写参数，例如产品id为1、库存为10

  ```json
  {
    "productID": 1,
    "stock": 10
  }
  ```

- 测试提交订单分别请求无缓冲队列和缓冲队列的api，在swagger页面找到对应的api，填写参数，例如用户id为1、产品id为1、产品数量为1、订单金额为100

   ```json
   {
     "userID": 1,
     "productID": 1,
     "productCount": 1,
     "amount": 100,
     "couponID": 0
   }
   ```

> 注：如果设置couponID不为0，表示使用优惠券，如果优惠券失效或过期，会导致订单失败。如果想使得订单成功的话，在swagger页面找到创建优惠券api创建新的优惠券，然后获取得到优惠券id，填写到提交订单api的couponID字段。

<br>

**测试秒杀抢购API**

- 秒杀抢购使用了kafka的消息队列、dtm+redis的二阶段消息、dtm+mysql的saga两个分布式事务模式，主要验证秒杀抢购、扣减库存、创建订单、扣减库存、创建支付订单数据是否一致。

- 为了避免库存不足导致订单失败，测试前先设置库存，在swagger页面找到设置产品库存api，填写参数，例如产品id为1、库存为10

   ```json
   {
    "productID": 1,
    "stock": 10
   }
   ```

- 测试秒杀抢购api，验证数据是否一致，在swagger页面找到秒杀抢购api，填写参数，例如用户id为1、产品id为1

   ```json
   {
     "userID": 1,
     "productID": 1,
     "amount": 100
   }
   ```

<br>

### 压力测试

这里使用压力测试工具[k6](https://github.com/grafana/k6)对eshop_gw的api网关服务进行压力测试，验证系统在高并发场景下的性能。压测前先设置足够的库存数量，否则会导致订单失败。

1. 压测提交订单 API的场景，使用k6的脚本 [submitOrder.js](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/submitOrder.js) ，运行命令如下：

    ```bash
    # 1000个虚拟用户，持续10秒
    k6 run --vus 1000 --duration 10s test/k6/submitOrder.js
   
   # 或者，指定虚拟用户数和请求迭代次数，例如1000个虚拟用户共同完成迭代请求100000次
   k6 run -u 1000 -i 100000 submit_order.js
    ```

2. 压测秒杀抢购 API的场景，使用k6的脚本 [flashSale.js](https://github.com/zhufuyi/sponge_examples/blob/main/_14_eshop/test/flashSale.js) ，运行命令如下：

   ```bash
   # 10000个虚拟用户，持续1秒
   k6 run --vus 10000 --duration 1s test/k6/flashSale.js
   ```

> 注：压测结果与机器配置、网络环境、数据库配置等因素有关，请根据实际情况进行调整。

<br>
