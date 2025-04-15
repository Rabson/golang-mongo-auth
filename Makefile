hello:
	echo "golang-mongo-auth"

build:
	go build -o bin/main cmd/app/main.go

seed:
	go run main.go seed $(module) --dry-run

clean:
	go mod tidy
	go mod vendor
	go clean
	rm -rf bin
	rm -rf vendor

dev:
	air

docker-build:
	docker build -t golang_mongo_auth deployment

docker-run:
	docker run -d -p 8080:8080 --env-file .env golang_mongo_auth

kub-run-deployment:
	kubectl apply -f deployment/kubernetes/deployment.yaml

kub-run-scv:
	kubectl apply -f deployment/kubernetes/service.yaml

kub-show:
	kubectl get deployments
	kubectl get services

lint:
	golangci-lint run

run:
	go run main.go

start:
	air

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm cmd/app/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 cmd/app/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/app/main.go

all: hello build