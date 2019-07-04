# go-sf
[WIP] Golang library for Salesforce

----

## Enable git hooks when contributing

Git hooks will stop you from committing if there are issues highlighted by `golint` and `go vet`. There are two methods to setting up the githooks, one is using a symlink for older git versions, the other is setting the hooks folder in the local git configuration.

#### Pre-requisites

Install the `go vet` and `golint` packages if you don't already have them installed. `govet` is a standard Go package that _should_ come with your Go installation. `golint` however has to be isntalled seperately.

```sh
go get -u golang.org/x/lint/golint
```

#### Setting up the hooks

<!--
We can probalby move this to a make make file or something, but for now, this will do i guess.
This should probably also go in contributing rather than here but, work in progress.
 -->


(_Skip for windows users_) Set the permissions on our version controlled githooks.
```sh
chmod -R +x .githooks/*
```

**If you are running git 2.9 or higher, do the following;**

Run the following command in the root directory of the project.
```sh
git config --local core.hooksPath .githooks/
```

**If you are on an older version of git,** run the following to clear out existing symlinks and create new ones;

```sh
find .git/hooks -type l -exec rm {} \; && find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;
```
