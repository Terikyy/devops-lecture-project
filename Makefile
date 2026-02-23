GO_CMD = go

all: build test

test:
	cd $(service) && $(GO_CMD) test -v ./...

build:
	cd $(service) && $(GO_CMD) build -o ../main ./cmd/main.go

docker-build:
	docker build --build-arg SERVICE=$(service) -t $(service) .

clean:
	rm -f main
