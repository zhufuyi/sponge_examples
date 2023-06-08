#!/bin/bash

dockerComposeFilePath="deployments/docker-compose"

mkdir -p ${dockerComposeFilePath}/configs
if [ ! -f "${dockerComposeFilePath}/configs/inventory.yml" ];then
  cp configs/inventory.yml ${dockerComposeFilePath}/configs
fi

# shellcheck disable=SC2164
cd ${dockerComposeFilePath}

if [ "$1"x = "stop"x ] ;then
  docker-compose down
  exit 0
fi

docker-compose up -d

echo "path is 'deployments/docker-compose'"
