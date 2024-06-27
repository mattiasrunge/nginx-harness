FROM docker.io/library/alpine:latest

RUN apk --no-cache add ca-certificates nginx

COPY .bin/nginx-harness /nginx-harness

EXPOSE 9898

VOLUME ["/storage"]

CMD ["/nginx-harness", "-dir", "/storage"]