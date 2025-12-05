FROM golang:1.25 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# -o main specifies the output binary name.
RUN CGO_ENABLED=0 GOOS=linux go build -o Go-WebDAV .

# Uses Alpine linux instead of Ubuntu or another distro since its fooprint is a lot smaller.
FROM alpine:latest  

WORKDIR /root/

COPY --from=build /app/Go-WebDAV .

EXPOSE 8080

# Runs the binary to start the server.
CMD ["./Go-WebDAV"]

