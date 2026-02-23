GO_CMD = go

all: build test

test:
	cd $(service) && $(GO_CMD) test -v ./...

build:
	cd $(service) && $(GO_CMD) build -o ../bin/$(service) ./cmd/main.go

docker-build:
	docker build --build-arg SERVICE=$(service) -t $(service):latest .

clean:
	rm -rf bin/
