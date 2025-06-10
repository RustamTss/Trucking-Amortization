# Trucking Amortization Platform

Полноценное веб-приложение для управления амортизацией и депрециацией в траковом бизнесе.

## Технический стек

### Backend
- **Golang** с Fiber framework
- **MongoDB** для хранения данных
- **JWT** для аутентификации
- **Docker** для контейнеризации

### Frontend
- **React** с TypeScript
- **Vite** для сборки
- **Tailwind CSS** для стилей
- **Zustand** для управления состоянием
- **React Router** для навигации
- **Axios** для HTTP запросов

## Функциональность

### Основные возможности
- 🔐 Аутентификация и авторизация пользователей
- 🏢 Управление компаниями
- 🚛 Управление активами (траки и трейлеры)
- 📊 Расчет графиков амортизации кредитов
- 📈 Расчет графиков депрециации активов
- 💰 Анализ общей задолженности по компании
- 📱 Адаптивный дизайн

### Расчеты
- **Амортизация кредита**: Расчет ежемесячных платежей по формуле аннуитета
- **Депрециация активов**: Прямолинейный метод амортизации
- **Бизнес-задолженность**: Сводка по всем кредитам компании

## Структура проекта

```
trucking-amortization/
├── backend/
│   ├── cmd/main.go                 # Точка входа
│   ├── internal/
│   │   ├── config/                 # Конфигурация
│   │   ├── handlers/               # HTTP обработчики
│   │   ├── models/                 # Модели данных
│   │   ├── services/               # Бизнес-логика
│   │   └── middleware/             # Middleware
│   ├── pkg/
│   │   ├── database/               # Подключение к БД
│   │   └── utils/                  # Утилиты
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── src/
│   │   ├── components/             # React компоненты
│   │   ├── pages/                  # Страницы
│   │   ├── services/               # API сервисы
│   │   ├── store/                  # Zustand stores
│   │   └── types/                  # TypeScript типы
│   ├── Dockerfile
│   ├── package.json
│   └── vite.config.ts
├── docker-compose.yml
├── init-mongo.js
└── README.md
```

## Быстрый старт

### Предварительные требования
- Docker и Docker Compose
- Node.js 18+ (для разработки frontend)
- Go 1.21+ (для разработки backend)

### Запуск с Docker Compose

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd trucking-amortization
```

2. Запустите все сервисы:
```bash
docker-compose up -d
```

3. Приложение будет доступно по адресам:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - MongoDB: localhost:27017

### Разработка

#### Backend
```bash
# Установка зависимостей
go mod tidy

# Запуск в режиме разработки
go run cmd/main.go
```

#### Frontend
```bash
cd frontend

# Установка зависимостей
npm install

# Запуск в режиме разработки
npm run dev
```

## API Endpoints

### Аутентификация
- `POST /api/auth/register` - Регистрация пользователя
- `POST /api/auth/login` - Вход в систему
- `GET /api/auth/me` - Получение информации о текущем пользователе

### Компании
- `GET /api/companies` - Список компаний пользователя
- `POST /api/companies` - Создание компании
- `PUT /api/companies/:id` - Обновление компании
- `DELETE /api/companies/:id` - Удаление компании

### Активы
- `GET /api/assets?company_id=` - Список активов компании
- `POST /api/assets` - Создание актива
- `PUT /api/assets/:id` - Обновление актива
- `DELETE /api/assets/:id` - Удаление актива

### Расписания
- `GET /api/schedules/amortization/:asset_id` - График амортизации
- `GET /api/schedules/depreciation/:asset_id` - График депрециации
- `GET /api/schedules/business-debt/:company_id` - Общая задолженность

## Модели данных

### User (Пользователь)
```go
type User struct {
    ID        string    `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    Companies []string  `json:"companies"`
    CreatedAt time.Time `json:"created_at"`
}
```

### Company (Компания)
```go
type Company struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    OwnerID   string    `json:"owner_id"`
    CreatedAt time.Time `json:"created_at"`
}
```

### Asset (Актив)
```go
type Asset struct {
    ID            string    `json:"id"`
    CompanyID     string    `json:"company_id"`
    Type          string    `json:"type"` // "truck" или "trailer"
    Make          string    `json:"make"`
    Model         string    `json:"model"`
    Year          int       `json:"year"`
    VIN           string    `json:"vin"`
    PurchaseDate  time.Time `json:"purchase_date"`
    PurchasePrice float64   `json:"purchase_price"`
    LoanInfo      *LoanInfo `json:"loan_info"`
    CreatedAt     time.Time `json:"created_at"`
}
```

### LoanInfo (Информация о кредите)
```go
type LoanInfo struct {
    LoanAmount     float64   `json:"loan_amount"`
    InterestRate   float64   `json:"interest_rate"`
    LoanTerm       int       `json:"loan_term"` // в месяцах
    StartDate      time.Time `json:"start_date"`
    MonthlyPayment float64   `json:"monthly_payment"`
}
```

## Переменные окружения

### Backend
- `MONGO_URI` - URI подключения к MongoDB
- `JWT_SECRET` - Секретный ключ для JWT
- `PORT` - Порт сервера (по умолчанию 8080)
- `DATABASE_NAME` - Имя базы данных

### Frontend
- `VITE_API_URL` - URL backend API

## Безопасность

- JWT токены для аутентификации
- CORS настройки
- Валидация входных данных
- Хеширование паролей с bcrypt
- Rate limiting (планируется)

## Мониторинг и логирование

- Структурированные JSON логи
- Health check endpoint: `GET /health`
- Middleware для логирования запросов
- Обработка ошибок

## Развертывание

### Production
1. Настройте переменные окружения
2. Используйте Docker Compose для production:
```bash
docker-compose -f docker-compose.prod.yml up -d
```

### CI/CD
Настроен GitHub Actions для автоматического развертывания (см. `.github/workflows/deploy.yml`)

## Лицензия

MIT License

## Поддержка

Для вопросов и предложений создавайте issues в репозитории. 