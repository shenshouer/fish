.PHONY: default meet server client all clean

default: all

meet:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o meet/meet meet/meet.go
	docker build -t shenshouer/fish-demo-meet -f ./meet/Dockerfile .
	docker push shenshouer/fish-demo-meet

server:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o server/server server/server.go
	docker build -t shenshouer/fish-demo-server -f ./server/Dockerfile .
	docker push shenshouer/fish-demo-server

client:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o client/client client/client.go
	docker build -t shenshouer/fish-demo-client -f ./client/Dockerfile .
	docker push shenshouer/fish-demo-client

clean: 
	rm meet/meet
	rm server/server
	rm client/client

all: meet server client clean
