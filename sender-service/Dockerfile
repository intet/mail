FROM golang:alpine as builder
RUN apk --no-cache add git

WORKDIR /app/sender-service
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sender-service


FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/sender-service/sender-service .
CMD ["./sender-service"]
