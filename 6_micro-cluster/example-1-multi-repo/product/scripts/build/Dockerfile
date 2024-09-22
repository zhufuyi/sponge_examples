FROM alpine:latest
MAINTAINER zhufuyi "g.zhufuyi@gmail.com"

# set the time zone to Shanghai
RUN apk add tzdata  \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# add grpc_health_probe for health check of grpc services
COPY grpc_health_probe /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe

COPY configs/ /app/configs/
COPY product /app/product
RUN chmod +x /app/product

# grpc and http port
EXPOSE 8282 8283


WORKDIR /app

CMD ["./product", "-c", "configs/product.yml"]
# if you use the Configuration Center, product.yml is changed to the Configuration Center configuration.
#CMD ["./product", "-c", "configs/product.yml", "-enable-cc"]
