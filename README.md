# Сервіс авторизації

Сервіс що виконує авторизацію користувача через видачу JWT токена.

Утиліта для міграцій
```bash
wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb && \
    dpkg -i migrate.linux-amd64.deb
```

Запуск міграцій
```bash
migrate -database $AUTH_DB_URL -path internal/database/migrations up
```

Необхідні змінні оточення (для локального запуску, без докера):
```bash
export AUTH_LOG_PATH="/home/inside/go/src/study/auth/log/log.json" && \
export AUTH_DB_DSN="user=postgres password=postgres host=127.0.0.1 port=5432 dbname=gousers sslmode=disable" && \
export AUTH_DB_URL="postgres://postgres:postgres@127.0.0.1:5432/gousers?sslmode=disable" && \
export AUTH_WEB_PORT=8001 && \
export AUTH_SECRET="mysecretkey"
```

Необхідні змінні оточення (для запуску в докер):
```bash
export AUTH_POSTGRES_PASSWORD="postgres" && \
export AUTH_POSTGRES_USER="postgres" && \
export AUTH_POSTGRES_DATABASE="gousers" && \
export AUTH_WEB_PORT=8001 && \
export AUTH_LOG_PATH="/app/log/log.json" && \
export AUTH_DB_DSN="user=postgres password=postgres host=postgres_auth port=5433 dbname=gousers sslmode=disable" && \
export AUTH_DB_URL="postgres://postgres:postgres@postgres_auth:5433/gousers?sslmode=disable" && \
export AUTH_SECRET="mysecretkey"
```
