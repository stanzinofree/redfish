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
**Short_Description**: Temporarily store modified tracked files
**Long_Description**: Saves your local modifications away and reverts the working directory to match the HEAD commit. Useful when you need to quickly switch context without committing unfinished work.

```sh
git stash
git stash pop
git stash list
git stash apply stash@{0}
```

## git commit
**Tags**: git, version-control, save, snapshot
**Keywords**: commit save snapshot message record changes
**Short_Description**: Record changes to the repository
**Long_Description**: Stores the current contents of the index in a new commit along with a log message describing the changes. Creates a permanent snapshot of your changes in the repository history.

```sh
git commit -m "Add new feature"
git commit --amend
git commit --amend --no-edit
```

## merge and delete branch
**Tags**: git, github, branch, merge, delete, pull-request, pr
**Keywords**: merge branch delete remove pull-request pr gh cleanup
**Short_Description**: Merge pull request and delete branch automatically
**Long_Description**: Combines changes from a pull request into the main branch and removes the feature branch both locally and remotely. Useful workflow automation for cleaning up after merging features.

```sh
gh pr merge 38 --merge --delete-branch
git branch -d feature-branch
git push origin --delete feature-branch
```

## create branch
**Tags**: git, branch, create, new
**Keywords**: branch create new checkout switch
**Short_Description**: Create and switch to a new branch
**Long_Description**: Creates a new branch from the current HEAD and immediately switches to it. Essential for starting new features or bug fixes while keeping work isolated from the main codebase.

```sh
git checkout -b feature/new-feature
git switch -c feature/new-feature
```

## git rebase
**Tags**: git, rebase, history, clean
**Keywords**: rebase interactive squash history rewrite
**Short_Description**: Rebase commits for cleaner history
**Long_Description**: Reapplies commits on top of another base tip, often used to maintain a linear project history. Interactive rebase allows you to edit, combine, or reorder commits before integration.

```sh
git rebase main
git rebase -i HEAD~3
git rebase --continue
```

## git reset
**Tags**: git, reset, undo, rollback
**Keywords**: reset undo rollback uncommit restore
**Short_Description**: Reset commits or changes
**Long_Description**: Resets the current HEAD to a specified state. Use --soft to keep changes staged, --mixed to unstage changes, or --hard to discard all changes completely. Powerful but use with caution.

```sh
git reset --soft HEAD~1
git reset --hard HEAD~1
git reset --mixed HEAD~1
```
