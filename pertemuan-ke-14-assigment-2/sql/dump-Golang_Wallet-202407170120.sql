--
-- PostgreSQL database dump
--

-- Dumped from database version 12.13
-- Dumped by pg_dump version 12.13

-- Started on 2024-07-17 01:20:16

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

DROP DATABASE "Golang_Wallet";
--
-- TOC entry 2842 (class 1262 OID 22353)
-- Name: Golang_Wallet; Type: DATABASE; Schema: -; Owner: -
--

CREATE DATABASE "Golang_Wallet" WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_United States.1252' LC_CTYPE = 'English_United States.1252';


\connect "Golang_Wallet"

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

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


--
-- TOC entry 2843 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- TOC entry 631 (class 1247 OID 22368)
-- Name: credit_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.credit_type AS ENUM (
    'topup',
    'transfer'
);


--
-- TOC entry 628 (class 1247 OID 22363)
-- Name: transaction_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.transaction_type AS ENUM (
    'debit',
    'credit'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 203 (class 1259 OID 22356)
-- Name: user_saldo; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_saldo (
    id integer NOT NULL,
    user_id integer NOT NULL,
    saldo double precision NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- TOC entry 205 (class 1259 OID 22375)
-- Name: user_saldo_history; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_saldo_history (
    id integer NOT NULL,
    user_id_from integer NOT NULL,
    user_id_to integer NOT NULL,
    type_transaction public.transaction_type NOT NULL,
    type_credit public.credit_type,
    total double precision NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- TOC entry 204 (class 1259 OID 22373)
-- Name: user_saldo_history_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_saldo_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2844 (class 0 OID 0)
-- Dependencies: 204
-- Name: user_saldo_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_saldo_history_id_seq OWNED BY public.user_saldo_history.id;


--
-- TOC entry 202 (class 1259 OID 22354)
-- Name: user_saldo_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_saldo_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2845 (class 0 OID 0)
-- Dependencies: 202
-- Name: user_saldo_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_saldo_id_seq OWNED BY public.user_saldo.id;


--
-- TOC entry 2699 (class 2604 OID 22359)
-- Name: user_saldo id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_saldo ALTER COLUMN id SET DEFAULT nextval('public.user_saldo_id_seq'::regclass);


--
-- TOC entry 2700 (class 2604 OID 22378)
-- Name: user_saldo_history id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_saldo_history ALTER COLUMN id SET DEFAULT nextval('public.user_saldo_history_id_seq'::regclass);


--
-- TOC entry 2834 (class 0 OID 22356)
-- Dependencies: 203
-- Data for Name: user_saldo; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.user_saldo VALUES (3, 2, 40624000, '2024-07-12 00:20:22.913844', '2024-07-12 00:20:22.913844');
INSERT INTO public.user_saldo VALUES (2, 1, 32030000, '2024-07-11 23:30:20.651433', '2024-07-11 23:30:39.3428');
INSERT INTO public.user_saldo VALUES (4, 0, 0, '2024-07-17 00:52:21.175406', '2024-07-17 00:52:21.175406');


--
-- TOC entry 2836 (class 0 OID 22375)
-- Dependencies: 205
-- Data for Name: user_saldo_history; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.user_saldo_history VALUES (1, 1, 1, 'credit', 'topup', 1000000, '2024-07-11 22:45:51.566865', '2024-07-11 22:45:51.566865');
INSERT INTO public.user_saldo_history VALUES (2, 1, 1, 'credit', 'topup', 2000000, '2024-07-11 22:59:39.405059', '2024-07-11 22:59:39.405059');
INSERT INTO public.user_saldo_history VALUES (3, 1, 1, 'credit', 'topup', 2000000, '2024-07-11 23:03:05.03421', '2024-07-11 23:03:05.03421');
INSERT INTO public.user_saldo_history VALUES (4, 1, 1, 'credit', 'topup', 2000000, '2024-07-11 23:25:57.522969', '2024-07-11 23:25:57.522969');
INSERT INTO public.user_saldo_history VALUES (5, 1, 1, 'credit', 'topup', 2000000, '2024-07-11 23:26:23.38899', '2024-07-11 23:26:23.38899');
INSERT INTO public.user_saldo_history VALUES (6, 1, 1, 'credit', 'topup', 2000000, '2024-07-11 23:28:09.634201', '2024-07-11 23:28:09.634201');
INSERT INTO public.user_saldo_history VALUES (7, 1, 1, 'credit', 'topup', 2000000, '2024-07-11 23:30:20.64462', '2024-07-11 23:30:20.64462');
INSERT INTO public.user_saldo_history VALUES (8, 1, 1, 'credit', 'topup', 1000000, '2024-07-11 23:30:39.337695', '2024-07-11 23:30:39.337695');
INSERT INTO public.user_saldo_history VALUES (9, 1, 1, 'credit', 'topup', 1000000, '2024-07-12 00:19:43.888836', '2024-07-12 00:19:43.888836');
INSERT INTO public.user_saldo_history VALUES (10, 1, 1, 'credit', 'topup', 1000000, '2024-07-12 00:20:22.905377', '2024-07-12 00:20:22.905377');
INSERT INTO public.user_saldo_history VALUES (11, 2, 2, 'credit', 'topup', 500000, '2024-07-12 00:20:22.912949', '2024-07-12 00:20:22.912949');
INSERT INTO public.user_saldo_history VALUES (12, 1, 1, 'credit', 'topup', 1000000, '2024-07-12 00:20:28.956633', '2024-07-12 00:20:28.956633');
INSERT INTO public.user_saldo_history VALUES (13, 2, 2, 'credit', 'topup', 500000, '2024-07-12 00:20:28.963078', '2024-07-12 00:20:28.963078');
INSERT INTO public.user_saldo_history VALUES (14, 1, 2, 'credit', 'transfer', 200000, '2024-07-12 00:22:06.74534', '2024-07-12 00:22:06.74534');
INSERT INTO public.user_saldo_history VALUES (15, 1, 2, 'debit', 'transfer', 200000, '2024-07-12 00:22:06.754276', '2024-07-12 00:22:06.754276');
INSERT INTO public.user_saldo_history VALUES (16, 1, 2, 'credit', 'transfer', 200000, '2024-07-12 00:46:10.891831', '2024-07-12 00:46:10.891831');
INSERT INTO public.user_saldo_history VALUES (17, 2, 1, 'debit', 'transfer', 200000, '2024-07-12 00:46:10.900548', '2024-07-12 00:46:10.900548');
INSERT INTO public.user_saldo_history VALUES (18, 1, 2, 'credit', 'transfer', 200000, '2024-07-12 00:46:43.165246', '2024-07-12 00:46:43.165246');
INSERT INTO public.user_saldo_history VALUES (19, 2, 1, 'debit', 'transfer', 200000, '2024-07-12 00:46:43.174625', '2024-07-12 00:46:43.174625');
INSERT INTO public.user_saldo_history VALUES (20, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:47:02.203569', '2024-07-12 00:47:02.203569');
INSERT INTO public.user_saldo_history VALUES (21, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:47:02.209665', '2024-07-12 00:47:02.209665');
INSERT INTO public.user_saldo_history VALUES (22, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:47:24.076537', '2024-07-12 00:47:24.076537');
INSERT INTO public.user_saldo_history VALUES (23, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:47:24.082072', '2024-07-12 00:47:24.082072');
INSERT INTO public.user_saldo_history VALUES (24, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:52:14.953969', '2024-07-12 00:52:14.953969');
INSERT INTO public.user_saldo_history VALUES (25, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:52:14.959841', '2024-07-12 00:52:14.959841');
INSERT INTO public.user_saldo_history VALUES (26, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:52:42.648608', '2024-07-12 00:52:42.648608');
INSERT INTO public.user_saldo_history VALUES (27, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:52:42.655462', '2024-07-12 00:52:42.655462');
INSERT INTO public.user_saldo_history VALUES (28, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:52:51.425234', '2024-07-12 00:52:51.425234');
INSERT INTO public.user_saldo_history VALUES (29, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:52:51.431653', '2024-07-12 00:52:51.431653');
INSERT INTO public.user_saldo_history VALUES (30, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:52:57.940294', '2024-07-12 00:52:57.940294');
INSERT INTO public.user_saldo_history VALUES (31, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:52:57.948069', '2024-07-12 00:52:57.948069');
INSERT INTO public.user_saldo_history VALUES (32, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:53:26.98509', '2024-07-12 00:53:26.98509');
INSERT INTO public.user_saldo_history VALUES (33, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:53:26.991795', '2024-07-12 00:53:26.991795');
INSERT INTO public.user_saldo_history VALUES (34, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:53:43.784501', '2024-07-12 00:53:43.784501');
INSERT INTO public.user_saldo_history VALUES (35, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:53:43.791852', '2024-07-12 00:53:43.791852');
INSERT INTO public.user_saldo_history VALUES (36, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:54:24.386829', '2024-07-12 00:54:24.386829');
INSERT INTO public.user_saldo_history VALUES (37, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:54:24.393973', '2024-07-12 00:54:24.393973');
INSERT INTO public.user_saldo_history VALUES (38, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:54:44.024904', '2024-07-12 00:54:44.024904');
INSERT INTO public.user_saldo_history VALUES (39, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:54:44.030392', '2024-07-12 00:54:44.030392');
INSERT INTO public.user_saldo_history VALUES (40, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:56:28.436873', '2024-07-12 00:56:28.436873');
INSERT INTO public.user_saldo_history VALUES (41, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:56:28.443205', '2024-07-12 00:56:28.443205');
INSERT INTO public.user_saldo_history VALUES (42, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:56:39.61599', '2024-07-12 00:56:39.61599');
INSERT INTO public.user_saldo_history VALUES (43, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:56:39.624098', '2024-07-12 00:56:39.624098');
INSERT INTO public.user_saldo_history VALUES (44, 1, 2, 'credit', 'transfer', 1000000, '2024-07-12 00:57:26.868806', '2024-07-12 00:57:26.868806');
INSERT INTO public.user_saldo_history VALUES (45, 2, 1, 'debit', 'transfer', 1000000, '2024-07-12 00:57:26.875399', '2024-07-12 00:57:26.875399');
INSERT INTO public.user_saldo_history VALUES (46, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 00:58:11.511897', '2024-07-12 00:58:11.511897');
INSERT INTO public.user_saldo_history VALUES (47, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 00:58:11.518091', '2024-07-12 00:58:11.518091');
INSERT INTO public.user_saldo_history VALUES (48, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 00:58:40.33344', '2024-07-12 00:58:40.33344');
INSERT INTO public.user_saldo_history VALUES (49, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 00:58:40.341533', '2024-07-12 00:58:40.341533');
INSERT INTO public.user_saldo_history VALUES (50, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 01:00:59.774226', '2024-07-12 01:00:59.774226');
INSERT INTO public.user_saldo_history VALUES (51, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 01:00:59.779016', '2024-07-12 01:00:59.779016');
INSERT INTO public.user_saldo_history VALUES (52, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 01:03:00.68869', '2024-07-12 01:03:00.68869');
INSERT INTO public.user_saldo_history VALUES (53, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 01:03:00.694642', '2024-07-12 01:03:00.694642');
INSERT INTO public.user_saldo_history VALUES (54, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 01:03:33.852994', '2024-07-12 01:03:33.852994');
INSERT INTO public.user_saldo_history VALUES (55, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 01:03:33.857172', '2024-07-12 01:03:33.857172');
INSERT INTO public.user_saldo_history VALUES (56, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 01:03:41.462099', '2024-07-12 01:03:41.462099');
INSERT INTO public.user_saldo_history VALUES (57, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 01:03:41.470154', '2024-07-12 01:03:41.470154');
INSERT INTO public.user_saldo_history VALUES (58, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 01:08:12.122547', '2024-07-12 01:08:12.122547');
INSERT INTO public.user_saldo_history VALUES (59, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 01:08:12.129296', '2024-07-12 01:08:12.129296');
INSERT INTO public.user_saldo_history VALUES (60, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 01:09:42.007329', '2024-07-12 01:09:42.007329');
INSERT INTO public.user_saldo_history VALUES (61, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 01:09:42.013901', '2024-07-12 01:09:42.013901');
INSERT INTO public.user_saldo_history VALUES (62, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 01:12:52.728434', '2024-07-12 01:12:52.728434');
INSERT INTO public.user_saldo_history VALUES (63, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 01:12:52.735518', '2024-07-12 01:12:52.735518');
INSERT INTO public.user_saldo_history VALUES (64, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 10:38:36.401909', '2024-07-12 10:38:36.401909');
INSERT INTO public.user_saldo_history VALUES (65, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 10:38:36.410773', '2024-07-12 10:38:36.410773');
INSERT INTO public.user_saldo_history VALUES (66, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 10:45:21.03021', '2024-07-12 10:45:21.03021');
INSERT INTO public.user_saldo_history VALUES (67, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 10:45:21.035779', '2024-07-12 10:45:21.035779');
INSERT INTO public.user_saldo_history VALUES (68, 1, 2, 'credit', 'transfer', 2000000, '2024-07-12 11:17:49.715978', '2024-07-12 11:17:49.715978');
INSERT INTO public.user_saldo_history VALUES (69, 2, 1, 'debit', 'transfer', 2000000, '2024-07-12 11:17:49.722202', '2024-07-12 11:17:49.722202');
INSERT INTO public.user_saldo_history VALUES (70, 1, 2, 'credit', 'transfer', 2000000, '2024-07-16 09:49:00.501776', '2024-07-16 09:49:00.501776');
INSERT INTO public.user_saldo_history VALUES (71, 2, 1, 'debit', 'transfer', 2000000, '2024-07-16 09:49:00.514013', '2024-07-16 09:49:00.514013');
INSERT INTO public.user_saldo_history VALUES (72, 1, 1, 'credit', 'topup', 6000, '2024-07-16 21:44:42.97519', '2024-07-16 21:44:42.97519');
INSERT INTO public.user_saldo_history VALUES (73, 1, 1, 'credit', 'topup', 6000, '2024-07-16 21:49:12.031599', '2024-07-16 21:49:12.031599');
INSERT INTO public.user_saldo_history VALUES (74, 1, 1, 'credit', 'topup', 6000, '2024-07-16 21:49:15.532804', '2024-07-16 21:49:15.532804');
INSERT INTO public.user_saldo_history VALUES (75, 1, 1, 'credit', 'topup', 6000, '2024-07-16 21:49:43.586557', '2024-07-16 21:49:43.586557');
INSERT INTO public.user_saldo_history VALUES (76, 1, 1, 'credit', 'topup', 6000, '2024-07-16 21:50:17.266708', '2024-07-16 21:50:17.266708');
INSERT INTO public.user_saldo_history VALUES (77, 1, 1, 'credit', 'topup', 6000, '2024-07-16 21:50:21.658336', '2024-07-16 21:50:21.658336');
INSERT INTO public.user_saldo_history VALUES (78, 1, 1, 'credit', 'topup', 6000, '2024-07-16 21:58:41.760561', '2024-07-16 21:58:41.760561');
INSERT INTO public.user_saldo_history VALUES (79, 1, 1, 'credit', 'topup', 6000, '2024-07-16 23:30:14.471224', '2024-07-16 23:30:14.471224');
INSERT INTO public.user_saldo_history VALUES (80, 1, 2, 'credit', 'transfer', 6000, '2024-07-16 23:30:54.214991', '2024-07-16 23:30:54.214991');
INSERT INTO public.user_saldo_history VALUES (81, 2, 1, 'debit', 'transfer', 6000, '2024-07-16 23:30:54.218318', '2024-07-16 23:30:54.218318');
INSERT INTO public.user_saldo_history VALUES (82, 1, 2, 'credit', 'transfer', 6000, '2024-07-16 23:31:03.682713', '2024-07-16 23:31:03.682713');
INSERT INTO public.user_saldo_history VALUES (83, 2, 1, 'debit', 'transfer', 6000, '2024-07-16 23:31:03.6856', '2024-07-16 23:31:03.6856');
INSERT INTO public.user_saldo_history VALUES (84, 1, 2, 'credit', 'transfer', 6000, '2024-07-16 23:31:16.343149', '2024-07-16 23:31:16.343149');
INSERT INTO public.user_saldo_history VALUES (85, 2, 1, 'debit', 'transfer', 6000, '2024-07-16 23:31:16.346564', '2024-07-16 23:31:16.346564');
INSERT INTO public.user_saldo_history VALUES (86, 0, 0, 'credit', 'topup', 0, '2024-07-17 00:52:21.167019', '2024-07-17 00:52:21.167019');
INSERT INTO public.user_saldo_history VALUES (87, 0, 0, 'credit', 'topup', 0, '2024-07-17 00:56:53.474247', '2024-07-17 00:56:53.474247');
INSERT INTO public.user_saldo_history VALUES (88, 0, 0, 'credit', 'topup', 0, '2024-07-17 00:57:33.69482', '2024-07-17 00:57:33.69482');
INSERT INTO public.user_saldo_history VALUES (89, 0, 0, 'credit', 'topup', 0, '2024-07-17 00:57:37.631853', '2024-07-17 00:57:37.631853');
INSERT INTO public.user_saldo_history VALUES (90, 0, 0, 'credit', 'topup', 0, '2024-07-17 00:57:41.086517', '2024-07-17 00:57:41.086517');
INSERT INTO public.user_saldo_history VALUES (91, 1, 2, 'credit', 'transfer', 6000, '2024-07-17 01:00:43.703475', '2024-07-17 01:00:43.703475');
INSERT INTO public.user_saldo_history VALUES (92, 2, 1, 'debit', 'transfer', 6000, '2024-07-17 01:00:43.706428', '2024-07-17 01:00:43.706428');
INSERT INTO public.user_saldo_history VALUES (93, 1, 1, 'credit', 'topup', 6000, '2024-07-17 01:00:46.542898', '2024-07-17 01:00:46.542898');
INSERT INTO public.user_saldo_history VALUES (94, 0, 0, 'credit', 'topup', 0, '2024-07-17 01:01:52.588626', '2024-07-17 01:01:52.588626');
INSERT INTO public.user_saldo_history VALUES (95, 0, 0, 'credit', 'topup', 0, '2024-07-17 01:01:53.553302', '2024-07-17 01:01:53.553302');
INSERT INTO public.user_saldo_history VALUES (96, 0, 0, 'credit', 'topup', 0, '2024-07-17 01:02:00.813215', '2024-07-17 01:02:00.813215');


--
-- TOC entry 2846 (class 0 OID 0)
-- Dependencies: 204
-- Name: user_saldo_history_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_saldo_history_id_seq', 96, true);


--
-- TOC entry 2847 (class 0 OID 0)
-- Dependencies: 202
-- Name: user_saldo_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_saldo_id_seq', 4, true);


--
-- TOC entry 2706 (class 2606 OID 22382)
-- Name: user_saldo_history user_saldo_history_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_saldo_history
    ADD CONSTRAINT user_saldo_history_pkey PRIMARY KEY (id);


--
-- TOC entry 2704 (class 2606 OID 22361)
-- Name: user_saldo user_saldo_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_saldo
    ADD CONSTRAINT user_saldo_pkey PRIMARY KEY (id);


-- Completed on 2024-07-17 01:20:16

--
-- PostgreSQL database dump complete
--

