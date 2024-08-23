# TelloService

The following components and tools are used in this API:

- [GoLang](https://www.go.dev/)
- [GORM](https://www.gorm.io/)

## Pre-requisites

The following is required to run the service locally:

- Docker

## Start up

The following command can be used to run the application locally

```bash
docker-compose up
```

or 

```bash
docker-compose up --build --remove-orphans --force-recreatess
```

or if you have nodemon installed globally ( local only )

```bash
nodemon
```

## Database Seeding
Seeding is done automatically but you can still run it with the command below:
```bash
go run ./seeding/main.go 
```
