FROM golang:1.16 AS build
WORKDIR /go/src
ENV CGO_ENABLED=0

COPY go.* ./

RUN go get -d -v ./...

COPY specification.yaml .
COPY main.go .

COPY ./pkg ./pkg

RUN go build -a -installsuffix cgo -o gatekeeper .

FROM scratch AS runtime
ENV GIN_MODE=release
EXPOSE 4554/tcp
ENTRYPOINT ["./gatekeeper"]

COPY --from=build /go/src/gatekeeper ./
