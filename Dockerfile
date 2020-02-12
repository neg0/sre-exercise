# Build
FROM golang:1.13.7-alpine3.10 as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go test ./...

RUN GOOS=linux GOARCH=amd64 go build -o dist/alerts cmd/alerts/alerts.go

# Deploy
FROM scratch
COPY --from=builder /app/dist/alerts /app/
COPY --from=builder /app/service.json /app/
EXPOSE 8080
ENTRYPOINT ["/app/alerts", "./app/service.json"]