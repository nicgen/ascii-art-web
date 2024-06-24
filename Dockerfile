# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu

FROM golang:1.22-alpine

# create a working directory inside the image

WORKDIR /app

# copy directory files i.e all files ending with .go

# copy all files/folder into /app
COPY . ./

# download Go modules and dependencies

# RUN go mod download

# compile application

RUN go build -o ./ascii_web
RUN apk update
RUN apk add --no-cache bash 

# tells Docker that the container listens on specified network ports at runtime

EXPOSE 8080

# command to be used to execute when the image is used to start a container

CMD [ "./ascii_web" ]
