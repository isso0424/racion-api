FROM golang:1.16.2-alpine3.13 as build

WORKDIR /usr/src/

COPY ./ /usr/src/

RUN go build -tags netgo -o racion-api .

FROM alpine:3.13

ENV PORT=8000 \
    HOST=localhost

COPY --from=build /usr/src/racion-api /usr/local/bin

CMD /usr/local/bin/racion-api --port $PORT --host $HOST
