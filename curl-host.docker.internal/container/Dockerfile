FROM alpine

RUN apk add curl

COPY entrypoint.sh /
RUN chmod 0755 /entrypoint.sh

ENTRYPOINT /entrypoint.sh