FROM golang:1.16 as builder
RUN apt-get update 

LABEL maintainer "martiiv@stud.ntnu.com"

ADD ./studentdb /cloudproject/endpoints
ADD ./mongodb /cloudproject/structs
ADD ./utils /cloudproject/utils
ADD ./webhooks /cloudproject/webhooks
ADD ./main /cloudproject/main.go
ADD ./go.mod /cloudproject/go.mod
ADD ./go.sum /cloudproject/go.sum

WORKDIR /cloudproject

RUN go get https://git.gvk.idi.ntnu.no/MartinIversen/cloudproject.git

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o studentdb

FROM scratch 

LABEL maintainer "martiiv@stud.ntnu.com"

WORKDIR /

COPY --from=builder /cloudproject /cloudproject

CMD ["/cloudproject"]