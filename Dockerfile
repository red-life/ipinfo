FROM golang:1.20-alpine AS builder
WORKDIR /src
COPY ./cmd/ ./cmd
COPY ./internal/ ./internal/
COPY ./main.go ./main.go
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o ipinfo main.go

FROM alpine:3.14
WORKDIR /ipinfo
COPY ./.env ./.env
COPY --from=builder /src/ipinfo ./ipinfo
CMD ["./ipinfo"]
