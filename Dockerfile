# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu

FROM golang:1.22-alpine

LABEL author1.name="Nicolas GENTY" author1.email="nic@genty.dev" \
    author2.name="Nattan COURTEL" author2.email="" \
    author3.name="Adam GONCALVES" author3.email="goncalvesadam@icloud.com" \
    description="Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of our last project, ascii-art."

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
