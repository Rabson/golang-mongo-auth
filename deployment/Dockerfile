FROM golang:1.23.0-alpine
WORKDIR /app
COPY . .
RUN go build -o golang_mongo_auth main.go
CMD ["./golang_mongo_auth"]
