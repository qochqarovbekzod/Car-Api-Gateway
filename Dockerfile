# Stage 1: Build stage
FROM golang:1.22.2 AS builder

WORKDIR /app

# Copy and download dependencies
# COPY go.mod .
# COPY go.sum .

# Copy the rest of the application
COPY . .
RUN go mod download

# Optionally copy the .env file if it's needed
COPY .env .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp .

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .
COPY --from=builder /app/logs/app.log ./logs/
COPY --from=builder /app/config/model.conf ./config/
COPY --from=builder /app/config/policy.csv ./config/

# Optionally copy the .env file if it's needed
COPY --from=builder /app/.env .

# Expose port 8082
EXPOSE 8082

# Command to run the executable
CMD ["./myapp"]
