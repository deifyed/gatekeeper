FROM golang:1.16 AS build
ENV CGO_ENABLED=0
WORKDIR /go/src

COPY go.* ./

COPY pkg ./pkg
COPY main.go .

RUN ls .
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
ENV GIN_MODE=release
EXPOSE 8080/tcp

COPY --from=build /go/src/openapi ./

ENTRYPOINT ["./openapi"]
