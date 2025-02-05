# Use a Go base image to build the application
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the project files to the container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 🔥 Build the application statically to avoid problemas en Alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# Use a smaller base image to run the application
FROM alpine:latest

WORKDIR /root/

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder image
COPY --from=builder /app/main .

# 🔥 Asegurar permisos de ejecución
RUN chmod +x ./main

# Expose the API port
EXPOSE 8080

# Command to run the API
CMD ["./main"]