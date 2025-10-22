localtest:
	go test ./...

vendor:
	go mod tidy && go mod vendor

test:
	codecrafters test

submit:
	codecrafters submit

run:
	go run ./...
