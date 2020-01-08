FROM golang:1.13-alpine

ADD . /workspace
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go build -o /bin/server /workspace/*.go

FROM modokipaas/no-app
