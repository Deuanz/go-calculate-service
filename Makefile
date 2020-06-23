run:
	go run main.go

build:
	go build

run-test:
	export GIN_MODE=test && go test -failfast ./test/... -coverprofile .coverage.txt