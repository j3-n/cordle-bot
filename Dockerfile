FROM golang:1.20-alpine

WORKDIR /app

ADD . ./
RUN go mod download

RUN make build

EXPOSE 8080

CMD ["/app/build/program/app"]