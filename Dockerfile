FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD . ./

RUN make build

EXPOSE 8080

CMD ["/app/build/program/app"]