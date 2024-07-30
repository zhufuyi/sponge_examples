
## English | [简体中文](readme-cn.md)

## Transfer service

The transfer service is an example service that demonstrates how to use DTM's service registration and discovery capabilities, as well as how to use DTM's two-stage message transactions.

Test dtm add [dtmdriver-sponge](https://github.com/zhufuyi/dtmdriver-sponge) driver, use service registration and discovery example, the services involved in the example run on the same machine, so the IP address is 127.0.0.1, if the services run on different machines, need to modify the transfer and dtm configuration file to change 127.0.0.1 to the corresponding host IP or domain name.

Key code:

* In [internal/rpcclient/dtmservice.go](internal/rpcclient/dtmservice.go) line 9, import the driver code.
* In [internal/service/transfer.go](internal/service/transfer.go) line 65, commit the transaction.

<br>

### Using etcd as service discovery example

1. Start etcd service.

2. Modify dtm configuration file.

```yaml
MicroService:
   Driver: 'dtm-driver-sponge'
   Target: 'etcd://127.0.0.1:2379/dtmservice'
   EndPoint: 'grpc://127.0.0.1:36790'
```

3. Start dtm service.

```bash
dmt -c conf.yml
```

4. Modify transfer configuration.

Open the configuration file `configs/transfer.yml`, modify the app and grpcClient fields `registryDiscoveryType` value, indicating the use of etcd as service registration and discovery.

```yaml
app:
  registryDiscoveryType: "etcd"       # registry and discovery types: consul, etcd, nacos, if empty, registration and discovery are not used

grpcClient:
  - name: "dtmservice"   # dtm service name, used for service discovery
    registryDiscoveryType: "etcd"    # registration and discovery types: consul, etcd, nacos, if empty, connecting to server using host and port
    host: "127.0.0.1"       # dtm service address, used for direct connection
    port: 36790               # dtm service port

etcd:
  addrs: ["127.0.0.1:2379"]
```

5. Start transfer service

```bash
# Compile and run the service
make run
```

6. Test

Open the generated grpc client test code `internal/service/transfer_client_test.go`, test before filling in the parameters, example

```go
		{
			name: "Transfer",
			fn: func() (interface{}, error) {
				// todo type in the parameters to test
				req := &transferV1.TransferRequest{
					Amount:     100,
					FromUserId: 1,
					ToUserId:   2,
				}
				return cli.Transfer(ctx, req)
			},
			wantErr: false,
		},
```

Open a new terminal, switch to the directory `internal/service`, execute the command to test:

```bash
go test -run Test_service_transfer_methods/Transfer
```

> Note: If using `Goland IDE` to open the code `internal/service/transfer_client_test.go`, click the green button on the left to test.

<br>

### Using consul as service discovery

Same as the 6 steps above, replace `etcd` with `consul`, restart dtm and transfer services.

### Using nacos as service discovery

Same as the 6 steps above, replace `etcd` with `nacos`, restart dtm and transfer services.

Note: The namespaceID in the dtm and transfer configuration files must be the same, the default namespaceID is `public`, if specifying other values, must modify both.
