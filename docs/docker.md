# Docker Guide

This document provides a practical quick-reference for working with Docker and Docker Compose.  
For an introduction to container concepts, see the [Docker Docs: Get Started](https://docs.docker.com/get-started/).  

---

## Essential Docker Commands
These are the most common commands you'll need for building, running, and managing your local containers.
### 1. Build and Run the Application
#### docker compose up
**Description**: This command will allow you to build and start the containers specified within a docker compose file.

**Common Use**:
- `docker compose up -d` -- If you are in the same directory as the `docker-compose` file this will compose up them without taking over your terminal with logs (detached mode).
- `docker compose -f <path-to-file> up` -- This one allows you to specify a docker compose file to one not in your directory.
- `docker compose up <service-name>` -- This will only turn on any service you say after the up in the compose file it finds.
- `docker compose -f <path-to-file> up <service-name> <service-name> -d` -- This combines them together to build and start two services at the specified docker compose file path in detached mode, so it doesn't take over your terminal.
#### docker compose down
**Description**: Allows you to stop and delete all containers of a specified file.

**Common Use**:
- `docker compose -f <path-to-file> down` -- This is standard use with or without `-f`
- `docker compose down -t <time>` -- This allows you to specify how long you want docker to wait before forcing shutdown on a container
#### docker compose stop
**Description**: This just stops all containers in specified file

**Common Use**:
- `docker compose -f <path-to-file> stop` -- This doesn't delete the containers like down does. Again with or without the `-f` flag if you are in the directory
- `docker compose stop <service-name>` -- This will just stop the specific service instead of the whole compose.
#### docker compose build
**Description**: This just builds the images for all the services in a docker compose.

**Common Use**:
- `docker compose -f <path-to-file> build` -- Will build images but not start or build containers.
- `docker compose build --no-cache` -- This will force the docker engine to not use cached layers in the build process of the image. Good for if you make a change docker doesn't auto detect

### 2. Monitoring and Troubleshooting
#### docker ps
**Description**: This will list all the containers on your system that are running be default

**Common Use**:
- `docker ps` -- List all running containers
- `docker ps -a` -- List all running and stopped containers on your system
#### docker exec
**Description**: This allows you to almost `ssh` into the containers when you want but also run commands inside of them.

**Common Use**:
- `docker exec -it <container-name/id> /bin/bash` -- This will open a shell inside the container like `ssh` although some containers don't have bash so you can use `/bin/sh` or other shells to remote in.

### 3. Management and Cleanup
#### docker stop
**Description**: This stops a specific container based on name or ID.

**Common Use**:
- `docker stop <container-name-id>` -- Will stop the specified container

#### docker system prune
**Description**: This will help resolve some issues with docker getting broken. This is sort of like a reboot for docker by cleaning out all the saved stuff. This will delete all images and stuff on your computer

**Common Use**:
- `docker system prune -af` -- This is the most common command here that deletes everything and doesn't ask you for every section if you are sure.
- `docker network prune -af` -- This will only clean out unused networks. Sometimes does the trick

> **NOTE**: Running containers won't be pruned by `docker prune`

## Common Workflow
```bash
# Spin up containers
docker compose -f <path-to-file> up -d

# See what happened and work
docker ps
docker network ls # show all docker networks on computer
docker exec -it <container-name/id> /bin/bash

# If you want to make a change to a service and test the change
docker compose -f <path-to-file> stop <service>
# Make edits
docker compose -f <path-to-file> build <service> --no-cache
docker compose -f <path-to-file> up <service> -d
# Repeate this section until you have done your work

# End
docker compose -f <path-to-file> down
```

> **NOTE**: We have command line tools to make this process easier or you can use docker desktop but this is the default docker workflow you might do if you are going through making a change to one container.

---

## Dockerfiles

A **Dockerfile** defines how an image is built. Common instructions:

- `FROM <image>` — Base image (e.g., `FROM python:3.12-slim`).  
- `WORKDIR <path>` — Sets working directory inside container.  
- `COPY <src> <dest>` — Copy files into the image.  
- `RUN <command>` — Execute a command at build time (e.g., install packages).  
- `EXPOSE <port>` — Document the port the app runs on (doesn’t publish it).  
- `CMD ["executable", "param1"]` — Default command run when container starts.  
- `ENTRYPOINT ["executable", "param1"]` — Like CMD, but arguments passed to `docker run` are appended.  

**Example Dockerfile:**
```dockerfile
FROM node:20-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install --production
COPY . .
EXPOSE 3000
CMD ["node", "server.js"]
```

---

## Docker Compose

A **docker-compose.yml** defines and manages multi-container applications.  

### Docker Sections

#### services
This is where you define the options for each of the containers in the compose 

##### image
use this option when you are using a prebuilt image like pulling one from docker hub.

##### build
This is for building the docker image yourself with a docker file
- **context**: The path to the build context.
- **dockerfile**: The name of the Dockerfile.
- **args**: Build arguments to pass to the Dockerfile.
- **target**: The target stage to build in a multi-stage Dockerfile.

##### command and entrypoint
These two options are for overriding the default operation inside the image.
- **command**: Overrides the default command specified in the image.
- **entrypoint**: Overrides the default entrypoint specified in the image.

##### Environment Variables
These are for setting the env options that are built into docker images
- **environment**: Sets environment variables within the service container.
- **env_file**: Specifies a file containing environment variables to load.

##### Ports and Networking
This is where the network options for the container are defined
- **ports**: Maps ports between the host and the container.
- **expose**: Exposes ports to linked services without publishing them to the host.
- **networks**: Connects the service to specified networks.
  - **aliases**: Defines network aliases for the service within a network.
  - **ipv4_address/ipv6_address**: Assigns a static IP address to the service within a network.


##### Volumes and Data Persistence
Mainly these are just volumes that act as a permanent storage outisde of the container
- **volumes**: Mounts host paths or named volumes into the container for data persistence.
- **tmpfs**: Mounts a temporary filesystem into the container.

##### Dependencies and Relationships
Here you can define if a service needs another to run
- **depends_on**: Defines dependencies between services, ensuring services start in a specific order.


##### Restart Policies
This only has the one option for defining how you want the container to work on exit/crash
- **restart**: Configures how containers should restart when they exit/crash (e.g., no, on-failure, always, unless-stopped).


##### Other Options
- **container_name**: Specifies a custom name for the container.
- **labels**: Adds metadata labels to the service.
- **privileged**: Grants extended privileges to the container.
- **working_dir**: Sets the working directory inside the container.

```yaml
services:
  web:
    build: .
    ports:
      - "8080:80"
    volumes:
      - .:/app
    networks:
      - app-net

  db:
    image: postgres:16
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: pass
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - app-net
```

#### Networking
The networking section is where the docker network options are defined. If you need to define a network do so in the networking compose file and then for every network needed in the other use the `external: true` option to inform docker that it is defined elsewhere.

```yaml
networks:
  assembly_network:
    external: true
  utilities_network:
    external: true
  packaging_network:
    external: true
  site_operations_network:
    external: true
```
#### Volumes
Volumes are what allo data to be persistant beyond a container lifetime. There aren't to many options that get used with them often

```yaml
volumes:
  db-data:
```

---

This file is meant as a **quick-reference**. For deeper dives, see:  
- [Docker Docs](https://docs.docker.com/)  
- [Docker Compose Reference](https://docs.docker.com/compose/compose-file/)  

