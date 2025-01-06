CREATE TABLE users (
    id SERIAL PRIMARY KEY,        -- Уникальный идентификатор пользователя, автоматически увеличивается
    name VARCHAR(255) NOT NULL,   -- Имя пользователя, текстовое поле, обязательное
    email VARCHAR(255) NOT NULL,  -- Email пользователя, текстовое поле, обязательное
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Время создания записи, значение по умолчанию — текущее время
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Время последнего обновления
);
