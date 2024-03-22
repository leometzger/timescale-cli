build:
	@go build -v ./...

fmt: 
	@go fmt ./...

test:
	go test -v -timeout 1m ./...

generate-mocks:
	@mockery --output internal/domain/aggregations/mocks --dir internal/domain/aggregations --all
	@mockery --output internal/domain/hypertables/mocks --dir internal/hypertables/hypertables --all

setup-test:
	@docker run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb:latest-pg15
