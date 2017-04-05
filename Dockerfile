FROM alpine
MAINTAINER  Steve Sloka <steve@stevesloka.com>

RUN apk add --update ca-certificates && \
  rm -rf /var/cache/apk/*

ADD _output/bin/microservice /usr/local/bin

CMD ["/bin/sh", "-c", "/usr/local/bin/microservice"]