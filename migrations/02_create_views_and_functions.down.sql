-- REST ------------------------------------
DROP VIEW IF EXISTS public.post_view;
DROP VIEW IF EXISTS public.answer_view;

DROP FUNCTION IF EXISTS public.get_broadcasts();
DROP FUNCTION IF EXISTS public.get_full_broadcast(idd integer);
DROP FUNCTION IF EXISTS public.get_posts(idd integer);
DROP FUNCTION IF EXISTS public.get_answers(idd integer);
DROP FUNCTION IF EXISTS public.get_media(idd integer);

