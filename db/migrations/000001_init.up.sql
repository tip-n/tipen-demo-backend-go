CREATE TABLE IF NOT EXISTS restaurant_menu_type (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    is_deleted BOOLEAN DEFAULT false,
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    type varchar(255) not null
);

CREATE TABLE IF NOT EXISTS restaurant_menu (
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP,
   is_deleted BOOLEAN DEFAULT false,
   id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
   name varchar(255) not null,
   slug varchar(255) not null,
   price varchar(50) not null,
   description varchar(255) not null,
   type_id int not null
);