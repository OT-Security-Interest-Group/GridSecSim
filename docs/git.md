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
