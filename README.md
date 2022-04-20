# be_brankas_test


### Table of Contents
- [Description](#description)
- [Project setup](#project-setup)
- [Contributor](#contributor)

## Description
This repository is back-end software engineer test application exclusive to (Nguyen Duy Kien) for job offer decision. 
This application allows to upload image < 8Mbs, store image metadata into database

# Project setup

### Technologies
- Golang
- "github.com/lib/pq"
- "github.com/dgrijalva/jwt-go"
- PostgreSQL (pgAdmin4)

## To build
- go get -u github.com/dgrijalva/jwt-go // JWT library
- go get -u github.com/lib/pq           // Postgres library
- go build .

## To start
- go run .

## To create schema and table (POSTGRE SQL)
- CREATE SCHEMA brankas_test
- CREATE TABLE brankas_test.images (
	id serial NOT NULL,
	file_name varchar(255) NOT NULL,
	file_size int8 NOT NULL,
	content_type varchar(500) NOT NULL,
	created_date timestamp NOT NULL default(now()),
	CONSTRAINT images_pkey PRIMARY KEY (id)
);

================
To change connection string to database
Change constants in database/database.go


## Contributor
Kien Duy Nguyen
