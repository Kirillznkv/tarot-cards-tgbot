FROM golang:1.16

WORKDIR /app


COPY . /app
RUN cd /app; go mod download

RUN cd /app; make test; make

ENTRYPOINT ["/app/main"]
