# base
FROM golang:1.18.1-alpine AS base

RUN apk upgrade --update && \
    apk --no-cache add git gcc libc-dev

# development
FROM base AS dev

ENV TZ=Asia/Tokyo \
    PS1='\h /\W\ # '

RUN cd /tmp && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/rakyll/gotest@latest && \
    go install github.com/go-swagger/go-swagger/cmd/swagger@v0.29.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

CMD ["bash"]

# production
FROM base AS builder

WORKDIR /build

COPY . .

RUN go mod download && \
    GOOS=linux go build -ldflags="-s -w" -trimpath -o main cmd/online-store-server/main.go

FROM alpine:3.14.2 AS production

ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /build/main ./

CMD ["/app/main", "--host", "0.0.0.0", "--port", "9000"]
