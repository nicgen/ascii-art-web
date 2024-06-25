# ascii-art-web

## Authors

- [agoncalv](https://zone01normandie.org/git/agoncalv)
- [ncourtel](https://zone01normandie.org/git/ncourtel)
- [ngenty](https://github.com/nicgen)

## Description

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, [ascii-art](https://github.com/01-edu/public/tree/master/subjects/ascii-art).

## Requirements

- [Go](https://golang.org/dl/) (min version 1.22.0)
- [Docker](https://docs.docker.com/engine/install/ubuntu/#installation-methods) if you want to buid and run the container version.

## Install

1. Clone the repo:

```bash
git clone https://zone01normandie.org/git/ncourtel/ascii-art-web
cd ascii-art-web
```

## Usage

### Manualy

1. Run the app:

```bash
go run main.go*
```

2. Go to `localhost:8080` and test it!

3. Have fun

### With Docker

1. Build the docker image

```sh
bash docker_build.sh 
```

2. Go to the specified adress `localhost:<port>` and test it!

3. Have fun

#### Uninstall

When finished, you can stop all containers and delete the image with the command:

```sh
bash docker_delete.sh 
```

## Project structure:

- **main.go**: main file that creates the server aand link the pages
- **lib**: contains the functions used by the app
- **static**: contains the static files like the css, images, icons and the ASCII themes used

## Attribution

- The FontStruction “[MS Sans Serif Bold](https://fontstruct.com/fontstructions/show/1384862)" by “lou” is licensed under a Creative Commons Attribution Share Alike license (http://creativecommons.org/licenses/by-sa/3.0/).

## Licence

MIT License

<!-- TODO:
- problem with static/export
- if empty, show msg
- error 404 text content is wrong
- print only useful logs
- handle errors
-->
