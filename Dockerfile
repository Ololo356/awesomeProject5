FROM golang:1.19

WORKDIR /app

COPY . .

RUN go build MyHTTP.go

CMD [".\MyHTTP.exe"]
