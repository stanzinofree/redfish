# Git Commands

## git pull
**Tags**: git, version-control, sync, remote, fetch
**Keywords**: pull fetch remote update sync download
**Short_Description**: Fetch and merge changes from remote repository
**Long_Description**: Downloads commits, files, and refs from a remote repository and immediately integrates them into the current branch. This is essentially a combination of git fetch followed by git merge.

```sh
git pull origin main
```

## git push
**Tags**: git, version-control, sync, remote, upload
**Keywords**: push remote upload sync send publish
**Short_Description**: Push local commits to remote repository
**Long_Description**: Uploads local repository content to a remote repository. After you've accumulated several local commits and are ready to share them with other team members, you can push them to the remote repository.

```sh
git push origin main
git push --force-with-lease
```

## git stash
**Tags**: git, version-control, temporary, save, workspace
**Keywords**: stash save temporary work-in-progress wip store

Temporarily store modified tracked files

```sh
git stash
git stash pop
git stash list
git stash apply stash@{0}
```

## git commit
**Tags**: git, version-control, save, snapshot
**Keywords**: commit save snapshot message record changes

Record changes to the repository

```sh
git commit -m "Add new feature"
git commit --amend
git commit --amend --no-edit
```

## merge and delete branch
**Tags**: git, github, branch, merge, delete, pull-request, pr
**Keywords**: merge branch delete remove pull-request pr gh cleanup

Merge pull request and delete branch automatically

```sh
gh pr merge 38 --merge --delete-branch
git branch -d feature-branch
git push origin --delete feature-branch
```

## create branch
**Tags**: git, branch, create, new
**Keywords**: branch create new checkout switch

Create and switch to a new branch

```sh
git checkout -b feature/new-feature
git switch -c feature/new-feature
```

## git rebase
**Tags**: git, rebase, history, clean
**Keywords**: rebase interactive squash history rewrite

Rebase commits for cleaner history

```sh
git rebase main
git rebase -i HEAD~3
git rebase --continue
```

## git reset
**Tags**: git, reset, undo, rollback
**Keywords**: reset undo rollback uncommit restore

Reset commits or changes

```sh
git reset --soft HEAD~1
git reset --hard HEAD~1
git reset --mixed HEAD~1
```
