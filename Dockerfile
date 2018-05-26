FROM golang:1.10 as builder
WORKDIR /go/src/github.com/suviano/A-Kong-Gateway
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags=\"-s\" -o app


FROM iron/go
WORKDIR /root
COPY --from=builder /go/src/github.com/suviano/A-Kong-Gateway/app .
ENTRYPOINT ["./app"]
