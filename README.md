
# Проект: Разработка REST API для управления задачами (To-Do List)
Роль: Junior Backend Developer (Golang)
Цель:
Разработать REST API для системы управления задачами, которая позволяет
пользователям создавать, просматривать, обновлять и удалять задачи.

# Реализация
Реализованы все основные функции:

1. Запуск веб-сервиса
2. Создание БД
3. Добавление задач (POST, http://localhost:port/tasks)
4. Получение списка всех задач (GET, http://localhost:port/tasks)
4. Получение одной задачи (GET, http://localhost:port/tasks/id)
5. Редактирование задач (PUT, http://localhost:port/tasks/id)
6. Удаление задач (DELETE, http://localhost:port/tasks/id)

Реализованы дополнительные функции:

1. Написан пользовательский интерфейс для проверки работы REST API(реализован на ЯП GO с использованием библиотеки FYNE.io)

# Запуск проекта

1. Склонируйте репозиторий на локальную машину или скачайте архив с проектом
2. Откройте проект в редакторе кода, в терминале выполните команду:
```
go mod tidy
```
3. В файле .env находятся переменные окружения для запуска проекта, задайте им значения для вашей локальной машины, если не указывать, будут использованы параметры по умолчанию:

```
# Port on which the application will run.
# Default value is "8080" if not provided.
TASK_PORT=""

# Database user for connecting to the PostgreSQL database.
# Default value is "postgres" if not provided.
TASK_USER=""

# Password for the database user.
# Default value is "root" if not provided.
TASK_PASSWORD=""

# Host where the PostgreSQL database is located.
# Default value is "127.0.0.1" (localhost) if not provided.
TASK_HOST=""

# Port on which the PostgreSQL database is listening.
# Default value is "5432" if not provided.
TASK_DBPORT=""

# SSL mode for the PostgreSQL connection.
# Default value is "disable" if not provided.
TASK_SSLMODE=""
```
4. В терминале, находясь в основной директории проекта выполните команду:
```
go run cmd/main.go
```
5. БД будет создана автоматически. Название БД: verbatasks, название таблицы: tasks.
# Проверка работоспособности 
## Проверка через пользовательский интерфейс
1. Запустите файл verba_task_manager.exe для Windows или verba_taskmanager_Linux для Linux.
2. В открывшемся окне укажите порт, на котором запущен сервер, нажмите кнопку Set Port.
3. Добавления задачи: Введите заголовок, описание задачи, дату в формате RFC3339(2000-09-01T12:40:00Z). Нажмите кнопку Add Task, при успешном исполнении в окне Result вернется код ответа и задача. В случае ошибки вернется код ответа и описание ошибки.
4. Получения всех задач: Нажмите кнопку Get All Tasks, при успешном исполнении в окне Result вернется код ответа и все задачи. В случае ошибки вернется код ответа и описание ошибки.
5. Получения задачи по ID. Введите ID задачи, нажмите кнопку Get Task By ID, при успешном исполнении в окне Result вернется код ответа и задача. В случае ошибки вернется код ответа и описание ошибки.
6. Редактирование задачи. Введите ID, заголовок, описание, дату в формате RFC3339(2000-09-01T12:40:00Z). Нажмите кнопку Update Task By ID, при успешном исполнении в окне Result вернется код ответа и задача.В случае ошибки вернется код ответа и описание ошибки.
7. Удаление задачи по ID. Введите ID задачи, нажмите кнопук Delete Task By ID, при успешном исполнении в окне Result вернется код ответа и сообщение об удачном удалении. В случае ошибки вернется код ответа и описание ошибки.
## Проверка через терминал
1. Добавления задачи:
```
curl -i -X POST -H "Content-Type: application/json" -d '{"title": "Title", "description": "Description", "due_date": "2024-08-28T14:58:11+03:00"}' http://localhost:8080/tasks
```
 2. Получение всех задач:
 ```
 curl -i -X GET -H "Content-type: application/json" http://localhost:8080/tasks
 ```
 3. Получение задачи по ID:
 ```
 curl -i -X GET -H "Content-type: application/json" http://localhost:8080/tasks/1
 ```
 4. Редактирование задачи:
 ```
 curl -i -X PUT -H "Content-Type: application/json" -d '{"title": "Update", "description": "Update", "due_date": "2024-08-28T14:58
:11+03:00"}' http://localhost:8080/tasks/9
```
5. Удаление задачи по ID: 
```
curl -i -X DELETE -H "Content-type: application/json" http://localhost:8080/tasks/9
```

