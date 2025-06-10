# Trucking Amortization Platform

Веб-приложение для управления амортизацией и расчета кредитных обязательств в транспортном бизнесе.

## 🚀 Быстрый старт

### Локальная разработка

1. **Клонируйте репозиторий:**
```bash
git clone <repository-url>
cd Trucking-Amortization
```

2. **Запустите backend:**
```bash
# Соберите приложение
go build -o bin/server cmd/simple_main.go

# Запустите сервер
./bin/server
```

3. **Сервер будет доступен по адресу:** `http://localhost:8080`

### Docker развертывание

1. **Соберите Docker образ:**
```bash
docker build -t trucking-amortization .
```

2. **Запустите контейнер:**
```bash
docker run -p 8080:8080 trucking-amortization
```

3. **Или используйте docker-compose:**
```bash
docker-compose up -d
```

## 📋 API Endpoints

### Аутентификация

- `POST /api/v1/auth/register` - Регистрация пользователя
- `POST /api/v1/auth/login` - Вход в систему
- `GET /api/v1/profile` - Получение профиля (требует авторизации)

### Системные

- `GET /health` - Проверка состояния сервиса

## 🧪 Тестирование API

### Регистрация пользователя:
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","name":"Test User"}'
```

### Вход в систему:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### Получение профиля:
```bash
curl -X GET http://localhost:8080/api/v1/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 🏗️ Архитектура

### Backend (Go)
- **Framework:** Fiber v2
- **Authentication:** JWT токены
- **Password Hashing:** bcrypt
- **Configuration:** Environment variables

### Структура проекта:
```
cmd/
  simple_main.go          # Точка входа приложения
internal/
  config/                 # Конфигурация
  handlers/               # HTTP обработчики
  middleware/             # Middleware (аутентификация, CORS)
  models/                 # Модели данных
  services/               # Бизнес-логика
pkg/
  utils/                  # Утилиты (JWT, хеширование)
```

## 🔧 Конфигурация

Создайте файл `.env` в корне проекта:

```env
PORT=8080
JWT_SECRET=your_jwt_secret_here_change_in_production
```

## 📝 Статус разработки

### ✅ Реализовано:
- Базовая аутентификация (регистрация/вход)
- JWT токены
- Middleware для авторизации
- Docker контейнеризация
- Health check endpoints

### 🚧 В разработке:
- MongoDB интеграция
- Управление компаниями
- Управление активами
- Расчеты амортизации
- Frontend интерфейс

### 📋 Планируется:
- Полная система управления активами
- Расчеты кредитных обязательств
- Отчеты и аналитика
- Уведомления
- Экспорт данных

## 🐛 Известные проблемы

1. **MongoDB зависимости:** Временно отключены из-за проблем с `xdg-go/bson@v1.1.0`
2. **In-memory хранилище:** Данные не сохраняются между перезапусками
3. **Frontend:** Пока не подключен к backend

## 🔄 Следующие шаги

1. Исправить проблемы с MongoDB зависимостями
2. Восстановить полную функциональность с базой данных
3. Добавить управление компаниями и активами
4. Реализовать расчеты амортизации
5. Подключить frontend

## 📞 Поддержка

Если у вас возникли проблемы с развертыванием, проверьте:

1. Версию Go (требуется 1.21+)
2. Доступность портов 8080
3. Права доступа к Docker (если используется)

Для получения помощи создайте issue в репозитории. 