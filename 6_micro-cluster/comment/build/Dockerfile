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
COPY comment /app/comment
RUN chmod +x /app/comment

# http port
EXPOSE 8080


WORKDIR /app

CMD ["./comment", "-c", "configs/comment.yml"]
# if you use the Configuration Center, comment.yml is changed to the Configuration Center configuration.
#CMD ["./comment", "-c", "configs/comment.yml", "-enable-cc"]
