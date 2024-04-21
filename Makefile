.PHONY: build

build: errcheck staticcheck
	go build -o gcurl main.go


.PHONY: errcheck
errcheck: 
	errcheck ./...

.PHONY: staticcheck
staticcheck: 
	staticcheck ./...

.PHONY: test
test:
	go test ./...

.PHONY: testv
testv:
	go test -v ./...

.PHONY: golden
golden:
	go test ./... -update
	
.PHONY: run
run: 
	go run main.go

.PHONY: clean
clean:
	rm -f ./gcurl

.PHONY: httpbin
httpbin:
	docker run -d -p 8080:80 kennethreitz/httpbin
