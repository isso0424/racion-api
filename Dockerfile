FROM golang:1.16.2-alpine3.13

WORKDIR /usr/src/

ARG PORT=8000

ARG HOST=localhost

COPY ./ /usr/src/

RUN echo "fetching packages..."

CMD go mod download

RUN echo "running linter..."

CMD go vet ./...

RUN echo "server starting..."

CMD go run main.go --port $PORT --host $HOST

EXPOSE $PORT
