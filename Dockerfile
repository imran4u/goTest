FROM golang:1.23.0-bookworm

WORKDIR /app

COPY . .

RUN go get -u github.com/gin-gonic/gin

RUN go build -o bin src/main/main.go

#CMD ["./bin"]
ENTRYPOINT [ "./bin" ]