FROM ubuntu:14.04
MAINTAINER tobe tobeg3oogle@gmail.com

RUN apt-get update

# Install Go
RUN apt-get install golang
# gopath
go get beego
go get bee

# Install build tools
RUN apt-get install -y git


RUN git clone git@github.com/tobegit3hub/mycomputer.xyz


CMD ["./mycomputer"]

# Todo: run in container, mysql not support utf-8 by default