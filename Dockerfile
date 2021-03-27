FROM golang:1.16-alpine
WORKDIR /app
COPY pkg pkg
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go build -v -o bin/app main.go
ENTRYPOINT ["bin/app"]
