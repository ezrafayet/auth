
dev:
	go run cmd/httpserver/main.go

# Tests, build and execution are handled in the Dockerfile.
# We can still have them for reference:

test:
	go test ./... -cover

build:
	go build -o out/app cmd/httpserver/main.go

start:
	./out/app
