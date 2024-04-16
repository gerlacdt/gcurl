.PHONY: build test testv run clean httpbin

build:
	staticcheck ./... && errcheck ./... && go build -o gcurl main.go

test:
	go test ./...

testv:
	go test -v ./...

golden:
	go test ./... -update
	
run: 
	go run main.go

clean:
	rm -f ./gcurl

httpbin:
	docker run -d -p 8080:80 kennethreitz/httpbin
