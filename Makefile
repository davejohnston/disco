all: controller

clean:
	rm -rf bin internal/api

controller:
	dep ensure
	go generate ./...
	go build -o bin/controller cmd/controller/controller.go

check:
	gofmt -s -w .
	golint ./...
	govet ./...

doc:
	open http://localhost:6060/pkg/github.com/davejohnston/disco/?m=all
	godoc -http=:6060



