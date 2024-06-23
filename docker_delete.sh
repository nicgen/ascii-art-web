#!/usr/bin/env bash

container_name="web"
image_name="ascii_web"

sudo docker stop $container_name & pid=$!
wait $pid
echo "docker container $container_name stopped"
sudo docker rm $container_name & pid=$!
wait $pid
echo $pid
echo "docker container $container_name removed"
sudo docker rmi $image_name & pid=$!
wait $pid
echo "docker image $image_name removed"
sudo docker images & pid=$!
wait $pid
echo "the image has been remove"

# sudo docker ps -a 
# if $?; then
#   echo "docker image removed"
# fi
