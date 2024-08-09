FROM golang:1.22 as builder
LABEL authors="saidkomilmakhamadkhojayev"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Install necessary packages
RUN apk add --no-cache ca-certificates postgresql-client curl
RUN apk add --no-cache --virtual .build-deps gcc musl-dev go

# Set up Go environment (only if compiling goose directly in Alpine)
ENV GOPATH /go
ENV PATH /go/bin:$PATH

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Clean up unnecessary packages
RUN apk del .build-deps

WORKDIR /root/
#go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/entrypoint.sh .

RUN chmod +x /root/entrypoint.sh

# Expose port 8000 to the outside world
EXPOSE 8000

ENTRYPOINT ["/root/entrypoint.sh"]

# Command to run the executable
CMD ["./main"]

#ENTRYPOINT ["top", "-b"]