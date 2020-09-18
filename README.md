# GoFiber in Docker

> A simple test app in Go, using Docker as the development environment.

## How this project was created:

1) Create a basic Golang `Dockerfile`:

```docker
FROM golang:1.15

WORKDIR /go/src/app
COPY . .

CMD ["go", "run", "main.go"]
```

...and `docker-compose.yml` file:

```yaml
version: '3'

services:
    app:
        build: .
        volumes:
            - .:/go/src/app
        ports:
            - "8080:8080"
        command: go run main.go
```

2) Build

```bash
docker-compose build
```

3) Run

```bash
docker-compose run app bash
```

4) Initialise a Go module

```bash
go mod init
```

5) Exit the container

6) Update the `Dockerfile` with dependencies

```Dockerfile
FROM golang:1.15

ENV PORT 8080
EXPOSE $PORT

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o main .

CMD ["./main"]
```