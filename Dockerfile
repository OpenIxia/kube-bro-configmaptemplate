FROM golang:alpine
MAINTAINER Ixia NetServices

WORKDIR /go/src/github.com/openixia/cmt-controller

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# CMD ["cmt-controller"]

FROM alpine
COPY --from=0 /go/bin/cmt-controller /usr/bin
ENTRYPOINT [ "/usr/bin/cmt-controller" ]
