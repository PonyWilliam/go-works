
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v2/cmd/protoc-gen-micro
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/works.proto
	ls proto/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
.PHONY: build
build:
	go build -o works *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t ponywilliam/go-work .
	docker tag ponywilliam/go-work ponywilliam/go-work
	docker push ponywilliam/go-work
