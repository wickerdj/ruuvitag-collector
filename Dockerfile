FROM golang:1

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o main .
CMD ["/app/main"]