FROM golang:1.13

ADD . /workspace
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go build -o /bin/server /workspace/*.go

FROM scratch

COPY --from=0 /bin/server /bin/

ENTRYPOINT ["/bin/server"]