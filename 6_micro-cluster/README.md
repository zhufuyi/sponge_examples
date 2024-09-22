## Microservices Cluster Example

Here is an example of quickly creating a microservices cluster using sponge based on protobuf files. It supports two directory structures: multi-repo and mono-repo for microservices. The difference between multi-repo and mono-repo can be seen in [Repository Type Introduction Document](https://go-sponge.com/repository-type).

- [**example-1-multi-repo**](https://github.com/zhufuyi/sponge_examples/tree/main/6_micro-cluster/example-1-multi-repo): microservice multi repository (multi-repo) example. In a multi-repo structure, multiple repositories are used, with each service completely decoupled and independent. Each service has its own `go.mod` file, as well as separate `api` and `third_party` directories.
- [**example-2-mono-repo**](https://github.com/zhufuyi/sponge_examples/tree/main/6_micro-cluster/example-2-mono-repo): microservice monolithic repository (mono-repo) example. In a mono-repo structure, there is only one repository, with a single `go.mod` file, and the `api` and `third_party` directories are shared as common directories among all services.

The development and code generation processes are the same for both structures; the difference lies in the organization of the code. The mono-repo structure is generally recommended.

<br>
<br>

Here is an example of quickly setting up a microservices cluster from scratch based on protobuf files. The setup steps are essentially the same for both `multi-repo` and `mono-repo` repository types, only the following two steps are different:

- When generating code, the `mono-repo` repository type must select the monolithic repository option, while the `multi-repo` repository type does not require selecting this option.
- For the `multi-repo` repository type, if the protobuf files in the current service depend on protobuf files from other services, you need to copy the dependent protobuf files into the `api` directory of the current service. A `make copy-proto` command is provided for convenient copying. The `mono-repo` repository type does not require copying the dependent protobuf files.

<br>

By taking a simple e-commerce microservice cluster as an example, the product details page contains product information, inventory information, and product evaluation information. These data are scattered in different microservices. The grpc gateway service assembles the required data and returns it to the product details page, as shown in the following figure:

![micro-cluster](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-cluster.png)

Four proto files have been prepared in advance. Each proto file generates corresponding service code:

- The comment.proto file defines grpc methods to obtain comment data based on the product ID and is used to generate the comment grpc service code.
- The inventory.proto file defines grpc methods to obtain inventory data based on the product ID and is used to generate the inventory grpc service code.
- The product.proto file defines grpc methods to obtain product details based on the product ID and is used to generate the product grpc service code.
- The shop_gw.proto file defines grpc methods to assemble the data required by the product details page based on the product ID and is used to generate the shop grpc gateway service code.

<br>

### Quickly generate and start comment, inventory, and product microservices

#### Generate comment, inventory, and product microservice code

Enter the Sponge UI interface, click the left menu bar 【Protobuf】→【Create microservice project】, fill in the respective parameters of the comment, inventory, and product, and generate the comment, inventory, and product service code respectively.

The microservice framework uses [grpc](https://github.com/grpc/grpc-go) and also contains common service governance function codes, build deployment scripts, etc. You can choose whatever database you want.

Quickly create a comment service as shown below:

![micro-rpc-pb-comment](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-pb-comment.png)

Quickly create an inventory service as shown below:

![micro-rpc-pb-inventory](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-pb-inventory.png)

Quickly create a product service as shown below:

![micro-rpc-pb-product](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-pb-product.png)

Open three terminals, one for each of the comment, inventory, and product services.

<br>

#### Start the comment service

Switch to the comment directory and perform the following steps:
(1) Generate pb.go code, generate template code, and generate test code

```shell
make proto
```

(2) Open `internal/service/comment.go`, which is the generated template code. There is a line of panic code prompting you to fill in the business logic code. Fill in the business logic here, for example, fill in the return value:

```go
	return &commentV1.ListByProductIDReply{
		Total:     11,
		ProductID: 1,
		CommentDetails: []*commentV1.CommentDetail{
			{
				Id:       1,
				Username: "Mr Zhang",
				Content:  "good",
			},
			{
				Id:       2,
				Username: "Mr Li",
				Content:  "good",
			},
			{
				Id:       3,
				Username: "Mr Wang",
				Content:  "not good",
			},
		},
	}, nil
```

(3) Open the `configs/comment.yml` configuration file, find grpc, and modify the port and httpPort values under it:

```yaml
grpc:   
  port: 18203              # listen port   
  httpPort: 18213        # profile and metrics ports  
```

(4) Compile and start the comment service

```shell
make run
```

<br>

#### Start the inventory service

Switch to the inventory directory and perform the same steps as for the comment:

(1) Generate pb.go code, generate template code, and generate test code

```shell
make proto  
```

(2) Open `internal/service/inventory.go`, which is the generated template code. There is a line of panic code prompting you to fill in the business logic code. Fill in the business logic here, for example, fill in the return value:

```go
	return &inventoryV1.GetByIDReply{
		InventoryDetail: &inventoryV1.InventoryDetail{
			Id:      1,
			Num:     999,
			SoldNum: 111,
		},
	}, nil
```

(3) Open the `configs/inventory.yml` configuration file, find grpc, and modify the port and httpPort values under it:

```yaml
grpc:    
  port: 28203              # listen port  
  httpPort: 28213        # profile and metrics ports   
```

(4) Compile and start the inventory service

```shell
make run
```

<br>

#### Start the product service

Switch to the product directory and perform the same steps as for the comment:

(1) Generate pb.go code, generate template code, and generate test code

```shell
make proto
```

(2) Open `internal/service/product.go`, which is the generated template code. There is a line of panic code prompting you to fill in the business logic code. Fill in the business logic here, for example, fill in the return value:

```go
	return &productV1.GetByIDReply{
		ProductDetail: &productV1.ProductDetail{
			Id:          1,
			Name:        "Data cable",
			Price:       10,
			Description: "Android type c data cable",
		},
		InventoryID: 1,
	}, nil
```

(3) Open `the configs/product.yml` configuration file, find grpc, and modify the port and httpPort values under it:

```yaml
grpc:    
  port: 38203              # listen port  
  httpPort: 38213        # profile and metrics ports
```

(4) Compile and start the product service

```shell
make run
```

After the comment, inventory, and product microservices have started successfully, you can now generate and start the gateway service.

<br>

### Quickly generate and start the grpc gateway service

Enter the Sponge UI interface, click the left menu bar 【Protobuf】→【create grpc gateway project】, fill in some parameters to generate the grpc gateway project code.

The web framework uses [gin](https://github.com/gin-gonic/gin), which also contains swagger documentation, common service governance function codes, build deployment scripts, etc.

![micro-rpc-gw-pb-shopgw](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-gw-pb-shopgw.png)

In order to connect to the comment, inventory, and product grpc services, you need to generate additional connection grpc service codes. Click the left menu bar 【Public】→ 【generate grpc connection code】, fill in the parameters to generate the code, and then move the generated connection grpc service code to the grpc gateway project code directory.

![micro-rpc-cli](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-cluster-rpc-cli.png)

Switch to the shop_gw directory and perform the following steps:

(1) Copy proto file from grpc services comment,inventory,product.

```bash
make copy-proto SERVER=../comment,../inventory,../product
```

(2) Open the `configs/shop_gw.yml` configuration file, find grpcClient, and add the addresses of the comment, inventory, and product grpc services:

```yaml
grpcClient:
  - name: "comment"                # rpc service name, used for service discovery
    host: "127.0.0.1"              # rpc service address, used for direct connection
    port: 18282                    # rpc service port
    registryDiscoveryType: ""      # registration and discovery types: consul, etcd, nacos, if empty, connecting to server using host and port
  - name: "inventory"              # rpc service name, used for service discovery
    host: "127.0.0.1"              # rpc service address, used for direct connection
    port: 28282                    # rpc service port
    registryDiscoveryType: ""      # registration and discovery types: consul, etcd, nacos, if empty, connecting to server using host and port
  - name: "product"                # rpc service name, used for service discovery
    host: "127.0.0.1"              # rpc service address, used for direct connection
    port: 38282                    # rpc service port
    registryDiscoveryType: ""      # registration and discovery types: consul, etcd, nacos, if empty, connecting to server using host and port
```

(3) Generate pb.go code, generate registered routing code, generate template code, and generate swagger documentation

```shell
make proto
```

(4) Open `internal/service/shop_gw.go`, which is the generated API interface code. Fill in the business logic code here, fill in the following simple business logic code:

```go
package service

import (
	"context"

	commentV1 "shop_gw/api/comment/v1"
	inventoryV1 "shop_gw/api/inventory/v1"
	productV1 "shop_gw/api/product/v1"
	shop_gwV1 "shop_gw/api/shop_gw/v1"
	"shop_gw/internal/ecode"
	"shop_gw/internal/rpcclient"

	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
)

var _ shop_gwV1.ShopGwLogicer = (*shopGwClient)(nil)

type shopGwClient struct {
	commentCli   commentV1.CommentClient
	inventoryCli inventoryV1.InventoryClient
	productCli   productV1.ProductClient
}

// NewShopGwClient create a client
func NewShopGwClient() shop_gwV1.ShopGwLogicer {
	return &shopGwClient{
		commentCli:   commentV1.NewCommentClient(rpcclient.GetCommentRPCConn()),
		inventoryCli: inventoryV1.NewInventoryClient(rpcclient.GetInventoryRPCConn()),
		productCli:   productV1.NewProductClient(rpcclient.GetProductRPCConn()),
	}
}

// GetDetailsByProductID get page detail by product id
func (c *shopGwClient) GetDetailsByProductID(ctx context.Context, req *shop_gwV1.GetDetailsByProductIDRequest) (*shop_gwV1.GetDetailsByProductIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	// fill in the business logic code here

	productReply, err := c.productCli.GetByID(ctx, &productV1.GetByIDRequest{
		Id: req.ProductID,
	})
	if err != nil {
		return nil, err
	}
	logger.Info("get product info successfully", interceptor.CtxRequestIDField(ctx))

	inventoryReply, err := c.inventoryCli.GetByID(ctx, &inventoryV1.GetByIDRequest{
		Id: productReply.InventoryID,
	})
	if err != nil {
		return nil, err
	}
	logger.Info("get inventory info successfully", interceptor.CtxRequestIDField(ctx))

	commentReply, err := c.commentCli.ListByProductID(ctx, &commentV1.ListByProductIDRequest{
		ProductID: req.ProductID,
	})
	if err != nil {
		return nil, err
	}
	logger.Info("list comments info successfully", interceptor.CtxRequestIDField(ctx))

	return &shop_gwV1.GetDetailsByProductIDReply{
		ProductDetail:   productReply.ProductDetail,
		InventoryDetail: inventoryReply.InventoryDetail,
		CommentDetails:  commentReply.CommentDetails,
	}, nil
}
```

(5) Compile and start the shop_gw service

```shell
make run
```

You can test the API interface by opening [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html) in your browser.

![micro-cluster-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/micro-rpc-gw-pb-shopgw-swagger.png)

<br>
