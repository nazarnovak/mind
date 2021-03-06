--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.5
-- Dumped by pg_dump version 9.5.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: case_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE case_events (
    id integer NOT NULL,
    caseid integer,
    userid integer,
    typeid integer,
    created timestamp with time zone,
    content text
);


ALTER TABLE case_events OWNER TO postgres;

--
-- Name: case_events_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE case_events_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE case_events_id_seq OWNER TO postgres;

--
-- Name: case_events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE case_events_id_seq OWNED BY case_events.id;


--
-- Name: cases; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE cases (
    id integer NOT NULL,
    creatorid integer,
    doctorid integer DEFAULT 0
);


ALTER TABLE cases OWNER TO postgres;

--
-- Name: cases_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE cases_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE cases_id_seq OWNER TO postgres;

--
-- Name: cases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE cases_id_seq OWNED BY cases.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE users (
    id integer NOT NULL,
    name text,
    role integer
);


ALTER TABLE users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY case_events ALTER COLUMN id SET DEFAULT nextval('case_events_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY cases ALTER COLUMN id SET DEFAULT nextval('cases_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: case_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY case_events (id, caseid, userid, typeid, created, content) FROM stdin;
23	14	0	0	2017-02-09 20:16:44+02	Welcome to Min Doktor
24	15	0	0	2017-02-09 20:16:49+02	Welcome to Min Doktor
25	14	111	0	2017-02-09 20:18:07+02	Hello :)
26	14	111	0	2017-02-10 12:13:01+02	/submit something
27	14	111	0	2017-02-10 12:14:10+02	/submit
28	14	111	0	2017-02-10 12:14:31+02	/submit
29	15	111	0	2017-02-10 12:14:39+02	/submit
30	17	0	0	2017-02-10 12:15:06+02	Welcome to Min Doktor
31	17	111	0	2017-02-10 12:17:43+02	/submit
32	17	111	0	2017-02-10 12:20:21+02	/submit
33	14	9	0	2017-02-10 15:29:29+02	doctor here :)
34	17	9	0	2017-02-10 15:29:38+02	hi
35	18	0	0	2017-02-10 15:31:00+02	Welcome to Min Doktor
36	19	0	0	2017-02-11 11:21:20+02	Welcome to Min Doktor
37	14	9	0	2017-02-11 11:21:50+02	how are you
38	14	111	0	2017-02-11 11:26:32+02	fine :D
39	14	111	0	2017-02-11 11:26:37+02	:)
40	14	111	0	2017-02-11 11:26:40+02	@support help me
41	14	9	0	2017-02-11 11:31:34+02	:D
42	17	9	0	2017-02-11 11:31:51+02	how are you?
\.


--
-- Name: case_events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('case_events_id_seq', 42, true);


--
-- Data for Name: cases; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY cases (id, creatorid, doctorid) FROM stdin;
14	111	9
15	111	10
17	111	9
19	111	0
\.


--
-- Name: cases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('cases_id_seq', 19, true);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY users (id, name, role) FROM stdin;
0	MinDoktor	0
111	Nazar Novak	0
8	Patient X	0
9	Doctor X	1
10	Doctor Y	1
11	Doctor Z	1
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('users_id_seq', 11, true);


--
-- Name: case_events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY case_events
    ADD CONSTRAINT case_events_pkey PRIMARY KEY (id);


--
-- Name: cases_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY cases
    ADD CONSTRAINT cases_pkey PRIMARY KEY (id);


--
-- Name: users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

