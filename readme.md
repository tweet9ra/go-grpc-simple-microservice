# Простой микросервис на GO
### Технологии
- GO 1.14
- gRPC
- Protobuf 3
- Postgres 9.6 [(dockerhub)](https://hub.docker.com/_/postgres)
- Docker

### Запуск
1) Создать `.env` файл из шаблона: `cp src/.env.example src/.env`
2) Запустить postgres через docker-compose `docker-compose up -d`, по дефолту использует локальный порт `5433`
3) Собрать Docker image сервиса `docker build -t go-crud-microservice .`
4) Запустить image сервиса, например: `docker run --publish 4114:8080 --link simple-microservice_db_1:db --net=simple-microservice_default -t go-crud-microservice`  
Возможно нужно будет заменить: `simple-microservice_db_1` на имя контейнера с postgres (можно узнать через `docker-compose ps`) и сеть `go-crud-microservice`(смотреть в `docker network ls`)
5) Подключиться к gRPC с локальным портом `4114` используя файл `src/crud.proto`
Простая утилита для тестов [Evans](https://github.com/ktr0731/evans)

### Функционал
1) List: получить одну запись по ID или множество с фильтром по ID производителя
2) Create: создать одну или несколько записей
3) Delete: удалить одну или несколько записей
