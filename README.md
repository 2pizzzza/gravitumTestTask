# User Manager API

User Manager API — это RESTful сервис, написанный на Go, для управления пользователями. Он предоставляет CRUD-операции (Create, Read, Update, Delete) для работы с данными пользователей, хранящимися в базе данных PostgreSQL. Проект использует современные подходы к разработке, включая миграции базы данных, генерацию SQL-запросов с помощью `sqlc`, и документацию API через Swagger.

## Структура проекта

```
.
├── cmd                 # Точка входа приложения
│   └── userManager
│       └── main.go     # Основной файл запуска
├── config              # Конфигурация приложения
│   └── config.yml
├── database            # Миграции и запросы к БД
│   ├── migration       # SQL-миграции
│   └── queries         # SQL-запросы для sqlc
├── docker-compose.yml  # Конфигурация Docker Compose
├── Dockerfile          # Dockerfile для сборки приложения
├── docs                # Документация API (Swagger)
├── go.mod              # Зависимости Go
├── go.sum
├── internal            # Основная логика приложения
│   ├── app             # Инициализация приложения
│   ├── config          # Парсинг конфигурации
│   ├── domain          # Бизнес-логика
│   │   └── user        # Модуль пользователей
│   ├── http            # HTTP-обработчики
│   │   └── handler
│   ├── service         # Сервисный слой
│   ├── storage         # Слой работы с БД (PostgreSQL)
│   └── utils           # Утилиты (например, парсинг JSON)
├── Makefile            # Скрипты для автоматизации
├── pkg                 # Общие утилиты и пакеты
│   ├── httpserver      # HTTP-сервер
│   ├── logger          # Логирование
│   └── postgres        # Подключение к PostgreSQL
├── README.md           # Документация проекта
└── sqlc.yaml           # Конфигурация sqlc
```

## Зависимости

- Go 1.21+
- PostgreSQL
- [sqlc](https://sqlc.dev/) — для генерации кода запросов к БД
- [Swagger](https://swagger.io/) — для документации API
- [golangci-lint](https://golangci-lint.run/) — для линтинга кода

## Установка и запуск

### 1. Клонирование репозитория
```bash
git clone https://github.com/2pizzzza/gravitumTestTask.git
cd gravitumTestTask
```

### 2. Настройка конфигурации
Отредактируйте `config/config.yml` с вашими параметрами (например, подключение к БД: хост, порт, имя пользователя, пароль).

### 3. Запуск через Docker
Для запуска приложения с использованием Docker и PostgreSQL в контейнерах выполните:
```bash
docker-compose up --build
```
Это поднимет сервис приложения и базу данных PostgreSQL, определенные в `docker-compose.yml`.

Остановка:
```bash
docker-compose down
```

### 4. Запуск без Docker
Если вы хотите запустить приложение локально, без использования Docker:

#### Убедитесь, что PostgreSQL установлен и запущен
Установите PostgreSQL и создайте базу данных:
```bash
psql -U postgres -c "CREATE DATABASE user_manager;"
```

#### Сгенерируйте SQL-код
```bash
make generate
```

#### Сборка и запуск
```bash
go build -o bin/userManager ./cmd/userManager
./bin/userManager
```

Или используйте одну команду:
```bash
go run ./cmd/userManager
```

## API Эндпоинты

API предоставляет следующие маршруты для работы с пользователями:

### 1. Создание пользователя (`POST /user`)
- **Описание:** Создает нового пользователя.
- **Тело запроса:**
  ```json
  {
    "username": "john_doe",
    "email": "john@example.com"
  }
  ```
- **Успешный ответ (200 OK):**
  ```json
  {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "create_at": "2025-04-06T12:00:00Z",
    "update_at": "2025-04-06T12:00:00Z"
  }
  ```
- **Ошибки:**
    - `400 Bad Request` — Неверный формат тела запроса или пользователь с таким именем/почтой уже существует.
    - `400 Bad Request` — Внутренняя ошибка.

### 2. Получение пользователя (`GET /user`)
- **Описание:** Возвращает информацию о пользователе по ID.
- **Параметры запроса:** `id` (query-параметр, обязательный).
- **Пример:** `GET /user?id=1`
- **Успешный ответ (200 OK):**
  ```json
  {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "create_at": "2025-04-06T12:00:00Z",
    "update_at": "2025-04-06T12:00:00Z"
  }
  ```
- **Ошибки:**
    - `400 Bad Request` — Неверный или отсутствующий ID.
    - `404 Not Found` — Пользователь не найден.

### 3. Обновление пользователя (`PUT /user`)
- **Описание:** Обновляет данные пользователя по ID.
- **Параметры запроса:** `id` (query-параметр, обязательный).
- **Тело запроса:**
  ```json
  {
    "username": "john_doe_updated",
    "email": "john.updated@example.com"
  }
  ```
- **Пример:** `PUT /user?id=1`
- **Успешный ответ (200 OK):**
  ```json
  {
    "id": 1,
    "username": "john_doe_updated",
    "email": "john.updated@example.com",
    "create_at": "2025-04-06T12:00:00Z",
    "update_at": "2025-04-06T12:10:00Z"
  }
  ```
- **Ошибки:**
    - `400 Bad Request` — Неверный формат тела запроса или ID.
    - `404 Not Found` — Пользователь не найден.
    - `409 Conflict` — Пользователь с таким именем или почтой уже существует.

### 4. Удаление пользователя (`DELETE /user`)
- **Описание:** Удаляет пользователя по ID.
- **Параметры запроса:** `id` (query-параметр, обязательный).
- **Пример:** `DELETE /user?id=1`
- **Успешный ответ (200 OK):**
  ```json
  {
    "message": "Successes"
  }
  ```
- **Ошибки:**
    - `400 Bad Request` — Неверный или отсутствующий ID.
    - `404 Not Found` — Пользователь не найден.

## Модели данных

### CreateDTO
Используется для создания пользователя:
```go
type CreateDTO struct {
    Username string `json:"username"`
    Email    string `json:"email"`
}
```

### UpdateDTO
Используется для обновления пользователя:
```go
type UpdateDTO struct {
    Username string `json:"username"`
    Email    string `json:"email"`
}
```

### User
Модель пользователя, возвращаемая в ответах:
```go
type User struct {
    ID        int64     `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"create_at"`
    UpdatedAt time.Time `json:"update_at"`
}
```

## Ошибки домена

- `ErrUserNotFound` — Пользователь с указанным ID не найден.
- `ErrUserAlreadyExists` — Пользователь с таким именем или email уже существует.

## Документация API

Swagger-документация доступна в папке `docs`. После запуска приложения вы можете просмотреть интерактивную версию API по адресу `http://localhost:8080/swagger/index.html` (если настроен соответствующий маршрут).

## Разработка

### Генерация SQL-кода
```bash
make generate
```
Генерирует Go-код на основе SQL-запросов в `database/queries` с использованием `sqlc`.

### Линтинг кода
```bash
make lint
```
Запускает `golangci-lint` для проверки кода на соответствие стандартам.

### Форматирование кода
```bash
make fmt
```
Форматирует весь Go-код в проекте с помощью `go fmt`.

### Применение миграций
Используйте инструмент миграций (например, `migrate`) для применения файлов из `database/migration`. Пример:
```bash
migrate -path database/migration -database "postgres://postgres:password@localhost:5432/user_manager?sslmode=disable" up
```
