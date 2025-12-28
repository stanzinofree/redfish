# Docker Commands

## docker ps
**Tags**: docker, container, list, running, status
**Keywords**: ps list containers running active status show
**Short_Description**: List running containers
**Long_Description**: Display information about running Docker containers. Shows container ID, image, command, creation time, status, ports, and names. Use -a flag to include stopped containers. Useful for monitoring container status and getting container IDs for other commands.

```sh
docker ps
docker ps -a
docker ps --format "table {{.Names}}\t{{.Status}}"
```

## docker build
**Tags**: docker, image, build, create, dockerfile
**Keywords**: build image create dockerfile compile
**Short_Description**: Build Docker image from Dockerfile
**Long_Description**: Creates a Docker image from a Dockerfile and a build context (usually current directory). The -t flag tags the image with a name and optional version. Use --no-cache to rebuild from scratch without using cached layers. Multi-stage builds can target specific stages with --target flag.

```sh
docker build -t myapp:latest .
docker build --no-cache -t myapp:latest .
docker build --target production -t myapp:prod .
```

## docker run
**Tags**: docker, container, run, start, execute
**Keywords**: run start execute container create launch
**Short_Description**: Run a command in a new container
**Long_Description**: Creates and starts a new container from an image. The -d flag runs in detached mode (background), -p maps ports (host:container), -it provides interactive terminal, -v mounts volumes, and --rm removes container after exit. Essential command for running containerized applications.

```sh
docker run -d -p 8080:80 nginx
docker run -it ubuntu bash
docker run --rm -v $(pwd):/app myapp
```

## docker compose up
**Tags**: docker, compose, orchestration, multi-container, start
**Keywords**: compose up start orchestration multi-container services
**Short_Description**: Start multi-container application
**Long_Description**: Reads docker-compose.yml and starts all defined services. The -d flag runs in detached mode, --build forces image rebuild before starting, and --scale allows running multiple instances of a service. Creates networks and volumes as defined in the compose file.

```sh
docker compose up -d
docker compose up --build
docker compose up --scale web=3
```

## docker compose down
**Tags**: docker, compose, stop, cleanup, remove
**Keywords**: compose down stop remove cleanup teardown
**Short_Description**: Stop and remove containers, networks
**Long_Description**: Stops all services and removes containers, networks created by docker compose up. The -v flag also removes named volumes, and --rmi all removes all images used by services. Leaves volumes intact by default to preserve data.

```sh
docker compose down
docker compose down -v
docker compose down --rmi all
```

## docker logs
**Tags**: docker, logs, debugging, output
**Keywords**: logs debug output tail follow
**Short_Description**: View container logs
**Long_Description**: Fetches and displays logs from a running or stopped container. The -f flag follows log output in real-time (like tail -f), --tail N shows only the last N lines, and --since can filter logs by timestamp. Essential for debugging containerized applications.

```sh
docker logs container-name
docker logs -f container-name
docker logs --tail 100 container-name
```

## docker exec
**Tags**: docker, exec, shell, command, debug
**Keywords**: exec shell bash command execute debug
**Short_Description**: Execute command in running container
**Long_Description**: Runs a new command in an already running container. The -it flags provide an interactive terminal, commonly used to open a shell (bash, sh) for debugging. Can also run single commands without interactive mode. Container must be running.

```sh
docker exec -it container-name bash
docker exec container-name ls /app
```

## docker prune
**Tags**: docker, cleanup, prune, disk-space
**Keywords**: prune cleanup remove disk space garbage
**Short_Description**: Remove unused Docker objects
**Long_Description**: Cleans up unused Docker resources to free disk space. System prune removes stopped containers, unused networks, dangling images, and build cache. The -a flag also removes unused images (not just dangling). Volume and image prune target specific resource types.

```sh
docker system prune
docker system prune -a
docker volume prune
docker image prune -a
```
