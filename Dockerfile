FROM golang:1.15

ENV PORT 8080
EXPOSE $PORT

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o main .

CMD ["./main"]