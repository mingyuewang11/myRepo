FROM golang:alpine AS builder
WORKDIR /app
ADD . /app
RUN go build -o main main.go

FROM alpine
WORKDIR test
COPY --from=builder /app/main .
CMD ["./main"]