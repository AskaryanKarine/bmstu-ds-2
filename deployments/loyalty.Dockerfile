FROM golang:1.22.5 AS builder

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY .. ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main ./cmd/loyalty/main.go

FROM alpine:latest AS app

COPY --from=builder /go/bin/main ./go/

EXPOSE 8050
CMD ./go/main
