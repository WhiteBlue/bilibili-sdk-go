#!/bin/bash

IMAGENAME="bilibili-go"
TAGENAME="whiteblue/bilibili-go"

sudo -l || ( echo "Error: scripts need run with 'sudo'" && exit -1 )


( cid=`sudo docker ps -a |grep $IMAGENAME | awk '{printf $1" "}'`

if [ "$cid" !=  "" ];then

sudo docker rm -f $cid && echo "delete container $cid"

fi )

( iid=`sudo docker images | grep $IMAGENAME | awk '{printf $3" "}'`

if [ "$iid" !=  "" ];then

sudo docker rmi $iid && echo "delete image $iid"

fi )


sudo docker build -t $TAGENAME . || ( echo "Build image failed....." && exit -1 )

echo "build success"

sudo docker run -d --name $IMAGENAME -p 80:8080 $TAGENAME

echo "run container success"

exit 0
