FROM golang:1.10 as builder
WORKDIR /go/src/github.com/suviano/A-Kong-Gateway/
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo -o app


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/suviano/A-Kong-Gateway/app .
ENTRYPOINT ["./app"]
