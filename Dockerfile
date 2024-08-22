FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o bin src/main/main.go

#CMD ["./bin"]
ENTRYPOINT [ "./bin" ]