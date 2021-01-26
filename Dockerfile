# Builder
FROM golang:1.13.8-alpine3.11 as builder
WORKDIR /go/src/workspace/
COPY . .
RUN GOOS=linux CGO_ENABLED=0 ./build.sh

# Distribution
FROM alpine:3.11.3
WORKDIR /the-app
COPY --from=builder /go/src/workspace/build.gz .
RUN tar -xzf build.gz
CMD ["./tcp-dialer", "&"]
