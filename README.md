# Gotify matrix bot

This small program allows you to get any gotify push nodifications in a matrix chat.

## Installation guide

There are two ways to install this programm. (Using docker is recomended)

### Manual Installation

1. Pull the reposotory
2. create a `config.yaml`
3. Run `go mod download`
4. Build the program with `go build -o ./gotify-matrix-bot`
5. Run `./gotify-matrix-bot`

If your configuration is correct you should start reseving nodifications in matrix now.

### Docker Installation (recomendet)

1. Install the docker image using `docker pull ondolin/gotify-matrix-bot:latest`
2. Create a folder and a config.yaml inside it
3. Run `docker run --rm -d --name gotify-matrix-bot -v $(pwd):/data:z gotify-matrix-bot`
