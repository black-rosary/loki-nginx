FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./prober.go .
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/prober ./prober.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
EXPOSE 80
WORKDIR /go/bin/
COPY --from=build /go/src/app/bin /go/bin
COPY ./conf.yaml /go/bin/conf.yaml
ENTRYPOINT /go/bin/prober
