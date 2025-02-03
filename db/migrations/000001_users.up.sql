CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,                          -- Unique identifier, auto-increment
                       name VARCHAR(255) NOT NULL,                     -- User name, required
                       email VARCHAR(255) UNIQUE NOT NULL,             -- Email, required and must be unique
                       password VARCHAR(255) NOT NULL,                 -- Hashed password, required
                       is_confirmed BOOLEAN DEFAULT FALSE,             -- Email confirmation status, default is false
                       confirmation_token VARCHAR(255),                -- Token for email confirmation
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
                       role_id INT DEFAULT 1,                          -- Foreign key to the roles table, default role is "1"
                       CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id) -- Role relationship
);
