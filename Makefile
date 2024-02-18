.PHONY: build run clean

build:
	staticcheck ./... && errcheck ./... && go build -o gcurl main.go
	
run: 
	go run main.go

clean:
	rm -f ./gcurl
