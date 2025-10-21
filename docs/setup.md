This doc will guide you through the process of getting started with GridSecSim. The setup is meant to be minimal and clear to follow, so you can get started on OT security as quickly as possible.
#### Table of Contents
- [Environment Setup](#environment)
- [Working with git](#git)
- [Working with docker](#docker)
- [Using our Tools](#tooling)

<a id="environment"></a>
# Environment Setup
This will lead you through installing Docker if you don't have it. It will also have instructions for installing other tools, such as `git`, and 
## Windows
### Docker
1. Go to Docker Desktop install docs and pick the architecture you are on [Link](https://docs.docker.com/desktop/setup/install/windows-install/)

[install-docker-windows.png]

2. While that is installing, check if you have WSL installed by running wsl --version in CMD or PowerShell. You need `wsl` version 2.1.5 or higher. Skip to step 5 if you have the right version

```powershell
wsl --version
```


3. Open PowerShell or CMD in admin mode by typing into the Windows search bar for either, then selecting the option that shows up to run as administrator.
4. Now that you have admin PowerShell or CMD, run the appropriate command

```powershell
wsl --install
# or
wsl --update
```

> **NOTE**: If you are having trouble installing WSL, here is a link to Microsoft's docs. [Link](https://learn.microsoft.com/en-us/windows/wsl/install)

5. Now that you have `wsl` installed, it is time to run the `Docker Desktop Installer.exe` that by default is installed at `C:\Program Files\Docker\Docker`.
6. Make sure you select `wsl` instead of Hyper-V (unless you are using Hyper-V)
7. Follow the installation wizard
8. Once it is complete, if your user isn't an admin, then you need to add your user to the `docker-users` group

```powershell
net localgroup docker-users <user> /add
```

### git
If you don't already have git installed, go here [Link](https://git-scm.com/install/windows)

## Mac
### Docker
1. Go to Docker Desktop install docs and pick the type of chip you have [Link](https://docs.docker.com/desktop/setup/install/mac-install/)

![[install-docker-mac.png]]


2. Once downloaded, double-click on the dmg file to finish the installation.

### git
If you don't have `git` installed use homebrew to install

```sh
brew install git
```
<a id="git"></a>
# Working with git
`git` is a version control system that allows you to keep track of work. We are using GitHub as the project's remote repository. If you know nothing about `git`, then watch this [video](https://www.youtube.com/watch?v=hwP7WQkmECE) that covers the very basics.
## Git commands
This section covers the basics for working with git.
### git status
**Description**: This command lets you know the status of the git repository. It lists out all the staged and unstaged files in the repo, along with any changes you need to pull or push with the remote repository.

**Standard Uses**:
- `git status` -- This command doesn't have any common options

### git add
**Description**: Stage changes to be included in the next commit.

**Standard Uses**:
- `git add .` -- This will stage all the files within the current directory and subfolders.
- `git add <file/dir-name` -- This stages specific file or directory
- `git add -u` -- This stages all tracked files. So no new files
- `git add -A` -- This does stage all changes across all files in the repo

### git commit
**Description**: Takes all staged changes and records them in a commit.

**Standard Uses:**
- `git commit -m "<commit-message>"` -- This commits with a message directly in the command.
- `git commit -a -m "<commit-message>"` -- This will run the `git add -u` before committing, meaning it will stage all modified files but not new files before committing.
### git push
**Description**: This takes all the commits on the local and pushes them to the remote repository on GitHub.

**Standard Uses**:
- `git push` -- pushes commits to remote repository

### git pull
**Description**: This grabs all the commits from the remote repository.

**Standard Uses**:
- `git pull` -- gets commits from the remote repository
### git branch
**Description**: This command lets you work with branches by either listing, renaming, or deleting them.

**Standard Uses**:
- `git branch -a` -- List all remote and local branches
- `git branch <name>` -- Creates a new branch with the name
- `git branch -d <branch-name>` -- Deletes branch

### git switch
**Description**: This switches which branch you are working within. 

**Standard Uses**:
- `git switch <branch-name>` -- switches you to the specified branch name

## Git Workflow
There are general workflows when working with git, and there are also different strategies for working in the repository. We are going to use `feature branching` for the most part, though we aren't very strict about it beyond not making changes in the main branch.

```shell
# Get the new changes since you last worked on the project
git pull

# Start a new branch for your work
git branch <new-branch-name>
git switch <branch-name>

# Make changes
git status
git add -A
git status
git commit -m "<message>"
git status
git push

# When you are done with the branch
git switch main
git pull
git merge <branch-name> # Solve any merge conflicts

# Deleting Branch
git branch -d <branch-name>
git push origin --delete <branch-name>
```

## How do we use it
We use `git` with loose feature-branching workflows, where you create a branch for a feature you're working on and then merge it back in relatively quickly. We don't have a dev branch or anything like that.

<a id="docker"></a>
# Working with Docker
**Docker** is a platform that uses operating-system-level virtualization to deliver software in standardized units called **containers**. It essentially packages an application and all its dependencies—code, runtime, libraries, and settings—into one isolated environment, ensuring the application runs consistently regardless of the host machine.

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

<a id="tooling"></a>
# Our Tooling
To make the development process easier there are several scripts that we have created in order to make the development process faster. 

## tools.sh
`tools.sh` is built to provide several benefits to the standard Docker commands. If you didn't set up the bash config with `tools.sh` then all that needs to be done is `source tools.sh` in every shell you enter.

### What it does
`tools.sh`, when sourced, finds the Docker Compose files in the current and subdirectories to build a list of services. Then, using the command `start`, you can build and start whatever containers you want with tab autocomplete on their name.

### Example use
```sh
source tools.sh
start -u <service> <service> <service>

stop -d <service>
start -u <service>

stop -s <service> <service> <service>
```

## Controller Container
This is a container within the network Docker Compose that allows us to use custom routers. 

### What it does
It will listen on the Docker socket for any container start events. When it hears them, it will exec into the container and attempt to change the default route to what we want (our routers). 
