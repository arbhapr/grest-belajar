FROM golang:1.20-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o /app/main main.go

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 4001
CMD ["/app/main"]
