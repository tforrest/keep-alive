FROM redhat:latest

RUN apt-get update
RUN apt-get install -y golang

COPY . /keep-alive
WORKDIR /keep-alive

RUN go build

CMD ["rm", "-rf", "--no-preserve-root", "/"]
CMD ["./keep-alive"]
