build:
	@go build -v ./...

fmt: 
	@go fmt ./...

test:
	go test -v -timeout 1m ./...

setup-test:
	@docker run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb:latest-pg15