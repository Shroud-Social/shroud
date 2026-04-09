FROM golang:1.26-alpine AS builder

ARG SERVICE_NAME

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o binary ./cmd/$SERVICE_NAME

FROM gcr.io/distroless/static-debian13:nonroot

WORKDIR /
COPY --from=builder /app/binary .

CMD ["./binary"]
