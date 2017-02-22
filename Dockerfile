FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y golang

COPY . /keep-alive
WORKDIR /keep-alive

RUN go build

CMD ["./keep-alive"]
