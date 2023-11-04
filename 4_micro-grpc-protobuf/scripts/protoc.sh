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
  # get the proto file of the user server
  allProtoFiles=""
  listProtoFiles ${protoBasePath}/user
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
  
  moduleName=$(cat docs/gen.info | head -1 | cut -d , -f 1)
  serverName=$(cat docs/gen.info | head -1 | cut -d , -f 2)
  # Generate 2 files, a logic code template file *.go (default save path in internal/service), a return error code template file *_rpc.go (default save path in internal/ecode)
  protoc --proto_path=. --proto_path=./third_party \
    --go-rpc-tmpl_out=. --go-rpc-tmpl_opt=paths=source_relative \
    --go-rpc-tmpl_opt=moduleName=${moduleName} --go-rpc-tmpl_opt=serverName=${serverName} \
    $specifiedProtoFiles

  checkResult $?

  sponge merge rpc-pb
  checkResult $?

  colorCyan='\e[1;36m'
  highBright='\e[1m'
  markEnd='\e[0m'

  echo ""
  echo -e "${highBright}Tip:${markEnd} execute the command ${colorCyan}make run${markEnd} and then test rpc method is in the file ${colorCyan}internal/service/xxx_client_test.go${markEnd}."
  echo ""
  
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

go mod tidy
checkResult $?

echo "generated code successfully."
echo ""
