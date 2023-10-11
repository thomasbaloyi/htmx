FROM golang:1.21

WORKDIR /app

COPY src/go.mod src/go.sum /app

RUN go mod download

COPY /src /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]