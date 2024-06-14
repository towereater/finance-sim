FROM golang:1.22

WORKDIR /app

COPY . ./
RUN GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux go build -o /server

EXPOSE 13000

CMD ["/server"]