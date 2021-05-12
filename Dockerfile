FROM golang:1.16 as builder
RUN apt-get update 

LABEL maintainer "martiiv@stud.ntnu.com"
ADD ./main.go /
ADD ./endpoints /endpoints
ADD ./structs /structs
ADD ./utils /utils
ADD ./webhooks /webhooks
ADD ./go.mod /go.mod
ADD ./go.sum /go.sum

WORKDIR /

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main

FROM scratch 

LABEL maintainer "martiiv@stud.ntnu.com"

WORKDIR /

COPY --from=builder / /

ENTRYPOINT ["/main"]