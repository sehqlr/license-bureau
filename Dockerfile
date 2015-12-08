FROM golang:latest
MAINTAINER sehqlr

ADD . /bureau
RUN go get gopkg.in/yaml.v2

WORKDIR /bureau
CMD go run prototype.go license.yml
