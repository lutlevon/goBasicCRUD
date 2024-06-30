FROM golang:1.22.4

WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .

RUN go build -o bin
EXPOSE 3000
ENTRYPOINT ["/app/bin"]