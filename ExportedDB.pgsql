--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3
-- Dumped by pg_dump version 13.3

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
-- Name: notes; Type: TABLE; Schema: public; Owner: jawad_notesapp
--

CREATE TABLE public.notes (
    id integer NOT NULL,
    title character varying(200) NOT NULL,
    content text
);


ALTER TABLE public.notes OWNER TO jawad_notesapp;

--
-- Name: notes_id_seq; Type: SEQUENCE; Schema: public; Owner: jawad_notesapp
--

CREATE SEQUENCE public.notes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notes_id_seq OWNER TO jawad_notesapp;

--
-- Name: notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jawad_notesapp
--

ALTER SEQUENCE public.notes_id_seq OWNED BY public.notes.id;


--
-- Name: notes id; Type: DEFAULT; Schema: public; Owner: jawad_notesapp
--

ALTER TABLE ONLY public.notes ALTER COLUMN id SET DEFAULT nextval('public.notes_id_seq'::regclass);


--
-- Data for Name: notes; Type: TABLE DATA; Schema: public; Owner: jawad_notesapp
--

COPY public.notes (id, title, content) FROM stdin;
1	First Field	This is being populated tojut text out data
2	meow	bili
3	Suyog	Note bieng added to db
4	changed	Lallanwa
5	Will New	Lallanwa
\.


--
-- Name: notes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jawad_notesapp
--

SELECT pg_catalog.setval('public.notes_id_seq', 5, true);


--
-- Name: notes notes_pkey; Type: CONSTRAINT; Schema: public; Owner: jawad_notesapp
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

