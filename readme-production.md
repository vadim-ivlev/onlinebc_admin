Онлайн трансляции (onlinebc)
=============================

GraphQL и REST API онлайн трансляций.



Конфигурация
-------------
Файлы 
    
    configs/db.yaml
    configs/redis.yaml
    configs/imgserver.yaml

Должны содержать параметры подключения к базам данных Postgres Redis Imageserver соответственно. Структура файлов должна быть аналогична другим файлам в этой же директории.



Запуск программы
--------------

Под Linux

    ./onlinebc -serve 7777 -front


Параметр -serve 7777 задает порт, -front ограничивает API публичными функциями REST для RG.RU. Для получения полного списока параметров запуска, запустите программу без параметров.



Контроль запуска и доступности API 
-----------------------------------

В браузере откройте тестовую страницу доступного API: <http://localhost:7777>. 


Останов программы
-----------
    CTRL-C



Дополнительная информация
--------------------------
<https://git.rgwork.ru/web/onlinebc_admin>



