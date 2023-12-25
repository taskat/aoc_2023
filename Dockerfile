FROM golang:1.21.5

WORKDIR /go/src/app
COPY . .

RUN go mod download

ENV COUNT 1
ENV BENCHTIME 1s

ENTRYPOINT go test ./days/day1/... -run=^$ -bench=. -benchmem -count $COUNT -benchtime $BENCHTIME