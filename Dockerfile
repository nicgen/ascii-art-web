# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu

FROM golang:1.22-alpine

# create a working directory inside the image

WORKDIR /app

# TODO LABEL
# copy directory files i.e all files ending with .go

# copy all files/folder into /app
COPY . ./
# Lib
# COPY lib ./
# # CSS
# COPY static/css ./
# # Fonts
# COPY static/fonts/converted ./
# # HTML templates
# COPY static/html ./
# # icons
# COPY static/icon ./
# # images
# COPY static/img ./
# # ASCII themes
# COPY static/themes ./
# # mod & work
# COPY go.mod ./
# COPY go.work ./
# # main
# COPY main.go ./
# # COPY *.go ./
# COPY static ./

# download Go modules and dependencies

# RUN go mod download

# compile application

RUN go build -o ./ascii_web && apk update && apk add bash

RUN apk add --no-cache bash

# && apk update && apk add bash

# tells Docker that the container listens on specified network ports at runtime

EXPOSE 8080

# command to be used to execute when the image is used to start a container

CMD [ "./ascii_web" ]
