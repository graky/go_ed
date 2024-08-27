# File Storage API

Этот проект представляет собой веб-сервис на Go, который предоставляет API для регистрации пользователей, аутентификации и управления файлами.

## Функциональность

- Регистрация пользователей
- Аутентификация пользователей с использованием JWT
- Загрузка файлов (только изображения PNG и JPEG, до 10 МБ)
- Получение списка загруженных файлов

## Технологии

- Go 1.17+
- Gin Web Framework
- PostgreSQL
- JWT для аутентификации

## Структура проекта
```
├── cmd
│   └── api
│       └── main.go
├── pkg
│   ├── auth
│   ├── config
│   ├── database
│   ├── handlers
│   ├── middleware
│   ├── models
│   └── storage
├── internal
│   └── server
├── uploads
├── go.mod
└── go.sum
```

## Настройка и запуск
Установка зависимостей:

`go mod tidy`

Настройте базу данных PostgreSQL и создайте необходимые таблицы.

Настройте переменные окружения
export DATABASE_URL="postgres://username:password@localhost/dbname?sslmode=disable"
export SERVER_ADDRESS=":8080"
export JWT_SECRET="your_secret_key"

Запуск сервера

`go run cmd/api/main.go`

## API Endpoints

- `POST /sign-up`: Регистрация нового пользователя
- `POST /sign-in`: Аутентификация пользователя и получение JWT токена
- `POST /upload`: Загрузка файла (требуется JWT токен)
- `GET /files`: Получение списка загруженных файлов (требуется JWT токен)

### Аутентификация

`curl -X POST http://localhost:8080/sign-in -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"password123"}'`

