build:
	go build -o ./bin/custom-blockchain

run: build
	./bin/custom-blockchain

test:
	go test -v ./...