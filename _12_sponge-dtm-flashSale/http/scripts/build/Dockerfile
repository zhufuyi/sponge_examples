FROM alpine:latest
MAINTAINER zhufuyi "g.zhufuyi@gmail.com"

# set the time zone to Shanghai
RUN apk add tzdata  \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# add curl, used for http service checking, can be installed without it if deployed in k8s
RUN apk add curl

COPY configs/ /app/configs/
COPY flashSale /app/flashSale
RUN chmod +x /app/flashSale

# http port
EXPOSE 8080


WORKDIR /app

CMD ["./flashSale", "-c", "configs/flashSale.yml"]
# if you use the Configuration Center, flashSale.yml is changed to the Configuration Center configuration.
#CMD ["./flashSale", "-c", "configs/flashSale.yml", "-enable-cc"]
