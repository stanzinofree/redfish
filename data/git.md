# Git Commands

## git pull
**Tags**: git, version-control, sync, remote, fetch
**Keywords**: pull fetch remote update sync download

Fetch and merge changes from remote repository

```sh
git pull origin main
```

## git push
**Tags**: git, version-control, sync, remote, upload
**Keywords**: push remote upload sync send publish

Push local commits to remote repository

```sh
git push origin main
```

## git stash
**Tags**: git, version-control, temporary, save, workspace
**Keywords**: stash save temporary work-in-progress wip store

Temporarily store modified tracked files

```sh
git stash
git stash pop
git stash list
```

## git commit
**Tags**: git, version-control, save, snapshot
**Keywords**: commit save snapshot message record changes

Record changes to the repository

```sh
git commit -m "Add new feature"
git commit --amend
```

## merge and delete branch
**Tags**: git, github, branch, merge, delete, pull-request, pr
**Keywords**: merge branch delete remove pull-request pr gh cleanup

Merge pull request and delete branch automatically

```sh
gh pr merge 38 --merge --delete-branch
```

## create branch
**Tags**: git, branch, create, new
**Keywords**: branch create new checkout switch

Create and switch to a new branch

```sh
git checkout -b feature/new-feature
git switch -c feature/new-feature
```
