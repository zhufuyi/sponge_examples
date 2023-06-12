#!/bin/bash

dockerComposeFilePath="deployments/docker-compose"

mkdir -p ${dockerComposeFilePath}/configs
if [ ! -f "${dockerComposeFilePath}/configs/community_gw.yml" ];then
  cp configs/community_gw.yml ${dockerComposeFilePath}/configs
fi

# shellcheck disable=SC2164
cd ${dockerComposeFilePath}

if [ "$1"x = "stop"x ] ;then
  docker-compose down
  exit 0
fi

docker-compose up -d

echo "path is 'deployments/docker-compose'"
