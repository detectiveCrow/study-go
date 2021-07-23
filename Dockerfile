FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go mod init \
    && go mod tidy

CMD ["go", "run", "main.go"]
