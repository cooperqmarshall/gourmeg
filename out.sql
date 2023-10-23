--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.3 (Debian 15.3-0+deb12u1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: link; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.link (
    parent_id bigint NOT NULL,
    child_id bigint NOT NULL,
    child_type character varying(8)
);


ALTER TABLE public.link OWNER TO root;

--
-- Name: list; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.list (
    id bigint NOT NULL,
    name text,
    created_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.list OWNER TO root;

--
-- Name: list_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.list_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.list_id_seq OWNER TO root;

--
-- Name: list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.list_id_seq OWNED BY public.list.id;


--
-- Name: recipe; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.recipe (
    id bigint NOT NULL,
    url text,
    created_at timestamp with time zone DEFAULT now(),
    name text,
    ingredients text[],
    instructions text[]
);


ALTER TABLE public.recipe OWNER TO root;

--
-- Name: recipe_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.recipe_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipe_id_seq OWNER TO root;

--
-- Name: recipe_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.recipe_id_seq OWNED BY public.recipe.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public."user" (
    id bigint NOT NULL,
    username text,
    created_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public."user" OWNER TO root;

--
-- Name: user_user_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.user_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_user_id_seq OWNER TO root;

--
-- Name: user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.user_user_id_seq OWNED BY public."user".id;


--
-- Name: list id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.list ALTER COLUMN id SET DEFAULT nextval('public.list_id_seq'::regclass);


--
-- Name: recipe id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.recipe ALTER COLUMN id SET DEFAULT nextval('public.recipe_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_user_id_seq'::regclass);


--
-- Data for Name: link; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.link (parent_id, child_id, child_type) FROM stdin;
1	2	list
1	1	recipe
2	2	recipe
1	11	recipe
1	12	recipe
1	13	recipe
1	14	recipe
1	15	recipe
1	16	recipe
1	17	recipe
1	18	recipe
1	19	recipe
1	20	recipe
1	21	recipe
1	22	recipe
\.


--
-- Data for Name: list; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.list (id, name, created_at) FROM stdin;
1	italian	2023-09-22 01:11:09.924298+00
2	food	2023-09-22 01:11:16.521636+00
3	italian	2023-09-23 18:36:51.629121+00
4	italian	2023-09-23 18:36:54.39999+00
\.


--
-- Data for Name: recipe; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.recipe (id, url, created_at, name, ingredients, instructions) FROM stdin;
2	url	2023-09-22 01:09:25.940714+00	pizza	{cheese,sauce}	{"spread sauce","arrange cheese",bake}
11	a	2023-09-23 19:37:17.726257+00	test	{test,test2}	{test,test2}
12	a	2023-09-23 19:38:27.489831+00	test	{test,test2}	{test,test2}
13	a	2023-09-23 19:39:55.030964+00	test	{test,test2}	{test,test2}
14	a	2023-09-23 19:39:55.752284+00	test	{test,test2}	{test,test2}
15	a	2023-09-23 19:39:56.215441+00	test	{test,test2}	{test,test2}
16	a	2023-09-23 19:39:56.631919+00	test	{test,test2}	{test,test2}
17	a	2023-09-23 19:39:57.019646+00	test	{test,test2}	{test,test2}
18	a	2023-09-23 19:39:57.459218+00	test	{test,test2}	{test,test2}
19	a	2023-09-23 19:40:02.633631+00	test	{test,test2}	{test,test2}
20	a	2023-09-23 19:51:04.883531+00	test	{test,test2}	{test,test2}
21		2023-09-23 21:00:40.160643+00	test	{test,test2}	{test,test2}
22		2023-09-23 21:00:40.699893+00	test	{test,test2}	{test,test2}
1	url	2023-09-21 04:36:49.406213+00	spooky	{noodle,sauce}	{"boil water","cook noodle","add sauce"}
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public."user" (id, username, created_at) FROM stdin;
\.


--
-- Name: list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.list_id_seq', 4, true);


--
-- Name: recipe_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.recipe_id_seq', 22, true);


--
-- Name: user_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.user_user_id_seq', 1, false);


--
-- Name: list list_id_pk; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.list
    ADD CONSTRAINT list_id_pk PRIMARY KEY (id);


--
-- Name: recipe recipe_id_pk; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.recipe
    ADD CONSTRAINT recipe_id_pk PRIMARY KEY (id);


--
-- Name: user user_id_pk; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_id_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

