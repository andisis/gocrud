run: build
	@./bin/gocrud

test:
	@go fmt ./...
	@go vet ./...
	@go test -v -coverprofile=coverage.out ./...

coverage:
	@go tool cover -html=coverage.out

build:
	@go mod tidy
	@go build -o bin/gocrud cmd/gocrud/main.go

engine:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/docker/gocrud cmd/gocrud/main.go

image: engine
	@docker build -t gocrud .

migrate:
	@$(GOPATH)/bin/sql-migrate up -env="development"