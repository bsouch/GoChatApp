FROM golang:1.24.4

WORKDIR /app
COPY go.mod go.sum ./
EXPOSE 8080

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gochatapp

CMD ["/gochatapp"]