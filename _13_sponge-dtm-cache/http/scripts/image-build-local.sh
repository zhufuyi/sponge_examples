#!/bin/bash

# build the image for local docker, using the binaries, if you want to reduce the size of the image,
# use upx to compress the binaries before building the image.

serverName="stock"
# image name of the service, prohibit uppercase letters in names.
IMAGE_NAME="eshop/stock"
# Dockerfile file directory
DOCKERFILE_PATH="scripts/build"
DOCKERFILE="${DOCKERFILE_PATH}/Dockerfile"

mv -f cmd/${serverName}/${serverName} ${DOCKERFILE_PATH}/${serverName}

# compressing binary file
#cd ${DOCKERFILE_PATH}
#upx -9 ${serverName}
#cd -

mkdir -p ${DOCKERFILE_PATH}/configs && cp -f configs/${serverName}.yml ${DOCKERFILE_PATH}/configs/
echo "docker build -f ${DOCKERFILE} -t ${IMAGE_NAME}:latest ${DOCKERFILE_PATH}"
docker build -f ${DOCKERFILE} -t ${IMAGE_NAME}:latest ${DOCKERFILE_PATH}


if [ -f "${DOCKERFILE_PATH}/${serverName}" ]; then
    rm -f ${DOCKERFILE_PATH}/${serverName}
fi

if [ -d "${DOCKERFILE_PATH}/configs" ]; then
    rm -rf ${DOCKERFILE_PATH}/configs
fi

# delete none image
noneImages=$(docker images | grep "<none>" | awk '{print $3}')
if [ "X${noneImages}" != "X" ]; then
  docker rmi ${noneImages} > /dev/null
fi
exit 0
