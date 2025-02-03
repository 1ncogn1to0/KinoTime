CREATE TABLE roles (
                       id SERIAL PRIMARY KEY,        -- Уникальный идентификатор роли, автоматически увеличивается
                       name VARCHAR(255) NOT NULL,   -- Имя роли, текстовое поле, обязательное
                       code VARCHAR(255) NOT NULL    -- Код роли, обязательно
);