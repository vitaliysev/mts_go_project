FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/auth/*.sql migrations/auth/
ADD migration_auth.sh .
ADD .env .

RUN chmod +x migration_auth.sh

ENTRYPOINT ["bash", "migration_auth.sh"]