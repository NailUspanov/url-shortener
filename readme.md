# OZON url shortener REST API
## Test project

Создание и запуск образов, контейнеров
```sh
docker-compose up --build -d
```

> Note: Параметры запуска лежат в .env файле в корне проекта. По умолчанию проект запускается с использованием Postgres. Чтобы запустить с использованием памяти приложения, выставите параметр RDBMS=false

Ссылки хранятся 24 часа, после запускается коллектор.

### Порты:
- PostgreSQL 5436:5432
- API 8000:8000

### Endpoints:
- GET /api/:short_url
- POST /api
```json
{
    "long_url": "https://website.com"
}
```


### SCHEMA
```postgresql
CREATE TABLE public.urls
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 ),
    short_url text NOT NULL,
    long_url text NOT NULL,
    expiration_date timestamp without time zone NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (short_url),
    UNIQUE (long_url)
);
```