

build: vendor bin
	go build -o bin/nmap-aas .

vendor:
	dep ensure

bin:
	mkdir -p bin

test:
	go test -v github.com/factorysh/nmap-aas/nmap

docker-build:
	docker run --rm -ti \
		-v `pwd`:/go/src/github.com/factorysh/nmap-aas \
		-v `pwd`/.cache:/.cache \
		-w /go/src/github.com/factorysh/nmap-aas \
		-u `id -u` \
		bearstech/golang-dep:stretch \
		make build

docker-image:
	docker build -t nmap-aas .