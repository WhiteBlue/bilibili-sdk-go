#!/bin/bash


cid=`docker ps -a |grep bilibili-go | awk '{printf $1" "}'`

if [ "$cid" !=  "" ];then

docker rm -f $cid

echo "delete container $cid"

fi

iid=`docker images | grep bilibili-go | awk '{printf $3" "}'`

if [ "$iid" !=  "" ];then

docker rmi $iid

echo "delete image $iid"

fi

echo "pull image from daocloud"

docker pull daocloud.io/whiteblue/bilibili-go

niid=`docker images | grep bilibili-go | awk '{printf $3" "}'`

if [ "$niid" =  "" ];then

echo 'pull image failed....'

return 1

fi

docker run -d --name bilibili-go -p 80:8080 $niid

echo "start container success"