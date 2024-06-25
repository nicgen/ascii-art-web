#!/usr/bin/env bash

image_name="ascii_web"
container_name="web"

sudo docker stop $container_name & pid=$!
wait $pid
echo "docker container $container_name stopped"
sudo docker rm $container_name & pid=$!
wait $pid
echo "docker container $container_name removed"
sudo docker rmi $image_name & pid=$!
wait $pid
echo "docker image $image_name removed"
sudo docker images & pid=$!
wait $pid
echo "the image has been remove"
