FROM golang:1.16 as builder
RUN apt-get update 

LABEL maintainer "martiiv@stud.ntnu.com"

COPY ./main.go /service
COPY ./database /service/database
COPY ./test /service/test
COPY ./endpoints /service/endpoints
COPY ./structs /service/structs
COPY ./utils /service/utils
COPY ./webhooks /service/webhooks
COPY ./go.mod /service
COPY ./go.sum /service

EXPOSE 8080

WORKDIR /service

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o cloudproject

FROM scratch 

LABEL maintainer "martiiv@stud.ntnu.com"

WORKDIR /

CMD ["/service"]