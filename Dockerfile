FROM golang:1.16 as builder
RUN apt-get update 

LABEL maintainer "martiiv@stud.ntnu.com"

ADD ./endpoints /cloudproject/endpoints
ADD ./structs /cloudproject/structs
ADD ./utils /cloudproject/utils
ADD ./webhooks /cloudproject/webhooks
ADD ./go.mod /cloudproject/go.mod
ADD ./go.sum /cloudproject/go.sum

WORKDIR /cloudproject

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o cloudproject

FROM scratch 

LABEL maintainer "martiiv@stud.ntnu.com"

WORKDIR /

COPY --from=builder /cloudproject /cloudproject

CMD ["/cloudproject"]