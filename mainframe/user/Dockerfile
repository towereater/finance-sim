FROM golang:1.24.2

# Exposed server port required at build time
ARG SERVER_PORT

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE ${SERVER_PORT}

ENTRYPOINT ["/main"]
CMD ["./config.yml"]