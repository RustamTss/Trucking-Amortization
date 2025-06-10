// Инициализация базы данных MongoDB
db = db.getSiblingDB('trucking_db');

// Создаем коллекции
db.createCollection('users');
db.createCollection('companies');
db.createCollection('assets');

// Создаем индексы для оптимизации запросов
db.users.createIndex({ "email": 1 }, { unique: true });
db.companies.createIndex({ "owner_id": 1 });
db.assets.createIndex({ "company_id": 1 });
db.assets.createIndex({ "vin": 1 }, { unique: true });

print('Database initialized successfully!'); 