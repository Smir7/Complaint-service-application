CREATE TYPE user_role AS ENUM ('USER', 'ADMIN');

CREATE TABLE IF NOT EXISTS users(
       id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       user_uuid UUID NOT NULL UNIQUE,
       username TEXT NOT NULL UNIQUE,
       email TEXT,
       phone INT,
       role user_role DEFAULT 'USER',
       password TEXT
    );
CREATE INDEX IF NOT EXISTS idx_id ON users (id);
