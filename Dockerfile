FROM daocloud.io/golang:onbuild

RUN go get github.com/gin-gonic/gin
RUN go get github.com/bitly/go-simplejson

ENV VIRTUAL_HOST api.shiroblue.cn

EXPOSE 8080
