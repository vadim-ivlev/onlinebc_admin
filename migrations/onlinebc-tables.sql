
SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;



CREATE TABLE IF NOT EXISTS public.broadcast (
    id integer NOT NULL,
    title character varying(256),
    time_created integer,
    time_begin integer,
    is_ended integer,
    show_date integer,
    show_time integer,
    is_yandex integer,
    yandex_ids character varying(255),
    show_main_page integer,
    link_article character varying(256),
    link_img character varying(255),
    groups_create integer,
    is_diary integer,
    diary_author character varying(255)
);



CREATE TABLE IF NOT EXISTS public.post (
    id integer NOT NULL,
    id_parent integer,
    id_broadcast integer,
    text text,
    post_time integer,
    post_type integer,
    link character varying(256),
    has_big_img integer,
    author text
);



CREATE TABLE IF NOT EXISTS public.medium (
    id integer NOT NULL,
    post_id integer NOT NULL,
    uri character varying(255),
    thumb character varying(255),
    source character varying(255)
);



CREATE SEQUENCE IF NOT EXISTS public.broadcast_id_seq
    START WITH 400
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    OWNED BY public.broadcast.id
    CACHE 1;
-- ALTER SEQUENCE public.broadcast_id_seq OWNED BY public.broadcast.id;



CREATE SEQUENCE IF NOT EXISTS public.post_id_seq
    START WITH 30000
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    OWNED BY public.post.id
    CACHE 1;
-- ALTER SEQUENCE public.post_id_seq OWNED BY public.post.id;



CREATE SEQUENCE IF NOT EXISTS public.medium_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    OWNED BY public.medium.id
    CACHE 1;
-- ALTER SEQUENCE public.medium_id_seq OWNED BY public.medium.id;


ALTER TABLE ONLY public.broadcast ALTER COLUMN id SET DEFAULT nextval('public.broadcast_id_seq'::regclass);

ALTER TABLE ONLY public.post ALTER COLUMN id SET DEFAULT nextval('public.post_id_seq'::regclass);

ALTER TABLE ONLY public.medium ALTER COLUMN id SET DEFAULT nextval('public.medium_id_seq'::regclass);




ALTER TABLE ONLY public.broadcast
    ADD CONSTRAINT broadcast_pkey1 PRIMARY KEY (id);

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey1 PRIMARY KEY (id);

ALTER TABLE ONLY public.medium
    ADD CONSTRAINT medium_pk PRIMARY KEY (id);



CREATE INDEX IF NOT EXISTS medium_post_id_idx ON public.medium USING btree (post_id);

CREATE INDEX IF NOT EXISTS post_id_broadcast_idx ON public.post USING btree (id_broadcast);

CREATE INDEX IF NOT EXISTS post_id_parent_idx ON public.post USING btree (id_parent);



ALTER TABLE ONLY public.medium
    ADD CONSTRAINT medium_post_fk FOREIGN KEY (post_id) REFERENCES public.post(id) ON DELETE CASCADE DEFERRABLE;


ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_broadcast_fk FOREIGN KEY (id_broadcast) REFERENCES public.broadcast(id) ON DELETE CASCADE DEFERRABLE;


ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_post_fk FOREIGN KEY (id_parent) REFERENCES public.post(id) ON DELETE CASCADE DEFERRABLE;

