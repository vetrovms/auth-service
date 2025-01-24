FROM golang:latest

ARG ARG_AUTH_DB_DSN
ARG ARG_AUTH_DB_URL
ARG ARG_AUTH_WEB_PORT
ARG ARG_AUTH_LOG_PATH
ARG ARG_AUTH_SECRET

ENV AUTH_DB_DSN $ARG_AUTH_DB_DSN
ENV AUTH_DB_URL $ARG_AUTH_DB_URL
ENV AUTH_WEB_PORT $ARG_AUTH_WEB_PORT
ENV AUTH_LOG_PATH $ARG_AUTH_LOG_PATH
ENV AUTH_SECRET $ARG_AUTH_SECRET

# системні пакети
RUN apt-get update && \
    apt-get -y install --no-install-recommends --no-install-suggests \
    git

# утиліта для міграцій golang-migrate
RUN wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb && \
    dpkg -i migrate.linux-amd64.deb

# очистка
RUN apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# гіт
RUN rm -rf /app \
    && mkdir -p /app \
    && cd /app \
    && git init \
    && git remote -v \
    && git remote add origin https://github.com/vetrovms/auth-service.git \
    && git remote -v \
    && git fetch origin master \
    && git checkout master \
    && git config --global user.name "Thread" \
    && git config --global user.email thread@vipdesign.com

WORKDIR /app

# GO install
RUN go mod tidy

# білд
RUN go build -C cmd -o app

# міграції та запуск застосунку
ENTRYPOINT migrate -database "$AUTH_DB_URL" -path internal/database/migrations up && cmd/app