## Transfer 服务

transfer服务是一个示例服务，演示了如何使用dtm的服务注册与发现功能，以及如何使用dtm的二阶段消息事务。

在dtm添加[dtmdriver-sponge](https://github.com/zhufuyi/dtmdriver-sponge)驱动后，使用服务注册与发现示例，示例中所涉及的服务运行在同一个机器，因此ip地址都是127.0.0.1，如果有服务运行在不同机器，需要在transfer和dtm配置文件把127.0.0.1修改为对应的宿主机ip或域名。

关键代码：

- 在 [internal/rpcclient/dtmservice.go](internal/rpcclient/dtmservice.go) 第9行代码导入驱动代码。
- 在 [internal/service/transfer.go](internal/service/transfer.go) 第65行代码提交二阶段消息事务。

<br>

### 使用etcd作为服务发现示例

1. 启动etcd服务。

2. 修改dtm配置文件。

```yaml
MicroService:
   Driver: 'dtm-driver-sponge'
   Target: 'etcd://127.0.0.1:2379/dtmservice'
   EndPoint: 'grpc://127.0.0.1:36790'
```

3. 启动dtm服务。

```bash
dmt -c conf.yml
```

4. 修改transfer配置。

打开配置文件 `configs/transfer.yml`，修改app和grpcClient下字段`registryDiscoveryType`值，表示使用etcd作为服务注册与发现。

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

5. 启动 transfer 服务

```bash
# 编译和运行服务
make run
```

6. 测试

打开生成的grpc客户端测试代码 `internal/service/transfer_client_test.go`，测试前填写参数，示例

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

打开一个新终端，切换到目录`internal/service`，执行命令测试：

```bash
go test -run Test_service_transfer_methods/Transfer
```

> Note：如果使用`Goland IDE`打开代码`internal/service/transfer_client_test.go`，直接点击左边绿色按钮测试。


### 使用consul作为服务发现

与上面**使用etcd作为服务发现**的6个操作步骤一样，把`etcd`改为`consul`，重启dtm和transfer服务。

### 使用nacos作为服务发现

与上面**使用etcd作为服务发现**的6个操作步骤一样，把`etcd`改为`nacos`，重启dtm和transfer服务。

注：在dmt和transfer的配置文件的namespaceID必须一样，默认namespaceID是`public`，如果指定其他值，必须同时修改。
