FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/api ./cmd/api
COPY ./internal ./internal

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /automation-challenge cmd/api/main.go


EXPOSE 8080

# Run
CMD ["/automation-challenge"]