FROM golang:1.16 as builder
RUN apt-get update 

LABEL maintainer "martiiv@stud.ntnu.com"
ADD ./main.go /service
ADD ./database /service
ADD ./test /service
ADD ./endpoints /service
ADD ./structs /service
ADD ./utils /service
ADD ./webhooks /service
ADD ./go.mod /service
ADD ./go.sum /service

EXPOSE 8080

WORKDIR /service

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o cloudproject

FROM scratch 

LABEL maintainer "martiiv@stud.ntnu.com"

WORKDIR /

CMD ["/service"]