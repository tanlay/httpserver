FROM golang:alpine as builder
MAINTAINER tanlay
WORKDIR /data
COPY . /data
ENV GOPROXY https://goproxy.io,direct
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" main.go

FROM scratch as runner
WORKDIR /data
COPY --from=builder /data/bin/amd64/httpserver /data
EXPOSE 5678

ENTRYPOINT ["./httpserver"]