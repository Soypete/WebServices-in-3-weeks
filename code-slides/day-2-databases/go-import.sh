#!bin/bash
# This script will show you how to 
# get the go package from github

go mod init github.com/{username}/repo
go get "github.com/mattn/go-sqlite3"
go mod tidy

