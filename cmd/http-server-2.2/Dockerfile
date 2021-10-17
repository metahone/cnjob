#FROM scratch
FROM busybox:1.28.4

WORKDIR /

EXPOSE 80

COPY httpserver /

ENV VERSION=""

CMD ["./httpserver"]

