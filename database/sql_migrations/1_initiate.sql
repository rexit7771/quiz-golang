-- +migrate Up
-- +migrate StatementBegin

create table kategori(
     id SERIAL PRIMARY KEY,
    name varchar(250),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(50),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by varchar(50)
);

create table buku(
    id SERIAL PRIMARY KEY,
    title varchar(250),
    description varchar(250),
    image_url varchar(250),
    release_year INT,
    price INT,
    total_page INT,
    thickness varchar(250),
    category_id INT REFERENCES kategori(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(50),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by varchar(50)
);



create table users(
    id SERIAL PRIMARY KEY,
    username varchar(50),
    password varchar(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(50),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by varchar(50)
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS buku;
DROP TABLE IF EXISTS kategori;

-- +migrate StatementEnd