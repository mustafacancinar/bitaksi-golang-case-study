FROM --platform=linux/amd64 golang:1.23-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./build/main ./cmd/main.go


FROM --platform=linux/amd64 alpine:3.14 AS final

WORKDIR /app

COPY --from=builder /app/build/main .

CMD ["./main"]

EXPOSE 8080



