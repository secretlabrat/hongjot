FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app .

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /app/app /

USER nonroot:nonroot

ENTRYPOINT ["/app"]