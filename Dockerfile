FROM golang:1.17-alpine

WORKDIR /app
COPY go.mod ./
COPY *.go ./

RUN go mod tidy \
    && go get \
    && go build -o /apprunner-golang

EXPOSE 8080

CMD ["/apprunner-golang"]