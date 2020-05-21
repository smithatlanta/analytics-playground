FROM golang:1.13-alpine as build
WORKDIR /app
RUN apk add librdkafka-dev build-base
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -tags musl -a -o main .

FROM alpine:edge
WORKDIR /root/
COPY --from=build /app/main .
RUN apk add --no-cache librdkafka ca-certificates
EXPOSE 8080
CMD ["./main"] 