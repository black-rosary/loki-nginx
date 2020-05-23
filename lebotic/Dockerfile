FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY ./src/go.mod .
COPY ./src/go.sum .

RUN go mod download
COPY ./src .

RUN go build -o ./bin/lebotic ./main.go
#RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/lebotic ./main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 80
ENTRYPOINT /go/bin/lebotic
