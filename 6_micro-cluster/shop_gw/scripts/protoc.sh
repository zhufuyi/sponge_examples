#!/bin/bash

# the directory where the proto files are located
protoBasePath="api"
allProtoFiles=""

function checkResult() {
    result=$1
    if [ ${result} -ne 0 ]; then
        exit ${result}
    fi
}

# add the import of useless packages from the generated *.pb.go code here
function deleteUnusedPkg() {
  file=$1
  sed -i "s#_ \"github.com/envoyproxy/protoc-gen-validate/validate\"##g" ${file}
  sed -i "s#_ \"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options\"##g" ${file}
  sed -i "s#_ \"github.com/srikrsna/protoc-gen-gotag/tagger\"##g" ${file}
  sed -i "s#_ \"google.golang.org/genproto/googleapis/api/annotations\"##g" ${file}
}

function listProtoFiles(){
    cd $1
    items=$(ls)

    for item in $items; do
        if [ -d "$item" ]; then
            listProtoFiles $item
        else
            if [ "${item#*.}"x = "proto"x ];then
              file=$(pwd)/${item}
              protoFile="${protoBasePath}${file#*${protoBasePath}}"
              allProtoFiles="${allProtoFiles} ${protoFile}"
            fi
        fi
    done
    cd ..
}

function handlePbGoFiles(){
    cd $1
    items=$(ls)

    for item in $items; do
        if [ -d "$item" ]; then
            handlePbGoFiles $item
        else
            if [ "${item#*.}"x = "pb.go"x ];then
              deleteUnusedPkg $item
            fi
        fi
    done
    cd ..
}

function generateByAllProto(){
  # get all proto file paths
  listProtoFiles $protoBasePath
  if [ "$allProtoFiles"x = x ];then
    echo "Error: not found protobuf file in path $protoBasePath"
    exit 1
  fi

  # generate files *_pb.go
  protoc --proto_path=. --proto_path=./third_party \
    --go_out=. --go_opt=paths=source_relative \
    $allProtoFiles

  checkResult $?

  # generate files *_grpc_pb.go
  protoc --proto_path=. --proto_path=./third_party \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    $allProtoFiles

  checkResult $?


  # generate the file *_pb.validate.go
  protoc --proto_path=. --proto_path=./third_party \
    --validate_out=lang=go:. --validate_opt=paths=source_relative \
    $allProtoFiles

  checkResult $?

  # embed the tag field into *_pb.go
  protoc --proto_path=. --proto_path=./third_party \
    --gotag_out=:. --gotag_opt=paths=source_relative \
    $allProtoFiles

  checkResult $?
}

function generateBySpecifiedProto(){
  # get the proto file of the shop_gw server
  allProtoFiles=""
  listProtoFiles ${protoBasePath}/shop_gw
  cd ..
  specifiedProtoFiles=$allProtoFiles

  # Generate the swagger document and merge all files into docs/apis.swagger.json
  protoc --proto_path=. --proto_path=./third_party \
    --openapiv2_out=. --openapiv2_opt=logtostderr=true --openapiv2_opt=allow_merge=true --openapiv2_opt=merge_file_name=docs/apis.json \
    $specifiedProtoFiles

  checkResult $?

  sponge micro swagger --file=docs/apis.swagger.json > /dev/null
  checkResult $?

  echo ""
  echo "run server and see docs by http://localhost:8080/apis/swagger/index.html"
  echo ""

  moduleName=$(cat docs/gen.info | head -1 | cut -d , -f 1)
  serverName=$(cat docs/gen.info | head -1 | cut -d , -f 2)
  # A total of 4 files are generated, namely the registration route file _*router.pb.go (saved in the same directory as the protobuf file),
  # the injection route file *_service.pb.go (default save path in internal/routers), the logical code template file*_logic.go (saved in internal/service by default),
  # return error code template file*_rpc.go (saved in internal/ecode by default)
  protoc --proto_path=. --proto_path=./third_party \
    --go-gin_out=. --go-gin_opt=paths=source_relative --go-gin_opt=plugin=service \
    --go-gin_opt=moduleName=${moduleName} --go-gin_opt=serverName=${serverName} \
    $specifiedProtoFiles

  checkResult $?

}

# generate pb.go by all proto files
generateByAllProto

# generate pb.go by specified proto files
generateBySpecifiedProto

# delete unused packages in pb.go
handlePbGoFiles $protoBasePath

# delete json tag omitempty
sponge del-omitempty --dir=$protoBasePath --suffix-name=pb.go > /dev/null
checkResult $?

go mod tidy
checkResult $?

echo "execute protoc command successfully."
