FROM golang:1.21.5

WORKDIR /go/src/app
COPY . .

RUN go mod download

ENV COUNT 1
ENV BENCHTIME 1s
ENV TESTS ./...
ENV NAMES ^$

ENTRYPOINT go test $TESTS -run $NAMES -bench=. -benchmem -count $COUNT -benchtime $BENCHTIME