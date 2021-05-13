FROM golang:1.16 as builder
RUN apt-get update 

LABEL maintainer "martiiv@stud.ntnu.com"

ADD ./main.go /service
ADD ./database /service/database
ADD ./test /service/test
ADD ./endpoints /service/endpoints
ADD ./structs /service/structs
ADD ./utils /service/utils
ADD ./webhooks /service/webhooks
ADD ./go.mod /service
ADD ./go.sum /service

EXPOSE 8080

WORKDIR /service

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o cloudproject

FROM scratch 

LABEL maintainer "martiiv@stud.ntnu.com"

WORKDIR /

CMD ["/service"]