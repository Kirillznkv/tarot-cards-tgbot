FROM golang:1.16


RUN curl -fsSL \
        https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
        sh


WORKDIR /app


COPY .. /app
RUN cd /app; go mod download

RUN cd /app; make test; make

ENTRYPOINT ["/app/main"]
