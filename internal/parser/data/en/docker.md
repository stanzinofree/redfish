# Docker Commands

## docker ps
**Tags**: docker, container, list, running, status
**Keywords**: ps list containers running active status show

List running containers

```sh
docker ps
docker ps -a
docker ps --format "table {{.Names}}\t{{.Status}}"
```

## docker build
**Tags**: docker, image, build, create, dockerfile
**Keywords**: build image create dockerfile compile

Build Docker image from Dockerfile

```sh
docker build -t myapp:latest .
docker build --no-cache -t myapp:latest .
docker build --target production -t myapp:prod .
```

## docker run
**Tags**: docker, container, run, start, execute
**Keywords**: run start execute container create launch

Run a command in a new container

```sh
docker run -d -p 8080:80 nginx
docker run -it ubuntu bash
docker run --rm -v $(pwd):/app myapp
```

## docker compose up
**Tags**: docker, compose, orchestration, multi-container, start
**Keywords**: compose up start orchestration multi-container services

Start multi-container application

```sh
docker compose up -d
docker compose up --build
docker compose up --scale web=3
```

## docker compose down
**Tags**: docker, compose, stop, cleanup, remove
**Keywords**: compose down stop remove cleanup teardown

Stop and remove containers, networks

```sh
docker compose down
docker compose down -v
docker compose down --rmi all
```

## docker logs
**Tags**: docker, logs, debugging, output
**Keywords**: logs debug output tail follow

View container logs

```sh
docker logs container-name
docker logs -f container-name
docker logs --tail 100 container-name
```

## docker exec
**Tags**: docker, exec, shell, command, debug
**Keywords**: exec shell bash command execute debug

Execute command in running container

```sh
docker exec -it container-name bash
docker exec container-name ls /app
```

## docker prune
**Tags**: docker, cleanup, prune, disk-space
**Keywords**: prune cleanup remove disk space garbage

Remove unused Docker objects

```sh
docker system prune
docker system prune -a
docker volume prune
docker image prune -a
```
