CREATE TABLE IF NOT EXISTS users (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    is_deleted BOOLEAN DEFAULT false,
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    firstname varchar(255) not null,
    lastname varchar(255),
    password varchar(255) not null,
    email varchar(255) not null
);

CREATE TABLE IF NOT EXISTS user_logins (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    is_deleted BOOLEAN DEFAULT false,
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id int not null
);

-- CREATE TABLE IF NOT EXISTS products (
--     created_at TIMESTAMP,
--     updated_at TIMESTAMP,
--     deleted_at TIMESTAMP,
--     is_deleted BOOLEAN DEFAULT false,
--     id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
--     user_id int not null,
--     product_name varchar(255) not null,
--     product_image varchar(255) not null,

-- )