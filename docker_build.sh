#!/usr/bin/env bash

image_name="ascii_web"
container_name="web"

sudo echo "Let's build the docker image..."
# if sudo docker image build -f Dockerfile -t $image_name . & pid=$!; wait $pid;
if sudo docker image build -f Dockerfile -t $image_name . ;
then
  echo "Docker image is built"
else
  echo "Docker image failed"
fi
sudo docker images & pid=$!
wait $pid
if sudo docker container run -p 0:8080 --detach --name $container_name $image_name & pid=$!;wait $pid;
then
  container_port=$(sudo docker container port $container_name | tail -n 1 | grep --only-matching -P ":(\d+)")
  echo "Container $container_name from $image_name image is now started on port$container_port"
  sudo docker exec -it $container_name bash
else
  echo "container mount failed"
fi
# with port 0, you ask the kernel to handle the attribution of a free port

# xdg-open "http://localhost$container_port" < /dev/null &>/dev/null & disown
# problem with zombie process log
# [source](https://askubuntu.com/a/682547)
# `>/dev/null 2>&1` will prevent messages from the browser to be outputted to the terminal's window; & will put the process into the background
# `disown` will remove the job / process from the job list, preventing a SIGHUP signal to be propagated to it. 
