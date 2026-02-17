FROM golang:1.24-alpine

RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

RUN go install github.com/air-verse/air@v1.61.7

COPY go.mod go.sum* ./
RUN go mod download || true

COPY . .

CMD ["air", "-c", ".air.toml"]
