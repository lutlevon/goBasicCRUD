# goBasicCRUD
A project to learn Go by making a basic CRUD application, using the Chi library to manage the HTTP router.

# Docker

To create the docker image use the following command:

    docker build -t crud_golang .

To create the postgresql container:

    docker pull postgres

    docker run -d --name postgresCont -p 5432:5432 -e POSTGRES_PASSWORD=pass123 postgres


# Db migrations 
To create a new migration use the following command:

migrate create -ext sql -dir db/migrations -seq create_users_table

This command will create two files:
    db/migrations/000001_create_users_table.up.sql
    db/migrations/000001_create_users_table.down.sql

The .up.sql file applies the migration, and the .down.sql file reverts the migration.

to apply the migration use the following command: 
    
    migrate -database postgres://user:password@localhost:5432/mydb?sslmode=disable -path db/migrations up

to revet the migration use the following command:

    migrate -database postgres://user:password@localhost:5432/mydb?sslmode=disable -path db/migrations down

