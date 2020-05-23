FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY ./prober.go .
RUN go get ./...
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/prober ./prober.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 80
ENTRYPOINT /go/bin/prober
