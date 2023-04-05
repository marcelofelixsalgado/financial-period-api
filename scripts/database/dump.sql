CREATE DATABASE IF NOT EXISTS financial_period_db;
USE financial_period_db;

DROP TABLE IF EXISTS tenants;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS user_credentials;
DROP TABLE IF EXISTS periods;

CREATE TABLE tenants(
    id varchar(255) primary key,
    created_at datetime not null
);

CREATE TABLE users(
    id varchar(255) primary key,
    tenant_id varchar(255) not null,
    name varchar(255) not null,
    phone varchar(50) not null,
    email varchar(255) not null,
    created_at datetime not null,
    updated_at datetime default '0001-01-01 00:00:00',

    CONSTRAINT UC_Phone UNIQUE (phone),
    CONSTRAINT UC_Email UNIQUE (email),

    CONSTRAINT FK_UserTenant FOREIGN KEY (tenant_id)
    REFERENCES tenants(id) ON DELETE CASCADE    
);

CREATE TABLE user_credentials(
    id varchar(255) primary key,
    user_id varchar(255) not null,
    password varchar(255) not null,
    created_at datetime not null,
    updated_at datetime default '0001-01-01 00:00:00',

    CONSTRAINT FK_UserCredentials FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE transaction_types(
    code varchar(50) primary key,
    name varchar(255) not null
);

CREATE TABLE categories(
    id varchar(255) primary key,
    tenant_id varchar(255) not null,
    code varchar(50) not null,
    name varchar(255) not null,
    transaction_type_code varchar(50) not null,
    created_at datetime not null,
    updated_at datetime default '0001-01-01 00:00:00',

    CONSTRAINT UC_Category UNIQUE (tenant_id, code),
    CONSTRAINT FK_CategoryTransactionType FOREIGN KEY (transaction_type_code)
    REFERENCES transaction_types(code) ON DELETE CASCADE    
);

CREATE TABLE subcategories(
    id varchar(255) primary key,
    tenant_id varchar(255) not null,
    code varchar(50) not null,
    name varchar(255) not null,
    category_id varchar(255) not null,
    created_at datetime not null,
    updated_at datetime default '0001-01-01 00:00:00',

    CONSTRAINT UC_SubCategory UNIQUE (tenant_id, code),
    CONSTRAINT FK_SubCategoryCategory FOREIGN KEY (category_id)
    REFERENCES categories(id) ON DELETE CASCADE    
);

CREATE TABLE periods(
    id varchar(255) primary key,
    tenant_id varchar(255) not null,
    code varchar(50) not null,
    name varchar(255) not null,
    year varchar(4) not null,
    start_date datetime not null,
    end_date datetime not null,
    created_at datetime not null,
    updated_at datetime default '0001-01-01 00:00:00',

    CONSTRAINT UC_Period UNIQUE (tenant_id, code)
);

CREATE TABLE balance(
	id              varchar(255) primary key,
	tenant_id       varchar(255) not null,
	period_id       varchar(255) not null,
	category_id     varchar(255) not null,
	actual_amount   float,
	limit_amount    float,
    created_at      datetime not null,
    updated_at      datetime default '0001-01-01 00:00:00',

    CONSTRAINT UC_Balance UNIQUE (tenant_id, period_id, category_id)
);

INSERT INTO transaction_types (code, name) values ('EARNING', 'Receita');
INSERT INTO transaction_types (code, name) values ('EXPENSE', 'Despesa');
INSERT INTO transaction_types (code, name) values ('TRANSFER', 'Transferência');