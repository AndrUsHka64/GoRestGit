
FROM golang:1.20-alpine


WORKDIR /app


RUN apk add --no-cache gcc musl-dev


COPY go.mod go.sum ./


RUN go mod download


COPY database.go handlers.go models.go main.go ./



RUN go build -o main .


RUN chmod +x main



EXPOSE 8080


CMD ["./main.go"]
