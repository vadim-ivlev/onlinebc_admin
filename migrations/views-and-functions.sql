-- ФУНКЦИИ **********************************************************************************
-- Назначение функций преобразовть таблицы в JSON требуемый API RG.RU.
-- Выражение array_to_json(array_agg(row_to_json( t, false )),true) преобразует таблицу t JSON объект.
-- Все функции возвращают JSON.


CREATE OR REPLACE FUNCTION public.get_media(idd integer)
 RETURNS json
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN
    (
        select array_to_json(array_agg(row_to_json( t, false )),true) from
        (
            SELECT * FROM media
            WHERE post_id = idd
        ) t
    );

END;
$function$
;


CREATE OR REPLACE FUNCTION public.get_answers(idd integer)
 RETURNS json
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN
    (
        select array_to_json(array_agg(row_to_json( t, false )),true) from
        (
            SELECT * FROM answer_view
            WHERE id_parent = idd
        ) t
    );

END;
$function$
;

CREATE OR REPLACE FUNCTION public.get_posts(idd integer)
 RETURNS json
 LANGUAGE plpgsql
AS $function$
BEGIN   
    RETURN  
    (
        select array_to_json(array_agg(row_to_json( t, false )),true) from
        ( 
            SELECT * FROM post_view
            WHERE 
                    id_broadcast = idd 
                AND id_parent IS NULL 
            ORDER BY id DESC 
        ) t 
    );
    
END;
$function$
;


-- get_broadcast - гланая функция.
-- Преобразует плоские таблицы базы данных в JSON
-- с многоуровненвой иерархической структурой

CREATE OR REPLACE FUNCTION public.get_broadcast(idd integer)
 RETURNS json
 LANGUAGE plpgsql
AS $function$
BEGIN   
    RETURN
    (
        select array_to_json(array_agg(row_to_json( t, false )),true) from
        ( select *, get_posts(id) as posts  from broadcast where id = idd ) t
    );
END;
$function$
;

CREATE OR REPLACE FUNCTION public.get_broadcasts()
 RETURNS json
 LANGUAGE plpgsql
AS $function$
BEGIN   
    RETURN
    (
       select array_to_json(array_agg(row_to_json( t, false )),true) from
       ( select *  from broadcast ) t
    );
END;
$function$
;


CREATE OR REPLACE FUNCTION public.js(t anyelement)
 RETURNS json
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN array_to_json(array_agg(to_json( t )),true);
END;
$function$
;



-- ПРЕДСТАВЛЕНИЯ *****************************************************************************888
-- именуют поля таблицы post в соответствии с полями JSON API RG.RU
-- https://outer.rg.ru/plain/online_translations/api/online.php?id=354
-- Если нужно поменять названия полей, добавить или убрать поле
-- менять нужно здесь.


CREATE OR REPLACE VIEW public.answer_view
AS SELECT post.id,
    post.id_parent,
    post.id_broadcast,
    post.has_big_img,
    post.author     AS posts__answer__author,
    post.text       AS posts__answer__text,
    post.post_type  AS posts__answer__type,
    post.link       AS posts__answer__uri,
    to_char(to_timestamp(post.post_time::double precision), 'DD.MM.YYYY'::text) AS posts__answer__date,
    to_char(to_timestamp(post.post_time::double precision), 'HH24:MI'::text)    AS posts__answer__time,
    
    get_media(post.id)      AS posts__answer__media,
    get_answers(post.id)    AS posts__answer__answers
   FROM post;



CREATE OR REPLACE VIEW public.post_view
AS SELECT post.id,
    post.id_parent,
    post.id_broadcast,
    post.has_big_img,
    post.author AS posts__author,
    post.text AS posts__text,
    post.post_type AS posts__type,
    post.link AS posts__uri,
    to_char(to_timestamp(post.post_time::double precision), 'DD.MM.YYYY'::text) AS posts__date,
    to_char(to_timestamp(post.post_time::double precision), 'HH24:MI'::text) AS posts__time,
    
    get_media(post.id) AS posts__media,
    get_answers(post.id) AS posts__answers
   FROM post;

