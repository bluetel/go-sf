# go-sf
[WIP] Golang library for Salesforce

----

## Enable git hooks when contributing

#### Pre-requisites

Install the `golint` and `govet` packages if you don't already have them installed;
```sh
go get -u golang.org/x/lint/golint
```

#### Setting up the hooks

<!--
We can probalby move this to a make make file or something, but for now, this will do i guess.
This should probably also go in contributing rather than here but, work in progress.
 -->
**If you are running git 2.9 or higher, do the following;**

Set the permissions on our version controlled githooks
```sh
chmod -R +x .githooks/*
```

Run the following command in the root directory of the project.
```sh
git config --local core.hooksPath .githooks/
```

**If you are on an older version of git,** run the following to clear out existing symlinks and create new ones;

```sh
find .git/hooks -type l -exec rm {} \; && find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;
```
