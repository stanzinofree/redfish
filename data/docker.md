# Docker Commands

## docker ps
**Tags**: docker, container, list, running, status
**Keywords**: ps list containers running active status show

List running containers

```sh
docker ps
docker ps -a  # Show all containers
```

## docker build
**Tags**: docker, image, build, create, dockerfile
**Keywords**: build image create dockerfile compile

Build Docker image from Dockerfile

```sh
docker build -t myapp:latest .
docker build --no-cache -t myapp:latest .
```

## docker run
**Tags**: docker, container, run, start, execute
**Keywords**: run start execute container create launch

Run a command in a new container

```sh
docker run -d -p 8080:80 nginx
docker run -it ubuntu bash
```

## docker compose up
**Tags**: docker, compose, orchestration, multi-container, start
**Keywords**: compose up start orchestration multi-container services

Start multi-container application

```sh
docker compose up -d
docker compose up --build
```
