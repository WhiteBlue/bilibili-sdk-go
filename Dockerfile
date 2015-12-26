
FROM daocloud.io/golang:onbuild

RUN go get github.com/gin-gonic/gin
RUN go get github.com/bitly/go-simplejson


EXPOSE 8080

ENV VIRTUAL_HOST api.shiroblue.cn