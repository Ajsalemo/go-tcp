FROM golang:1.23.2-alpine3.20

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/app

EXPOSE 3080

# Run
CMD ["/usr/local/bin/app"]