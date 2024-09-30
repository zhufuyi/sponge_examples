#!/bin/bash

operate=$1

string="coupon eshop_gw flashSale order pay product stock user"
array=($string)

function checkResult() {
    result=$1
    if [ ${result} -ne 0 ]; then
        exit ${result}
    fi
}

function clean() {
  for serverName in "${array[@]}"
  do
      cd $serverName
      make run-nohup
      checkResult $?
      echo "Server $serverName started"
      cd ..
  done
}

function startServer() {
  for serverName in "${array[@]}"
  do
      cd $serverName
      make run-nohup
      checkResult $?
      echo "Server $serverName started"
      cd ..
  done
}

function stopServer() {
  for serverName in "${array[@]}"
  do
      cd $serverName
      make run-nohup CMD=stop
      checkResult $?
      echo "Server $serverName stopped"
      cd ..
  done
}

if [ "$operate" = "stop" ];then
  stopServer
elif [ "$operate" = "clean" ];then
  clean
else
  startServer
fi
