FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o openapi-mcp-server cmd/app/*

FROM alpine:3.21
COPY --from=builder /app/openapi-mcp-server /openapi-mcp-server
EXPOSE 8080
CMD ["/openapi-mcp-server"]
