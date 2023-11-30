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
COPY order_gw /app/order_gw
RUN chmod +x /app/order_gw

# http port
EXPOSE 8080


WORKDIR /app

CMD ["./order_gw", "-c", "configs/order_gw.yml"]
# if you use the Configuration Center, order_gw.yml is changed to the Configuration Center configuration.
#CMD ["./order_gw", "-c", "configs/order_gw.yml", "-enable-cc"]
