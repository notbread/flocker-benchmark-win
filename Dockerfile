FROM golang:1.16.4 as builder

RUN mkdir /build
COPY main.go go.mod /build/
WORKDIR /build
RUN CGO_ENABLED=0 go build -o app

FROM alpine:latest
RUN apk --no-cache add ca-certificates && update-ca-certificates
RUN mkdir /app
COPY --from=builder /build/app /app/app
ENTRYPOINT ["/app/app"]
