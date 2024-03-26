download:
	go mod download

build:
	rm -rf bin && go build -o ./bin/ ./src/*.go

build-orm:
	rm -rf bin && go build -o ./bin-orm/ ./src/application/orm/*.go

test:
	go test -v ./tests/*.go

test-coverage:
	rm -rf coverage
	mkdir coverage
	go test -v ./tests/*.go  -coverpkg=./src/... -coverprofile ./coverage/coverage.txt
	go tool cover -html ./coverage/coverage.txt -o ./coverage/coverage.html

lint:
	golangci-lint run ./src/...

api-run:
	go run ./src/main.go

docker-build:
	docker build -t go-api:latest . -f Dockerfile.local

docker-up-api:
	docker-compose up -d go-api

docker-up-store:
	docker-compose up -d go-store

docker-up-orm:
	docker-compose up -d go-orm

docker-down:
	docker-compose down --rmi all