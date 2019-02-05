Онлайн трансляции (onlinebc)
=============================

Рабочая среда для фронтэнд разработчиков. 
-----------------------------------------


- Тестовая страница API: <http://localhost:7700/> `GET`.

- JSON описание конечных точек GraphQL и REST API с параметрами запросов: <http://localhost:7700/routes> `GET`.

- Конечная точка GraphQL <http://localhost:7700/graphql> `POST`.
Документацию по типам данных, и возможным парамтрам запросов можно получить стандартными средствами GraphQL.




Требования к ПО
----------
- Docker, Docker-compose.



Запуск 
------

    docker-compose -f docker-compose-frontend.yml up -d


Просмотр логов работающего приложения.

    docker-compose -f docker-compose-frontend.yml logs -f runner 


Контроль запуска и доступности API 
------------------------

В браузере откройте тестовую страницу доступного API: <http://localhost:7700>. 


Останов программы
-----------

     docker-compose -f docker-compose-frontend.yml down



Дополнительная информация
--------------------------
<https://git.rgwork.ru/web/onlinebc_admin>




