hello:
	echo "golang-mongo-auth"

build:
	go build -o bin/main main.go

docker-build:
	docker build -t golang_mongo_auth .

docker-run:
	docker run -d -p 8080:8080 --env-file .env golang_mongo_auth

run:
	go run main.go

start:
	air

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

all: hello build