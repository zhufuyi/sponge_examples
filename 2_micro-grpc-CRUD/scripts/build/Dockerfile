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
COPY user /app/user
RUN chmod +x /app/user

# grpc and http port
EXPOSE 8282 8283


WORKDIR /app

CMD ["./user", "-c", "configs/user.yml"]
# if you use the Configuration Center, user.yml is changed to the Configuration Center configuration.
#CMD ["./user", "-c", "configs/user.yml", "-enable-cc"]
