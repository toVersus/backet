# Backet

:package: Backet is a command line tool to backup the git repositories managed by [Gitbucket](https://github.com/gitbucket/gitbucket) without any downtime. Backet is inspired by [backup shell script officially published](https://github.com/gitbucket/gitbucket/wiki/Backup) and rewritten by [Golang](https://github.com/golang/go).

## Features

Backet provides the following features for backup:

- [x] repositories and related wiki pages
- [ ] database containing users, groups and associated issues

## Usage
Backup all repositories and related wiki pages

```
$ backet backup -s <path/to/GITBUCKET_HOME> -d <path/to/backup/dir>
```
