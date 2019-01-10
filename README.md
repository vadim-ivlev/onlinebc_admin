# Онлайн трансляции - админка / onlinebc_admin

REST и GraphQL API для редактирования БД onlinebc


## Установка

Клонировать проект в  `~/go/src` и перейти в `onlinebc_admin/`

    git clone git@git.rgwork.ru:web/onlinebc_adin.git ~/go/src
    cd ~/go/src/onlinebc_admin

## Запуск БД 
Запустить БД
    
    docker-compose up -d    

Восстановить тестовую БД из дампа

    docker exec -it psql-com psql -U root -1 -d onlinebc -f /dumps/onlinebc-dump.sql


Postgres доступен на localhost:5432.

Аdminer - в браузере http://localhost:8080. 

Параметры доступа:
- System: PostgreSQL,
- Server: db,
- Username: root,
- Password: root,
- Database: onlinebc





## Запуск приложения

Если установлен go

    go run main.go

Если go не установлен, то под Linux

    ./onlinebc

Программа выдаст список возможных параметров запуска. Для запуска web приложения запустите с ключом `-serve`
    
    go run main.go -serve


--------------------

## Полезные команды




Запуск docker-compose

    docker-compose up -d



Останов docker-compose

    docker-compose down



Удаление базы данных после останова docker-compose

    sudo rm -rf  pgdata




Восстановление БД из дампа. Находится в `migrations/`.

    docker exec -it psql-com psql -U root -1 -d onlinebc -f /dumps/onlinebc-dump.sql



Дамп БД в файл в `migrations/`.
  
    docker exec -it psql-com pg_dump --file /dumps/onlinebc-dump.sql --host "localhost" --port "5432" --username "root"  --verbose --format=p --create --clean --if-exists --dbname "onlinebc"


Дамп схемы БД

    docker exec -it psql-com pg_dump --file /dumps/onlinebc-schema.sql --host "localhost" --port "5432" --username "root" --schema-only  --verbose --format=p --create --clean --if-exists --dbname "onlinebc"


Дамп только данных таблиц.

    docker exec -it psql-com pg_dump --file /dumps/onlinebc-data.sql --host "localhost" --port "5432" --username "root"  --verbose --format=p --dbname "onlinebc" --column-inserts --data-only --table=broadcast --table=post --table=medium


Можно добавить  -$(date +"-%Y-%m-%d--%H-%M-%S") к имени файла для приклеивания штампа даты-времени.


Показ структуры таблицы TABLE_NAME

    docker exec -it psql-com pg_dump -U root -d onlinebc -t TABLE_NAME --schema-only



Командная строка Postgres

	docker exec -it psql-com psql -U root onlinebc


