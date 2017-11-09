FROM jaehue/golang-onbuild
MAINTAINER jang.jaehue@eland.co.kr
ARG CERT_FILE

# install go packages
RUN go get github.com/relax-space/lemon-wxpay && \
    go get github.com/relax-space/go-kit/...

# setup swagger-ui
ADD $CERT_FILE /tmp/wxcert.tar.gz

# add application
ADD . /go/src/lemon-wxpay-api
WORKDIR /go/src/lemon-wxpay-api
RUN tar xf /tmp/wxcert.tar.gz
RUN go install

EXPOSE 5000

CMD ["lemon-wxpay-api"]