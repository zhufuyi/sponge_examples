[web-gin-protobuf 中文说明](https://www.bilibili.com/read/cv23040234)

<br>

### Quickly Create a Web Project

Prepare a proto file before creating a web service. The proto file must contain **route description information** and **swagger description information**.

Enter the UI interface of sponge, click 【Protobuf】→【Create Web Project】in the left menu bar, and fill in some parameters to generate the web service project code.

The web framework uses [gin](https://github.com/gin-gonic/gin). It also includes swagger documents, common service governance function codes, and build and deployment scripts. You can choose which database to use.

![web-http-pb](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_web-http-pb.png)

Change to the web directory and execute the command:

```bash
# Generate API code, generate registered routing code, and generate swagger docs 
make proto

# Open internal/handler/user.go, which is the generated API code. There is 
# a line of panic code prompting to fill in the business logic code. Fill in the business logic here.
# The registration route of the API, the go structure code of the input parameter and the returned result, 
# the swagger document, and the definition error code have all been generated. 
# Just fill in the business logic code. 

# compile and start web services
make run 
```

Open [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html) in your browser to test the API.

![web-http-pb-swagger2](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_web-http-pb-swagger2.png)

<br>

**More detailed development web documentation** https://go-sponge.com/web-development-protobuf

<br>
