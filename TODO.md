
TODO
====

- доделать выборку GetBroadcastList
- построить полнотекстовые индексы на full_broadcast

- Заменить загрузку изображений не через base64.

- Прикрутить профилирование pprof https://habr.com/ru/company/badoo/blog/301990/






----------------
```sql

DROP FUNCTION list_broadcast(anyelement);

CREATE OR REPLACE FUNCTION list_broadcast(idd anyelement )
-- RETURNS json
RETURNS TABLE 
(
    id int4 ,
    title varchar(256) ,
    time_created int4 ,
    time_begin int4 ,
    is_ended int4 ,
    show_date int4 ,
    show_time int4 ,
    show_main_page int4 ,
    link_article varchar(256) ,
    link_img varchar(255) ,
    groups_create int4 ,
    is_diary int4 ,
    diary_author varchar(255) 
    
)
 LANGUAGE plpgsql
AS $function$
BEGIN   
    RETURN QUERY
    (  
--        select array_to_json(array_agg(row_to_json( t, false )),true) from
        ( select *, get_posts(id) as posts  from broadcast where id = idd) 
--        t
    );
END;
$function$
;


SELECT public.list_broadcast(354);

```