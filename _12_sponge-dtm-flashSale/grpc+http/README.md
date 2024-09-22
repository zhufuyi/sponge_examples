
## Overview

This is a sample project that demonstrates how to use Sponge to implement a flash sale service (grpc+http) using Redis as the storage engine and dtm as the transaction manager. The project is a simple implementation of a flash sale service that allows users to purchase a limited number of items for a specific price.

<br>

### Getting Started Quickly

1. **Start Redis Service**

2. **Start dtm Service**
    - Download the [dtm](https://github.com/dtm-labs/dtm/releases/tag/v1.18.0) executable file, modify the default dtm configuration to use Redis as storage, and start the dtm service:
      ```bash
      dtm -c conf.yml
      ```

3. **Configure the Project**
    - Clone the project code locally and open the configuration file `configs/flashSale.yml`. Update the `Redis` and `dtmservice` configuration settings by replacing the default IP addresses (e.g., 192.168.3.37 and 192.168.3.90) with the actual IP addresses of your environment,  if all services are running locally, fill in 127.0.0.1.

4. **Compile and Start the Service**
    - you can compile and run the service using the following command:
      ```bash
      cd cmd/flashSale
      go run main.go
      ```
    - Alternatively, if sponge is installed, you can directly run the service with:
      ```bash
      make run
      ```

<br>

### Test the Flash Sale API

The service supports both HTTP and GRPC protocols, so you can test the flash sale API on the Swagger page or use the built-in GRPC client to test the flash sale API.

**(1) Example of testing gRPC API:**

Open the code file `internal/service/flashSale_client_test.go`, fill in the parameters, and start testing.

You can use Goland IDE to test:

![flash-sale-grpc-client-test](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/flash-sale-grpc-client-test.png)

You can also use the go test command to test:

```bash
	#Fill in parameters before executing commands

	#Set inventory quantity
	go test -run Test_service_flashSale_methods/SetProductStock
	 
	#Flash sale
	go test -run Test_service_flashSale_methods/FlashSale
```

<br>

**(2) Testing HTTP api**

Access [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html) in your browser and use the Swagger UI to test the flash sale API.

![flash-sale-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/flash-sale-swagger.png)

**API Testing Examples:**

1. **Set Inventory API**
    - Request example parameters:
      ```json
      {
        "productID": 1,
        "stock": 3
      }
      ```

2. **Flash Sale Request API**
    - Request example parameters:
      ```json
      {
        "productID": 1,
        "amount": 100,
        "userID": 1
      }
      ```
    - If the inventory is insufficient, a 409 status code will be returned, indicating the flash sale has failed.

<br>

### Performance Testing

For performance testing, it is recommended to run the service in the background to avoid terminal log output affecting the service performance. If sponge is installed, you can start the service in the background using the following command:
```bash
make run-nohup
```

**gRPC api performance pressure test**

Open the file `internal/service/flashSale_client_test.go`, you can use Goland IDE test, you can also use command line test.

```bash
#Please fill in the parameter
go test -run Test_service_flashSale_benchmark/FlashSale
```

**HTTP api performance test**

Use `ab`, `wrk` and other performance pressure test tools to test api.

<br>