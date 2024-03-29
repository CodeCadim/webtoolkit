ifndef VERBOSE
.SILENT:
endif
.PHONY: default restart start stop clean test build-client build-server build

default: start

restart: stop start

test:
	go test ./...

install:
	go mod tidy
	cd ui/cli && npm -s i

start:
	tmux send-keys 'goconvey' C-m \; split-window -h -p 66 \; send-keys 'watcher -run ui/ops/*.go' C-m \; split-window -h \; send-keys 'cd ui/cli && npm run dev' C-m \;

stop:
	echo "Stopping dev env"
	docker-compose stop 

clean:
	echo "cleaning application"
	rm -f server
	docker-compose down

build-client:
	echo "build client"
	cd ui/cli && npm -s run build

build-server:
	echo "build server"
	go build -o ./build/ ui/srv/*.go

build: build-client build-server

