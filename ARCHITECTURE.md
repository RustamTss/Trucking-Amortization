# Архитектура Trucking Amortization Platform

## Обзор системы

Trucking Amortization Platform - это полноценное веб-приложение для управления амортизацией и депрециацией в траковом бизнесе, построенное по микросервисной архитектуре.

## Архитектурная диаграмма

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│                 │    │                 │    │                 │
│    Frontend     │◄──►│     Backend     │◄──►│    MongoDB      │
│   (React/TS)    │    │   (Go/Fiber)    │    │   (Database)    │
│                 │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│     Nginx       │    │   JWT Auth      │    │   Data Models   │
│   (Reverse      │    │  Middleware     │    │   & Indexes     │
│    Proxy)       │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## Компоненты системы

### 1. Frontend (React + TypeScript)

**Технологии:**
- React 18 с TypeScript
- Vite для сборки
- Tailwind CSS для стилей
- Zustand для управления состоянием
- React Router для навигации
- Axios для HTTP запросов

**Структура:**
```
frontend/src/
├── components/          # Переиспользуемые компоненты
├── pages/              # Страницы приложения
├── services/           # API сервисы
├── store/              # Zustand stores
├── types/              # TypeScript типы
└── utils/              # Утилиты
```

**Основные страницы:**
- Аутентификация (Login/Register)
- Dashboard (обзор компаний)
- Управление компаниями
- Управление активами
- Графики амортизации
- Графики депрециации
- Анализ задолженности

### 2. Backend (Go + Fiber)

**Технологии:**
- Go 1.21
- Fiber v2 (веб-фреймворк)
- MongoDB Driver
- JWT для аутентификации
- bcrypt для хеширования паролей

**Архитектурные слои:**

#### Handlers (Контроллеры)
- `AuthHandler` - аутентификация и авторизация
- `CompanyHandler` - управление компаниями
- `AssetHandler` - управление активами
- `ScheduleHandler` - расчет графиков

#### Services (Бизнес-логика)
- `UserService` - операции с пользователями
- `CompanyService` - операции с компаниями
- `AssetService` - операции с активами
- `CalculationService` - расчеты амортизации и депрециации

#### Models (Модели данных)
- `User` - пользователи системы
- `Company` - компании пользователей
- `Asset` - активы (траки и трейлеры)
- `LoanInfo` - информация о кредитах
- `Schedule` - графики платежей

#### Middleware
- `AuthMiddleware` - проверка JWT токенов
- `CORSMiddleware` - настройки CORS
- `LoggerMiddleware` - логирование запросов

### 3. База данных (MongoDB)

**Коллекции:**
- `users` - пользователи
- `companies` - компании
- `assets` - активы

**Индексы:**
- `users.email` (unique)
- `companies.owner_id`
- `assets.company_id`
- `assets.vin` (unique)

## Бизнес-логика

### 1. Расчет амортизации кредита

Используется формула аннуитетного платежа:

```
PMT = PV × (r × (1 + r)^n) / ((1 + r)^n - 1)

где:
PMT = ежемесячный платеж
PV = сумма кредита
r = месячная процентная ставка
n = количество месяцев
```

### 2. Расчет депрециации активов

Используется прямолинейный метод:

```
Годовая амортизация = (Первоначальная стоимость - Остаточная стоимость) / Срок полезного использования

Сроки полезного использования:
- Траки: 7 лет
- Трейлеры: 10 лет
- Остаточная стоимость: 10% от первоначальной стоимости
```

### 3. Анализ задолженности

Система рассчитывает:
- Текущий остаток по каждому кредиту
- Общую сумму задолженности по компании
- Ежемесячные платежи
- Количество оставшихся месяцев

## Безопасность

### Аутентификация и авторизация
- JWT токены с временем жизни 24 часа
- Хеширование паролей с bcrypt
- Middleware для проверки токенов

### Валидация данных
- Валидация на уровне моделей
- Санитизация входных данных
- Проверка прав доступа к ресурсам

### Безопасность API
- CORS настройки
- Rate limiting (планируется)
- HTTPS в production

## Развертывание

### Docker контейнеры
- `trucking_backend` - Go приложение
- `trucking_frontend` - React приложение с Nginx
- `trucking_db` - MongoDB

### Переменные окружения
```bash
# Backend
MONGO_URI=mongodb://admin:password@mongodb:27017/trucking_db?authSource=admin
JWT_SECRET=your_secret_key
PORT=8080

# Frontend
VITE_API_URL=http://localhost:8080
```

### CI/CD Pipeline
1. **Test** - запуск тестов и линтеров
2. **Build** - сборка Docker образов
3. **Push** - публикация в GitHub Container Registry
4. **Deploy** - развертывание на production сервере

## Мониторинг и логирование

### Логирование
- Структурированные JSON логи
- Логирование всех HTTP запросов
- Логирование ошибок с контекстом

### Мониторинг
- Health check endpoint: `GET /health`
- Метрики производительности (планируется)
- Алерты при ошибках (планируется)

## Масштабирование

### Горизонтальное масштабирование
- Stateless backend сервисы
- Load balancer для распределения нагрузки
- Репликация MongoDB

### Кэширование
- Redis для кэширования сессий (планируется)
- CDN для статических ресурсов
- Кэширование расчетов

## Будущие улучшения

### Функциональность
- Экспорт отчетов в PDF/Excel
- Уведомления о платежах
- Интеграция с банковскими API
- Мобильное приложение

### Техническая часть
- Микросервисная архитектура
- Event-driven архитектура
- GraphQL API
- Kubernetes для оркестрации

## Производительность

### Оптимизации
- Индексы в MongoDB для быстрых запросов
- Пагинация для больших списков
- Lazy loading компонентов
- Сжатие статических ресурсов

### Метрики
- Время отклика API < 200ms
- Время загрузки страницы < 2s
- Доступность 99.9%

## Тестирование

### Backend
- Unit тесты для бизнес-логики
- Integration тесты для API
- Тесты производительности

### Frontend
- Unit тесты компонентов
- E2E тесты критических путей
- Тесты доступности

## Документация

- API документация (Swagger/OpenAPI)
- Руководство пользователя
- Техническая документация
- Диаграммы архитектуры 