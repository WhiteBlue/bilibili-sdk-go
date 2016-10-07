FROM golang:onbuild

MAINTAINER whiteblue0616@gmail.com

ADD . $GOPATH/src/github.com/whiteblue/bilibili-go

WORKDIR $GOPATH/src/github.com/whiteblue/bilibili-go


RUN go get -u github.com/go-playground/log \
    && go get -u github.com/gin-gonic/gin \
    && go get -u github.com/gin-gonic/contrib/gzip

EXPOSE 8080

CMD ["./run.sh"]

