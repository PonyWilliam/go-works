FROM nimmis/golang

MAINTAINER  "william"
RUN mkdir /app
WORKDIR /app/
ADD . /app/
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY="https://goproxy.io"
RUN go env -w GOOS="linux"
RUN go build -o works .
ENV GOPATH=/root/go
ENTRYPOINT  ["./works"]
