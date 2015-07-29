FROM golang

ADD . /go/src/github.com/whobrokethebuild/goingup
#ADD goingup-example/templates/ /go/bin/templates
#ADD goingup-example/static/ /go/bin/static

RUN go get github.com/gorilla/mux
RUN go install github.com/whobrokethebuild/goingup/goingup-example

WORKDIR /go/bin
ENTRYPOINT /go/bin/goingup-example

EXPOSE 8080