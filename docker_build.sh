#!/usr/bin/env bash

container_name="web"
image_name="ascii_web"

sudo echo "Let's build the docker image..."
sudo docker image build -f Dockerfile -t $image_name . & pid=$!
wait $pid
echo "Docker image is built"
sudo docker images & pid=$!
wait $pid
sudo docker container run -p 0:8080 --detach --name $container_name $image_name & pid=$!
# with port 0, you ask the kernel to handle the attribution of a free port
wait $pid
container_port=$(sudo docker container port web | tail -n 1 | grep --only-matching -P ":(\d+)")
echo "Container $container_name from $image_name image is now started on port$container_port"
# sudo docker ps -a & pid=$!
# wait $pid
xdg-open "http://localhost$container_port" > /dev/null & disown
# [source](https://askubuntu.com/a/682547)
# `>/dev/null 2>&1` will prevent messages from the browser to be outputted to the terminal's window; & will put the process into the background
# `disown` will remove the job / process from the job list, preventing a SIGHUP signal to be propagated to it. 
# sudo docker exec -it $container_name bash
# sudo docker exec -it web /bin/bash