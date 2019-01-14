# Онлайн трансляции  (onlinebc_admin)

 GraphQL и REST API онлайн трансляций.



Требования к ПО
--------------

На компьютере разработчика должны быть установлены Docker, Docker-compose, Go.



Загрузка файлов проекта
----------------------
    git clone git@git.rgwork.ru:web/onlinebc_admin.git ~/go/src
    cd ~/go/src/onlinebc_admin



Запуск Postgres и Redis
-----------------------    
    docker-compose up -d    



Запуск приложения
-----------------
    go run main.go -serve 7777

Для просморта списка возможных параметров запустите программу без параметров.



Сборка для фронтэнд разработчиков и production
----------------------------------------------

    /build.sh

Скрипт `build.sh` генерирует исполняемые файлы windows (`onlinebc.exe`) и Linux (`onlinebc`) в директорию `build/` вместе с настроечными файлами и `README.md` для фронтэнд разработчиков. 

Из `build/`файлы копируются в директорию `../onlinebc/`, расположенную в том же каталоге что и `onlinebc_admin`, которая содержит клон репозитория <https://git.rgwork.ru/web/onlinebc>. Таким образом обновляются файлы проекта onlinebc для использования фронтэнд разработчиками и для размещения на продакшн сервере. 





-------------------------------------------------------


Просмотр состояния базы данных
------------------------------

Postgres доступен на localhost:5432.

Если блок adminer разкомментирован в `docker-compose.yml`, то в браузере откройте <http://localhost:8080>. 

Параметры доступа:
- System: PostgreSQL,
- Server: db,
- Username: root,
- Password: root,
- Database: onlinebc








Другие команды
--------------------


Останов базы данных
    
    docker-compose down



Удаление файлов базы данных после останова docker-compose

    sudo rm -rf  pgdata



Дамп базы данных в файл в директорию `migrations/`.
  
    docker exec -it psql-com pg_dump --file /dumps/onlinebc-dump.sql --host "localhost" --port "5432" --username "root"  --verbose --format=p --create --clean --if-exists --dbname "onlinebc"


Восстановление БД из дампа в `migrations/`.

    docker exec -it psql-com psql -U root -1 -d onlinebc -f /dumps/onlinebc-dump.sql



Дамп схемы БД

    docker exec -it psql-com pg_dump --file /dumps/onlinebc-schema.sql --host "localhost" --port "5432" --username "root" --schema-only  --verbose --format=p --create --clean --if-exists --dbname "onlinebc"


Дамп только данных таблиц.

    docker exec -it psql-com pg_dump --file /dumps/onlinebc-data.sql --host "localhost" --port "5432" --username "root"  --verbose --format=p --dbname "onlinebc" --column-inserts --data-only --table=broadcast --table=post --table=medium


Можно добавить  -$(date +"-%Y-%m-%d--%H-%M-%S") к имени файла для приклеивания штампа даты-времени.



Показ структуры таблицы TABLE_NAME

    docker-compose exec db pg_dump -U root -d onlinebc -t TABLE_NAME --schema-only



Командная строка Postgres

	docker-compose exec db psql -U root onlinebc



Командная строка Redis

    docker-compose exec redis redis-cli



---------------------


Замечания о программе
=====================





