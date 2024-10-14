build:
	@go build -o bin/featureflag

run: build
	@./bin/featureflag

test:
	@go test -v ./...