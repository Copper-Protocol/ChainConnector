build:
	go build -o bin/chain-connecor src/main.go

run: build
	./bin/chain-connecor
