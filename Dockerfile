FROM golang:1.15

ENV GOPATH /go
ENV USER root

RUN mkdir /app

COPY . /go/src/github.com/mskutle/hello-go
WORKDIR /go/src/github.com/mskutle/hello-go

RUN go get ./...

RUN cd /go/src/github.com/mskutle/hello-go
COPY . /app/

EXPOSE 3000
WORKDIR /app

CMD ["go", "run", "main.go"]
