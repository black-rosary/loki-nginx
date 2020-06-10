FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./backyard.go .
COPY ./prober.go .
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/prober

FROM alpine:3.9
RUN apk --no-cache add ca-certificates && mkdir -p /opt/prober
WORKDIR /opt/prober
COPY --from=build /go/src/app/bin /opt/prober/
ENTRYPOINT /opt/prober/prober
