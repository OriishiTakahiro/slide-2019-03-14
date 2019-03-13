FROM golang:1.12.0-alpine

ADD app /src/app

ENV HOGE 'hogehoge yeaaaah!'

WORKDIR /src/app
RUN     go build -o /opt/server server.go && \
        mkdir /var/logs && \
        touch /var/logs/server.log && \
        adduser -D executer && \
        chown executer /opt/server && \
        chown executer /var/logs/server.log

USER executer

CMD ["/opt/server"]

LABEL maintainer "Oriishi Takahiro <takahiro0914@live.jp>"
