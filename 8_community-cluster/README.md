[**community-cluster 中文**](https://juejin.cn/post/7255590272557056059)

<br>

## Community-cluster Introduction

The community-cluster is a microservice cluster composed of two types of services: **gRPC services** and **rpc gateway services**. gRPC services are independent functional implementation modules, while rpc gateway services mainly serve to forward requests to gRPC services and assemble data. The community-cluster service cluster is built using the sponge tool, which automatically separates business logic and non-business logic code when generating gRPC service and rpc gateway service code. The benefit of separating business logic from non-business logic is that it allows developers to focus on core business logic code, greatly reducing the difficulty of building a microservice cluster and reducing the need for manual coding. Click to view the complete [project code](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster) of the microservice cluster, and the framework diagram is shown below:

![community-cluster-frame](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_community-cluster-frame.png)

The code composition structure of gRPC service is based on [grpc](https://github.com/grpc/grpc-go) encapsulation, including rich service governance plug-ins, construction, deployment scripts, and the code composition structure of gRPC service is shown in the following figure:

![micro-rpc-pb-anatomy](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-pb-anatomy.png)
<p align="center">Figure 1 gRPC service code structure diagram</p>

As can be seen from Figure 1, developing a complete microservice focuses on **defining data tables**, **defining api interfaces**, and **writing specific business logic code in template code**. These three nodes have already existed in the monolithic web service [community-single](https://github.com/zhufuyi/sponge_examples/tree/main/7_community-single), and there is no need to rewrite them. Just move these codes over. That is, the egg yolk (core business logic code) remains unchanged, only need to change the eggshell (web framework changed to gRPC framework) and egg white (http handler related code changed to rpc service related code), using tools sponge, it is easy to complete the conversion from web service to gRPC service.

<br>

The code of rpc gateway service is based on [gin](https://github.com/gin-gonic/gin) encapsulation, including rich service governance plug-ins, construction, deployment scripts, and the code composition structure of rpc gateway service is shown in the following figure:

![micro-rpc-gw-pb-anatomy](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_micro-rpc-gw-pb-anatomy.png)
<p align="center">Figure 2 rpc gateway service code structure diagram</p>

As can be seen from Figure 2, developing a complete rpc gateway service focuses on **defining api interfaces** and **writing specific business logic code in template code**. Among them, **defining api interfaces** already exists in monolithic web services [community-single](https://github.com/zhufuyi/sponge_examples/tree/main/7_community-single), no need to rewrite, just copy proto files over.

![community-cluster-gen](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_community-cluster-gen.png)

This is a comparison chart of proto files that monolithic web services and microservice clusters depend on. On the left are proto files that monolithic web services depend on, and all proto files are in the same service. On the right are proto files that microservices depend on, according to each gRPC service depends on its own proto file.

In rpc gateway services, if you need to assemble data from multiple microservices into a new api interface, fill in the description information of this assembled new api interface into community_gw.proto file.

<br>

Below we use tools sponge from 0 to complete microservice cluster process. The development process depends on tools sponge. You need to install sponge first. Click to view [installation instructions](https://github.com/zhufuyi/sponge/blob/main/assets/install-cn.md#%E5%9C%A8linux%E6%88%96macos%E4%B8%8A%E5%AE%89%E8%A3%85sponge).

Create a directory community-cluster and move each independent microservice code under this directory.

<br>

## gRPC service

### Create user, relation, creation services

Enter the sponge UI interface, click on the left menu bar [Protobuf] -> [RPC type] -> [Create rpc project], fill in the parameters, and generate three microservice codes separately.

#### Create user service

This is the proto file [user.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/user/api/user/v1/user.proto) copied from the monolithic service community-single, which is used to quickly generate user (user) service code, as shown below:

![community-rpc-pb-user](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_community-rpc-pb-user.png)

Unzip the code, change the directory name to user, and then move the user directory under the community-cluster directory.

<br>

#### Create relation service

This is the proto file [relation.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/relation/api/relation/v1/relation.proto) copied from the monolithic service community-single, which is used to quickly generate relationship (relation) services, as shown below:

![community-rpc-pb-relation](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_community-rpc-pb-relation.png)

Unzip the code, change the directory name to relation, and then move the relation directory under the community-cluster directory.

<br>

#### Create creation service

These are proto files [post.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/creation/api/creation/v1/post.proto), [comment.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/creation/api/creation/v1/comment.proto), [like.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/creation/api/creation/v1/like.proto), and [collect.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/creation/api/creation/v1/collect.proto) copied from monolithic service community-single to quickly generate creation (creation) services, as shown below:

![community-rpc-pb-creation](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_community-rpc-pb-creation.png)

Unzip the code, change the directory name to creation, and then move the creation directory under the community-cluster directory.

<br>

After simple interface operations, three gRPC services (user, relation, creation) were created. That is to complete each gRPC service's own Figure 1 eggshell part. Next, complete Figure 1 egg white and egg yolk two parts of code.

<br>

### Write business logic code for user, relation, creation services

From Figure 1 above, it can be seen that after sponge stripping, business logic code only includes two parts: egg white and egg yolk. Writing business logic code is carried out around these two parts.

#### Write business logic code for user service

Write business logic code for user service in three steps.

**The first step is to generate template code**. Enter the project user directory, open a terminal and execute the command:

```bash
make proto
```

This command generates service template code, error codes, rpc client test code. These codes correspond to Figure 1 egg white part.

- **Service template code**, in the `internal/service` directory. The file name is consistent with the proto file name and the suffix is `_login.go`. The method functions in the file correspond one-to-one with the rpc method names defined in the proto file. By default, there are simple usage examples under each method function. Just write specific logic code in each method function.
- **Interface error codes**, in the `internal/ecode` directory. The file name is consistent with the proto file name and the suffix is `_rpc.go`. The default error code variables in the file correspond one-to-one with the rpc method names defined in the proto file. Add or change business-related error codes here. Note that error codes cannot be repeated, otherwise panic will be triggered.
- **rpc client test code**, in `internal/service` directory. The file name is consistent with the proto file name and suffix is `_client_test.go`. The method functions in this file correspond one-to-one with rpc method names defined in proto files. Fill in parameters for each rpc method.

**The second step is to migrate dao code**. Copy user-starting code files under `internal/model`, `internal/cache`, `internal/dao`, and `internal/ecode` directories of monolithic web service community-single to user service directory. The copied directories and filenames remain unchanged. These codes correspond to Figure 1 egg white part.

**The third step is to migrate specific logic code**. Copy specific logic code under each method function of monolithic web service community-single's `internal/handler/user_logic.go` into same-named functions of user service's `internal/service/user_logic.go`. These codes are the part of writing business logic code in Figure 1 egg yolk.

<br>

#### Write business logic code for relation service

Write business logic code for relation service in three steps, referring to the three steps of the above user service.

<br>

#### Write business logic code for creation service

Write business logic code for creation service in three steps, referring to the three steps of the above user service.

<br>

### Test rpc methods of user, relation, creation services

#### Test rpc methods of user service

After writing the business logic code, start the service to test the rpc method. Before starting the service for the first time, open the configuration file (`user/configs/user.yml`) to set mysql and redis addresses, set grpc and grpcClient related parameters, and then execute the command to compile and start the service:

```bash
# Compile and run the service
make run
```

Open user service code in goland IDE, enter `user/internal/service` directory, find code files with suffix `_client_test.go`, fill in parameters for each rpc method and test.

<br>

#### Test rpc methods of relation service

To test rpc methods of relation service, please refer to testing rpc methods of user service above.

<br>

#### Test rpc methods of creation service

To test rpc methods of creation service, please refer to testing rpc methods of user service above.

<br>

## rpc gateway service

After completing the user, relation, and creation services, the rpc gateway service community_gw needs to be completed next. Community_gw serves as the unified entrance for user, relation, and creation services.

### Create community_gw service

Enter the sponge UI interface, click on the left menu bar [Protobuf] -> [Web type] -> [Create rpc gateway project], fill in some parameters to generate rpc gateway service code.

The proto files([user_gw.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/api/community_gw/v1/user_gw.proto)、[relation_gw.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/api/community_gw/v1/relation_gw.proto)、[post_gw.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/api/community_gw/v1/post_gw.proto)、[comment_gw.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/api/community_gw/v1/comment_gw.proto)、[like_gw.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/api/community_gw/v1/like_gw.proto)、[collect_gw.proto](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/api/community_gw/v1/collect_gw.proto)) copied from the monolithic service community-single to quickly create the rpc gateway service community_gw, as shown below:

![community-rpc-gw-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_community-rpc-gw-pb.png)

Unzip the code and change the directory name to community_gw.

<br>

Because community_gw service serves as the request entrance and communicates with user, relation, and creation services using rpc method, it is necessary to generate code that connects to user, relation, and creation services. Enter the sponge UI interface, click on the left menu bar [Public] -> [Generate rpc service connection code], fill in some parameters to generate rpc service connection code, as shown below:

![community-rpc-conn](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_community-rpc-conn.png)

Unzip the code and move the internal directory to the community_gw service directory.

At the same time, copy the proto files of user, relation, and creation services to the api directory of community_gw as shown in the list below. The proto files in v1 directory of community_gw are used to define http api interface information. It is recommended to uniformly agree on suffix `_gw.proto`.

```
.
├── community_gw
│   └── v1
│       ├── collect_gw.proto
│       ├── comment_gw.proto
│       ├── like_gw.proto
│       ├── post_gw.proto
│       ├── relation_gw.proto
│       └── user_gw.proto
├── creation
│   └── v1
│       ├── collect.proto
│       ├── comment.proto
│       ├── like.proto
│       └── post.proto
├── relation
│   └── v1
│       └── relation.proto
└── user
    └── v1
        └── user.proto
```

Through simple operations, we have completed creating rpc gateway service community_gw.

<br>

### Write business logic code for community_gw service

From Figure 2 above, it can be seen that after sponge stripping, business logic code only includes two parts: egg white and egg yolk. Writing business logic code is carried out around these two parts.

#### Write business logic code related to proto files

Enter project community-gw directory, open terminal and execute command:

```bash
make proto
```

This command generates four parts of code based on proto files in `community_gw/api/community_gw/v1` directory: service template code, registration route code, api interface error code, swagger document. That is Figure 2 egg white part.

(1) **Generated service template code**, in `community_gw/internal/service` directory. The file name is consistent with proto file name and suffix is `_logic.go`. The names are:

[collect_gw_logic.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/service/collect_gw_logic.go), [comment_gw_logic.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/service/comment_gw_logic.go), [like_gw_logic.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/service/like_gw_logic.go), [post_gw_logic.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/service/post_gw_logic.go), [relation_gw_logic.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/service/relation_gw_logic.go), [user_gw_logic.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/service/user_gw_logic.go)

In these files, method functions correspond one-to-one with rpc method names defined in proto files. There are default usage examples under each method function. Just make simple adjustments to call rpc methods of user, relation, and creation services. The above files are codes after writing specific logic.

<br>

(2) **Generated registration route code**, in `community_gw/internal/routers` directory. The file name is consistent with proto file name and suffix is `_service.pb.go`. The names are:

[collect_gw_service.pb.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/routers/collect_gw_service.pb.go), [comment_gw_service.pb.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/routers/comment_gw_service.pb.go), [like_gw_service.pb.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/routers/like_gw_service.pb.go), [post_gw_service.pb.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/routers/post_gw_service.pb.go), [relation_gw_service.pb.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/routers/relation_gw_service.pb.go), [user_gw_service.pb.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/routers/user_gw_service.pb.go)

In these files, set middleware for api interface such as jwt authentication. Each interface already has middleware template code. Just uncommenting code can make middleware effective. Support setting gin middleware for route groups and individual routes.

<br>

(3) **Generated interface error codes**, in `community_gw/internal/ecode` directory. The file name is consistent with proto file name and suffix is `_rpc.go`. The names are:

[collect_gw_rpc.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/ecode/collect_gw_rpc.go), [comment_gw_rpc.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/ecode/comment_gw_rpc.go), [like_gw_rpc.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/ecode/like_gw_rpc.go), [post_gw_rpc.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/ecode/post_gw_rpc.go), [relation_gw_rpc.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/ecode/relation_gw_rpc.go), [user_gw_rpc.go](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/internal/ecode/user_gw_rpc.go)

In these files, default error code variables correspond one-to-one with rpc method names defined in proto files. Add or change business-related error codes here. Note that error codes cannot be repeated otherwise panic will be triggered.

Note: If the called rpc method itself contains an error code, it can be returned directly.

<br>

(4) **Generated swagger document**, in `community_gw/docs` directory. The name is [apis.swagger.json](https://github.com/zhufuyi/sponge_examples/blob/main/8_community-cluster/community_gw/docs/apis.swagger.json)

<br>

If you add or change api interfaces in proto files, you need to execute command `make proto` again to update codes. You will find that there are date-time-suffixed code files under `community_gw/internal/handler`, `community_gw/internal/routers`, `community_gw/internal/ecode` directories. Open file and copy new or modified part of codes into same-named file. After copying new codes execute command `make clean` to clear these date-suffix files.

The code generated by `make proto` command is a bridge connecting web framework code and core business logic code which is egg white part. The advantage of this layered generation of codes is that it reduces writing codes.

<br>

### Test api interface

After writing the business logic code, start the service to test the api interface. Before starting the service for the first time, open the configuration file (`community_gw/configs/community_gw.yml`) to set the connection rpc service configuration information as shown below:

```yaml
# grpc client settings, support for setting up multiple rpc clients
grpcClient:
  - name: "user"
    host: "127.0.0.1"
    port: 18282
    registryDiscoveryType: ""
    enableLoadBalance: false
  - name: "relation"
    host: "127.0.0.1"
    port: 28282
    registryDiscoveryType: ""
    enableLoadBalance: false
  - name: "creation"
    host: "127.0.0.1"
    port: 38282
    registryDiscoveryType: ""
    enableLoadBalance: false
```

Execute the command to compile and start the service:

```bash
# Compile and run the service
make run
```

Visit `http://localhost:8080/apis/swagger/index.htm` in your browser to enter the swagger interface, as shown below:

![community-gw-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/community-gw-swagger.png)

From the figure, it can be seen that some api interfaces have a lock mark on the right, indicating that the request header will carry authentication information Authorization. Whether the server receives a request for authentication is determined by the server. If the server needs to authenticate, it can be set in `community_gw/internal/routers` directory suffix file `_service.pb.go`, that is, uncommenting authentication code to make api interface authentication middleware effective.

<br>

## Service Governance

gRPC services (user, relation, creation) and rpc gateway service (community-gw) all contain rich service governance plug-ins (logging, flow control, circuit breaking, link tracking, service registration and discovery, metric collection, performance analysis, resource statistics, configuration center), some service governance plug-ins are closed by default and can be turned on as needed.

In addition to the governance plug-ins provided by the service itself, you can also use your own service governance plug-ins. To add your own service governance plug-ins:

- For microservices (user, relation, creation), add your own plug-ins in the code file `service name/internal/server/grpc.go`. If your service governance plug-in (interceptor) is of unary type, add it to the `unaryServerOptions` function. If your service governance plug-in (interceptor) is of stream type, add it to the `streamServerOptions` function.
- For rpc gateway service community-gw, add your own plug-ins (gin middleware) in the code file `community-gw/internal/routers/routers.go`.

<br>

Below are the default service governance plug-in opening and setting instructions. Set uniformly in each service's configuration file `service name/configs/service name.yml`.

**Logging**

The logging plug-in (zap) is turned on by default and outputs to the terminal by default. The default output log format is console. You can set the output format to json and set the log to be saved to a specified file. The log file is cut and retained for a period of time.

Set in the `logger` field in the configuration file:

```yaml
# logger settings
logger:
  level: "info"             # Output log level debug, info, warn, error, default is debug
  format: "console"     # Output format, console or json, default is console
  isSave: false           # false: output to terminal, true: output to file, default is false
  logFileConfig:          # valid when isSave=true
    filename: "out.log"            # File name, default value out.log
    maxSize: 20                     # Maximum file size (MB), default value 10MB
    maxBackups: 50               # Maximum number of old files retained, default value 100
    maxAge: 15                     # Maximum number of days old files are retained, default value 30 days
    isCompression: true          # Whether to compress/archive old files, default value false
```

<br>

**Flow Control**

The flow control plug-in is turned off by default and adapts itself without setting other parameters.

Set in the `enableLimit` field in the configuration file:

```yaml
  enableLimit: false    # Whether to enable flow control (adaptive), true: turn on, false: turn off
```

<br>

**Circuit Breaking**

The circuit breaking plug-in is turned off by default and adapts itself. It supports custom error codes (default 500 and 503) to trigger circuit breaking. Set in `internal/routers/routers.go`.

Set in the `enableCircuitBreaker` field in the configuration file:

```yaml
  enableCircuitBreaker: false    # Whether to enable circuit breaking (adaptive), true: turn on, false: turn off
```

<br>

**Link Tracking**

The link tracking plug-in is turned off by default and depends on jaeger service.

Set in the `enableTrace` field in the configuration file:

```yaml
  enableTrace: false    # Whether to enable tracking , true : enabled , false : closed , if it is true , must set jaeger configuration .
  tracingSamplingRate : 1.0      # Link tracking sampling rate , range 0 ~1.0 floating point number , 0 means no sampling , 1.0 means sampling all links


# jaeger settings
jaeger:
  agentHost : "192.168.3.37"
  agentPort : 6831
```

View link tracking information on jaeger interface [document description](https://go-sponge.com/zh-cn/service-governance?id=%e9%93%be%e8%b7%af%e8%b7%9f%e8%b8%aa).

<br>

**Service Registration and Discovery**

The service registration and discovery plug-in is turned off by default and supports three types of consul, etcd, and nacos.

Set in the `registryDiscoveryType` field in the configuration file:

```yaml
  registryDiscoveryType : ""    # Registration and discovery type : consul , etcd , nacos , if it is empty , it means that service registration and discovery are turned off.


# Set parameters according to registryDiscoveryType value . For example , use consul as service discovery , just set consul .
# consul settings
consul:
  addr : "192.168.3.37:8500"

# etcd settings
etcd:
  addrs : ["192.168.3.37:2379"]

# nacos settings
nacosRd:
  ipAddr : "192.168.3.37"
  port : 8848
  namespaceID : "3454d2b5-2455-4d0e-bf6d-e033b086bb4c" # namespace id
```

<br>

**Metric Collection**

The metric collection function is turned on by default and provides data for prometheus to collect. The default route is `/metrics`.

Set in the `enableMetrics` field in the configuration file:

```yaml
  enableMetrics: true    # Whether to enable metric collection, true: enabled, false: closed
```

Use prometheus and grafana to collect metrics and monitor services [document description](https://go-sponge.com/zh-cn/service-governance?id=%e7%9b%91%e6%8e%a7).

<br>

**Performance Analysis**

The performance analysis plug-in is turned off by default. The default route for collecting profile is `/debug/pprof`. In addition to supporting the default profile analysis provided by go language itself, it also supports io analysis. The route is `/debug/pprof/profile-io`.

Set in the `enableHTTPProfile` field in the configuration file:

```yaml
  enableHTTPProfile: false    # Whether to enable performance analysis, true: enabled, false: closed
```

The method of collecting profile through routing for performance analysis is usually used during development or testing. If it is turned on online, there will be a little performance loss because the program records profile-related information at regular intervals in the background. The service generated by sponge itself has made some improvements. Normally, it stops collecting profile. When the user actively triggers a system signal, it starts and stops collecting profile. Collecting profile is saved to `/tmp/service name_profile directory`, and the default collection is 60 seconds. After 60 seconds, it automatically stops collecting profile. If you only want to collect for 30 seconds, send the first signal to start collecting, and send the second signal about 30 seconds later to stop collecting profile, similar to a switch.

This is the operation step of collecting profile:

```yaml
# View service pid by name
ps aux | grep service name

# Send signal to service
kill -trap pid value
```

Note: Only supports linux and darwin systems.

<br>

**Resource Statistics**

The resource statistics plug-in is turned on by default and outputs to the log once every minute by default. Resource statistics include cpu and memory-related data for both system and service itself.

Resource statistics also include automatic triggering of profile collection function. When the average value of CPU or memory of this service is continuously counted for 3 times, and the average value of CPU or memory occupies more than 80% of system resources, it will automatically trigger collection of profile. The default collection is 60 seconds. Collecting profile is saved to `/tmp/service name_profile directory`, thus realizing adaptive collection of profile, which is another improvement over manually sending system signals to collect profile.

Set in the `enableHTTPProfile` field in the configuration file:

```yaml
  enableStat: true    # Whether to enable resource statistics, true: enabled, false: closed
```

<br>

**Configuration Center**

Currently supports nacos as configuration center. Configuration center file `configs/user_cc.yml`, configuration content is as follows:

```yaml
# nacos settings
nacos:
  ipAddr: "192.168.3.37"    # Service address
  port: 8848                      # Listening port
  scheme: "http"                # Supports http and https
  contextPath: "/nacos"       # Path
  namespaceID: "3454d2b5-2455-4d0e-bf6d-e033b086bb4c" # namespace id
  group: "dev"                    # Group name: dev, prod, test
  dataID: "community.yml"  # Configuration file id
  format: "yaml"                 # Configuration file type: json,yaml,toml
```

And copy the service's configuration file `configs/user.yml` to nacos interface for configuration. To use nacos as a configuration center, start service command needs to specify configuration center file, command as follows:

```bash
./user -c configs/user_cc.yml -enable-cc
```

Use nacos as configuration center [document description](https://go-sponge.com/zh-cn/service-governance?id=%e9%85%8d%e7%bd%ae%e4%b8%ad%e5%bf%83).

<br>

## Service Stress Testing

Some tools used when stress testing services:

- http stress testing tool [wrk](https://github.com/wg/wrk) or [go-stress-testing](https://github.com/link1st/go-stress-testing).
- The service opens the metric collection function and uses prometheus to collect service metrics and system metrics for monitoring.
- The service's own adaptive profile collection function.

<br>

Stress test indicators:

- **Concurrency**: Gradually increase the number of concurrent users to find the maximum concurrency of the service and determine the maximum number of users the service can support.
- **Response time**: Pay attention to the average response time and response time distribution of the service when the number of concurrent users increases. Ensure that even under high concurrency, the response time is within an acceptable range.
- **Error rate**: Observe the probability of errors or exceptions in the service when concurrency increases. Use stress testing tools for long-term concurrent testing and count the number and type of errors at each concurrency level.
- **Throughput**: Find the maximum throughput of the service and determine the maximum request volume that the service can support under high concurrency. This requires continuous increase in concurrency until the throughput saturation point is found.
- **Resource utilization**: Pay attention to the utilization rate of resources such as CPU, memory, disk I/O, and network when concurrency increases, and find the resource bottleneck of the service.
- **Bottleneck detection**: By observing performance indicators and resource utilization rates of services under high concurrency, find hardware or software bottlenecks in systems and services for optimization.
- **Stability**: Long-term high-concurrency operation can detect potential problems in services, such as memory leaks, connection leaks, etc., to ensure stable operation of services. This requires long-term concurrent stress testing and observation of service operation indicators.

The purpose of stress testing a service is to evaluate its performance, determine its maximum concurrency and throughput, discover current bottlenecks, and detect the stability of service operation for optimization or capacity planning.

<br>

## Summary

This article introduces the specific practice process of splitting monolithic service community-single into microservice cluster community-cluster. The microservice cluster includes three independent services: user service (user), relationship service (relation), and content creation service (creation). A microservice entrance gateway service (community_gw), these service codes (the eggshell and egg white parts in Figures 1 and 2) are all generated by tool sponge, and core business logic code can be directly manually seamlessly transplanted.

It is easy to build a microservice cluster using tool sponge. The advantages of microservice cluster are:

- High performance: High-performance communication protocol based on Protobuf, with both high-concurrency processing and low-latency characteristics.
- Scalability: Rich plug-in and component mechanisms. Developers can customize and expand framework functions according to actual needs.
- High reliability: Provides functions such as service registration and discovery, flow control, circuit breaking, link tracking, monitoring alarm, etc., which improves the reliability of microservices.
