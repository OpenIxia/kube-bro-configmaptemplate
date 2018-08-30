FROM golang:alpine
MAINTAINER Ixia NetServices

WORKDIR /go/src/github.com/openixia/kube-bro-configmaptemplate

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# CMD ["kube-bro-configmaptemplate"]

FROM alpine
COPY --from=0 /go/bin/kube-bro-configmaptemplate /usr/bin
ENTRYPOINT [ "/usr/bin/kube-bro-configmaptemplate" ]
