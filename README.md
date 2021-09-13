# Go-Todo

Это простой пример Todo приложения, предоставляющего REST API для модели.

Входная точка приложения находится в `cmd/web/main.go`

Все роуты находятся в `cmd/web/routes.go`

Все обработчики находятся в `cmd/web/handlers.go`

Модель Todo находится в `pkg/models/models.go`

Функции для работы с БД находятся в `pkg/models/pg/todolist.go`

Для работоспособности приложения требуется создать `.env` файл в корне проекта и указать
ряд значений для следующих полей

```dotenv
PORT=
POSTGRES_NAME=
POSTGRES_PASSWORD=
POSTGRES_DATABASE=
```

## Запуск приложения

Требуется запустить файл `cmd/web/main.go` с помощью команды `go run ./web/main.go`


#REST API описан ниже

## Получить список всех todo задач

### Request

`GET /`

    curl -i -H 'Accept: application/json' http://localhost:8080/

### Response

    &{1 task 1 false 2021-09-08 21:07:12.688493 +0000 +0000}
    &{2 task 2 true 2021-09-08 21:07:12.688493 +0000 +0000}
    &{5 task3 true 2021-09-08 21:40:20.961389 +0000 +0000}


## Создать новую задачу

`POST /todo/create`

    curl -i -H 'Accept: application/json' -d 'title=Learn golang&isCompleted=false' http://localhost:8080/todo/create

### Response

    &{1 Learn golang 1 false 2021-09-08 21:07:12.688493 +0000 +0000}


## Удалить задачу

`DELETE /todo/`

    curl -i -H 'Accept: application/json' -d 'id=1' http://localhost:8080/todo/create


## Обновить задачу

`PATCH /todo/`

    curl -i -H 'Accept: application/json' -d 'id=1&isCompleted=true' http://localhost:8080/todo/create

### Response

    &{1 Learn golang 1 true 2021-09-08 21:07:12.688493 +0000 +0000}