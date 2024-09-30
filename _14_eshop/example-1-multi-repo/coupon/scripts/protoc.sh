#!/bin/bash

# the directory where the proto files are located
protoBasePath="api"
allProtoFiles=""

specifiedProtoFilePath=$1
specifiedProtoFilePaths=""

function checkResult() {
    result=$1
    if [ ${result} -ne 0 ]; then
        exit ${result}
    fi
}

# get specified proto files, if empty, return 0 else return 1
function getSpecifiedProtoFiles() {
  if [ "$specifiedProtoFilePath"x = x ];then
    return 0
  fi

  specifiedProtoFilePaths=${specifiedProtoFilePath//,/ }

  for v in $specifiedProtoFilePaths; do
    if [ ! -f "$v" ];then
      echo "Error: not found specified proto file $v"
	    echo "example: make proto FILES=api/user/v1/user.proto,api/types/types.proto"
      checkResult 1
    fi
  done

  return 1
}

# add the import of useless packages from the generated *.pb.go code here
function deleteUnusedPkg() {
  file=$1
  osType=$(uname -s)
  if [ "${osType}"x = "Darwin"x ];then
    sed -i '' 's#_ \"github.com/envoyproxy/protoc-gen-validate/validate\"##g' ${file}
    sed -i '' 's#_ \"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options\"##g' ${file}
    sed -i '' 's#_ \"github.com/srikrsna/protoc-gen-gotag/tagger\"##g' ${file}
    sed -i '' 's#_ \"google.golang.org/genproto/googleapis/api/annotations\"##g' ${file}
  else
    sed -i "s#_ \"github.com/envoyproxy/protoc-gen-validate/validate\"##g" ${file}
    sed -i "s#_ \"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options\"##g" ${file}
    sed -i "s#_ \"github.com/srikrsna/protoc-gen-gotag/tagger\"##g" ${file}
    sed -i "s#_ \"google.golang.org/genproto/googleapis/api/annotations\"##g" ${file}
  fi
  checkResult $?
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
  getSpecifiedProtoFiles
  if [ $? -eq 0 ]; then
    listProtoFiles $protoBasePath
  else
    allProtoFiles=$specifiedProtoFilePaths
  fi

  if [ "$allProtoFiles"x = x ];then
    echo "Error: not found proto file in path $protoBasePath"
    exit 1
  fi
  echo "generate *pb.go by proto files: $allProtoFiles"
  echo ""

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
  # get the proto file of the coupon server
  allProtoFiles=""
  listProtoFiles ${protoBasePath}/coupon
  cd ..
  specifiedProtoFiles=""
  getSpecifiedProtoFiles
  if [ $? -eq 0 ]; then
    specifiedProtoFiles=$allProtoFiles
  else
	  for v1 in $specifiedProtoFilePaths; do
      for v2 in $allProtoFiles; do
        if [ "$v1"x = "$v2"x ];then
          specifiedProtoFiles="$specifiedProtoFiles $v1"
        fi
      done
	  done
  fi

  if [ "$specifiedProtoFiles"x = x ];then
    return
  fi
  echo "generate template code by proto files: $specifiedProtoFiles"
  
  # generate the swagger document and merge all files into docs/apis.swagger.json
  protoc --proto_path=. --proto_path=./third_party \
    --openapiv2_out=. --openapiv2_opt=logtostderr=true --openapiv2_opt=allow_merge=true --openapiv2_opt=merge_file_name=docs/apis.json \
    $specifiedProtoFiles

  checkResult $?

  sponge web swagger --file=docs/apis.swagger.json
  checkResult $?

  moduleName=$(cat docs/gen.info | head -1 | cut -d , -f 1)
  serverName=$(cat docs/gen.info | head -1 | cut -d , -f 2)
  suitedMonoRepo=$(cat docs/gen.info | head -1 | cut -d , -f 3)

  protoc --proto_path=. --proto_path=./third_party \
    --go-rpc-tmpl_out=. --go-rpc-tmpl_opt=paths=source_relative \
    --go-rpc-tmpl_opt=moduleName=${moduleName} --go-rpc-tmpl_opt=serverName=${serverName} --go-rpc-tmpl_opt=suitedMonoRepo=${suitedMonoRepo} \
    $specifiedProtoFiles

  checkResult $?

  sponge merge rpc-pb
  checkResult $?

  protoc --proto_path=. --proto_path=./third_party \
    --go-gin_out=. --go-gin_opt=paths=source_relative --go-gin_opt=plugin=mix \
    --go-gin_opt=moduleName=${moduleName} --go-gin_opt=serverName=${serverName} --go-gin_opt=suitedMonoRepo=${suitedMonoRepo} \
    $specifiedProtoFiles

  checkResult $?

  sponge merge http-pb
  checkResult $?

  colorCyan='\033[1;36m'
  highBright='\033[1m'
  markEnd='\033[0m'

  echo ""
  echo -e "${highBright}Tip:${markEnd} execute the command ${colorCyan}make run${markEnd} and then"
  echo -e "    1. test http api in your browser ${colorCyan}http://localhost:8080/apis/swagger/index.html${markEnd}"
  echo -e "    2. test grpc api in the file ${colorCyan}internal/service/xxx_client_test.go${markEnd}"
  echo ""
  

  if [ "$suitedMonoRepo" == "true" ]; then
    sponge patch adapt-mono-repo
  fi
}

# generate pb.go by all proto files
generateByAllProto

# generate pb.go by specified proto files
generateBySpecifiedProto

# delete unused packages in pb.go
handlePbGoFiles $protoBasePath

# delete json tag omitempty
sponge patch del-omitempty --dir=$protoBasePath --suffix-name=pb.go > /dev/null

# modify duplicate numbers and error codes
sponge patch modify-dup-num --dir=internal/ecode
sponge patch modify-dup-err-code --dir=internal/ecode

echo "generated code successfully."
echo ""
