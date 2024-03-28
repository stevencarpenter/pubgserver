FROM golang:1.21
LABEL authors="steven@sjcarpenter.com"

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /pubgserver

EXPOSE 80 8080 8090 5000

# Run
CMD ["/pubgserver"]
