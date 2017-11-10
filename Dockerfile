FROM jaehue/golang-onbuild
MAINTAINER jang.jaehue@eland.co.kr

# install go packages
RUN go get github.com/relax-space/lemon-wxpay && \
    go get github.com/relax-space/go-kit/...


# add application
ADD . /go/src/lemon-wxpay-api
WORKDIR /go/src/lemon-wxpay-api
RUN tar xf tmp/wxcert.tar.gz
RUN go install

EXPOSE 5000

CMD ["lemon-wxpay-api"]