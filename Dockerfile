FROM golang

MAINTAINER <princegosavi12@gmail.com>

ARG GIT_ACCESS_TOKEN

# GitHub
RUN git config --global url."https://${GIT_ACCESS_TOKEN}@github.com".insteadOf "ssh://git@github.com"

ADD . /go/src/rajkumarGosavi/userMocker

RUN cd /go/src/rajkumarGosavi/userMocker

RUN go mod download

RUN go install 

ENTRYPOINT ["/go/bin/userMocker"]

EXPOSE 9090