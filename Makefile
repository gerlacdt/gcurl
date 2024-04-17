.PHONY: build errcheck staticcheck test testv run clean httpbin

build: errcheck staticcheck
	go build -o gcurl main.go

errcheck: 
	errcheck ./...

staticcheck: 
	staticcheck ./...

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
