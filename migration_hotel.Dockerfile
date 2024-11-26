FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/hotel/*.sql migrations/hotel/
ADD migration_hotel.sh .
ADD .env .

RUN chmod +x migration_hotel.sh

ENTRYPOINT ["bash", "migration_hotel.sh"]