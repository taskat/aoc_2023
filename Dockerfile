FROM golang:1.21.5

WORKDIR /go/src/app

ENV COUNT 1
ENV BENCHTIME 1s
ENV TESTS ./...
ENV NAMES .*

ENTRYPOINT go test $TESTS -run='^$' -bench=$NAMES -benchmem -count $COUNT -benchtime $BENCHTIME