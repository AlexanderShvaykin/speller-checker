run:
	go run cmd/ya-speller/main.go
test:
	go test ./...
build:
	go build -o $(GOPATH)/bin/ya-speller cmd/ya-speller/main.go
run-maillog:
	go run cmd/tb-maillog/main.go
compile:
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/ya-speller-mac-amd64 cmd/ya-speller/main.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/ya-speller-linux-amd64 cmd/ya-speller/main.go