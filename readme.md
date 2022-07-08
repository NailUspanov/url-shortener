# OZON url shortener REST API
## Test project

Создание и запуск образов, контейнеров
```sh
docker-compose up --build -d
```

> Note: Параметры запуска лежат в .env файле в корне проекта. По умолчанию проект запускается с использованием Postgres. Чтобы запустить с использованием памяти приложения, выставите параметр RDBMS=false

Ссылки хранятся 24 часа, после запускается коллектор.

Порты:
- PostgreSQL 5436:5432
- API 8000:8000

Endpoints:
- GET /api/:short_url
- POST /api
```json
{
    "long_url": "https://website.com"
}