FROM golang:1.19-alpine3.16 as builder

WORKDIR $GOPATH/src/github.com/asty-org/asty

COPY . .

RUN go clean -modcache && \
    apk add --no-cache alpine-sdk git make

RUN make
RUN cp bin/asty /

FROM alpine:3.16

COPY --from=builder /asty /asty

ENTRYPOINT ["/asty"]
