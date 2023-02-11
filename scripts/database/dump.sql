CREATE DATABASE IF NOT EXISTS financial_db;
USE financial_db;

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

CREATE TABLE periods(
    id varchar(255) primary key,
    tenant_id varchar(255) not null,
    code varchar(50) not null,
    name varchar(255) not null,
    year varchar(4) not null,
    start_date datetime not null,
    end_date datetime not null,
    created_at datetime not null,
    updated_at datetime default '0001-01-01 00:00:00'
);


