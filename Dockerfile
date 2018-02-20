FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u -v
RUN go build -o main .

EXPOSE 3000
EXPOSE 8080
CMD ["/app/main"]