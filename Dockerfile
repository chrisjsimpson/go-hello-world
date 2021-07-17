FROM golang:1.16.6-alpine3.14 AS build

WORKDIR /go/hello-world

COPY . .

RUN go build -o /bin/hello-world

FROM golang:1.16.6-alpine3.14

WORKDIR /root

COPY --from=build /bin/hello-world /bin/hello-world

EXPOSE 80

CMD /bin/hello-world
