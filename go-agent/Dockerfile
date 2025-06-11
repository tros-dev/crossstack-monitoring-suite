FROM golang:1.20
WORKDIR /app
COPY . .
RUN go build -o agent main.go
CMD ["./agent"]
