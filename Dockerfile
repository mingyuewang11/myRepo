FROM golang:alphine AS builder
WORKDIR /app
ADD . /app
RUN go build -o main main.go

FROM alphine
WORKDIR test
COPY --from=builder /app/main .
CMD ["./main"]