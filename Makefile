

build: vendor bin
	go build -o bin/nmap-aas .

vendor:
	dep ensure

bin:
	mkdir -p bin