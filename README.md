# Онлайн трансляции  (onlinebc_admin)

GraphQL и REST API онлайн трансляций.

- Тестовая страница API: <http://localhost:7777/> `GET`.

- JSON описание конечных точек GraphQL и REST API с параметрами запросов: <http://localhost:7777/routes> `GET`.

- Конечная точка GraphQL <http://localhost:7777/graphql> `POST`.
Документацию по типам данных, и возможным парамтрам запросов можно получить стандартными средствами GraphQL.


Требования к ПО
--------------

На компьютере разработчика должны быть установлены Docker, Docker-compose, Go.



Клонирование проекта на локальный компьютер
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



Контроль запуска и доступности API 
-----------------------------------

В браузере откройте тестовую страницу доступного API: <http://localhost:7777>. 





Сборка для фронтэнд разработчиков и production
----------------------------------------------

    /build.sh

Скрипт `build.sh` генерирует исполняемые файл Linux (`onlinebc`) в директории `build/` вместе с настроечными файлами. Эта директория должна копироваться на продакшн сервер. 

Описание настроек в файле <readme-production.md>.


Из `build/`файлы копируются в директорию `../onlinebc/`, расположенную в том же каталоге что и `onlinebc_admin`, которая содержит клон репозитория <https://git.rgwork.ru/web/onlinebc>. Таким образом обновляются файлы проекта onlinebc для использования фронтэнд разработчиками.  

Читай <readme-front.md>



---------------------


Замечания о программе
=====================

Данные
-------

Для хранения данных используется СУБД Postgres v11. 


<img src="docs/tables.png">

Таблицы БД восстанавливаются и наполняются тестовыми данными при каждом запуске приложения.
- broadcast  - трансляции
- post  - посты к трансляциям. Таблица рекурсивно ссылается на саму себя для организации ответов к посту.
- medium - изображения к постам

Для обеспечения ссылочной целостности БД на таблицы `post` и `medium` наложены ограничения внешних ключей с каскадным удалением из подчиненных таблиц. Первичные ключи `id` автогенерируются. На все ключи построены индексы.


Описание полей таблиц находятся в файле `controller/controller-graphql.go`. Во время исполнеия могут быть получены стандартными средствами GraphQL. 




Файлы и директории
-------------------



    configs/

Содержит настроечные файлы соединений с Postgres и Redis. Файл `routes.yaml` описывает маршруты и возможные параметры запросов для тестовой страницы API <http://localhost:7777/>. Файл `routes-front.yaml` - усеченная версия файла `routes.yaml`, содержит маршруты публичного REST API для показа на страницах RG.RU, подключается вместо  `routes.yaml` если при запуске программы задан параметр `-front`.




    controller/

Файл `controller-graphql.go` содежит функции GraphQL API, `controller-rest.go` функции REST API.




    middleware/

Каждый запрос к программе обрабатывается двумя функциями middleware до того как будет обработан основным контроллером.

```
Жизненный цикл запроса

(req) --> HeadersMiddleware --> RedisMiddleware --> router --> controller --> (resp)
```
`HeadersMiddleware` добавляет json и возможно CORS HTTP заголовки к ответу сервера.  `RedisMiddleware` кэширует ответы сервера для публичных REST маршрутов начинающихся с `/api/` и возвращает кэшированные ответы клиенту.




    model/
        db/
        redis/

Содержит файлы отвечающие за работу с базой данных и редис. Формируют запросы к базам данных и формирование ответов в JSON формате.

Приложение не использует библиотек ORM. Формирование ответов сервера в JSON формате в части REST API возложено на представления и функции postgres. Таким образом снижается объем трафика между БД и приложением, ускоряются запросы, снижается нагрузка на  хостирующий сервер, и уменьшается объем кода. 

Для формирование JSON в части GraphQL используется библиотека <github.com/jmoiron/sqlx>, что позволило применить отображения вместо структур и точного описания типов БД в коде Go. Это уменьшает связность кода БД и приложения, позволяет менять структуру объектов базы данных без необходимости вносить соответствующие изменения в код go приложения и уменьшает суммарный объем кода.



    router/

Cопосталяет маршруты функциям-контроллерам, присоединяет middleware, и запускает сервер. Сами маршруты с именами соответствующих функций вынсены в настроечные файлы `configs/routes.yaml` и `routes-front.yaml`.


    migrations/
        create-tables.sql
        create-views-and-functions.sql
        add-data.sql

Директория содержит SQL скрипты для порождения таблиц базы данных если таковые отсутствуют, а так же представлений и функций. Скрипт `add-data.sql` наполняет базу данных тестовыми данными с идентификаторами трансляций между `321` и `354`. Все три скрипта выполняются при каждом запуске программы, так что программа будет корректно работать даже при изначально пустой базе данных.



**Второстепенные файлы**


    etc/
        .pgpass

Файл используется докер контейнером db Postgres, для того чтобы не вводить пароли при дампе и восстановлении базы данных.


    templates/


Шаблоны приветственного сообщения приложения и тестовой страницы API <http://localhost:7777/>.


    docs/

Файлы для документации и проч.


    pgdata/

Директория где postgres хранит файлы базы данных. Может быть удалена. Восстанавливается при каждом новом запуске приложения.

    docker-compose.yml     
    main.go
    README-FRONT.md              # Для фронтэнд разработчиков
    README-PRODUCTION.md         # Для админов
    README.md                    # Этот файл
    build.sh*                    # Скрипт сборки
    run_tests.sh*                # Запуск тестов
    Dockerfile-frontend          # Используется в build.sh
    docker-compose-frontend.yml  # Файл запуска для фронтэнд разработчиков. 
                                 # Копируется в onlinebc скриптом build.sh
    TODO.md                      # Недоделки








-------------------------------------------------------

Другие команды
--------------------


Просмотр состояния базы данных


Postgres доступен на localhost:5432.

Если блок adminer разкомментирован в `docker-compose.yml`, то в браузере откройте <http://localhost:8080>. 

Параметры доступа:
- System: PostgreSQL,
- Server: db,
- Username: root,
- Password: root,
- Database: onlinebc




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











