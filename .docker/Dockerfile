FROM golang:1.14.4 as builder

WORKDIR /go/src/github.com/lgdevlop/go-hpa/app

COPY ./app/hpa.go .

COPY ./app/loop ./loop

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app hpa.go && strip app

FROM scratch

COPY --from=builder /go/src/github.com/lgdevlop/go-hpa/app/app /app

EXPOSE 8000

CMD ["/app"]
