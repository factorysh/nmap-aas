

build: vendor bin
	go build -o bin/nmap-aas .

vendor:
	dep ensure

bin:
	mkdir -p bin

test:
	go test -v github.com/factorysh/nmap-aas/nmap