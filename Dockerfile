FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod init app
RUN go mod download
WORKDIR /app/bin/todoapp
RUN go build -o main .
CMD ["./todoapp"]
