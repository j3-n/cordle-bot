FROM golang:1.20-alpine

WORKDIR /app

ADD . ./
RUN go mod download

RUN go build -o build/program/app cmd/cli/main.go 

EXPOSE 8080

CMD ["/app/build/program/app"]