--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

-- Started on 2025-02-11 09:28:24

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
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


--
-- TOC entry 4954 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 226 (class 1259 OID 16673)
-- Name: admin; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.admin (
    username text NOT NULL,
    password text NOT NULL
);


--
-- TOC entry 225 (class 1259 OID 16666)
-- Name: auth; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.auth (
    username character varying(255) NOT NULL,
    password character varying(255),
    hit_chance integer,
    status character varying(50),
    name text,
    email text
);


--
-- TOC entry 224 (class 1259 OID 16650)
-- Name: detail_transaksi; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.detail_transaksi (
    id integer NOT NULL,
    id_transaksi integer,
    id_layanan integer,
    quantity integer
);


--
-- TOC entry 223 (class 1259 OID 16649)
-- Name: detail_transaksi_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.detail_transaksi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4955 (class 0 OID 0)
-- Dependencies: 223
-- Name: detail_transaksi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.detail_transaksi_id_seq OWNED BY public.detail_transaksi.id;


--
-- TOC entry 220 (class 1259 OID 16625)
-- Name: layanan; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.layanan (
    id integer NOT NULL,
    nama_layanan character varying(255),
    satuan character varying(10),
    harga integer
);


--
-- TOC entry 219 (class 1259 OID 16624)
-- Name: layanan_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.layanan_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4956 (class 0 OID 0)
-- Dependencies: 219
-- Name: layanan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.layanan_id_seq OWNED BY public.layanan.id;


--
-- TOC entry 218 (class 1259 OID 16618)
-- Name: mst_pegawai; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.mst_pegawai (
    id integer NOT NULL,
    nama_pegawai character varying(255)
);


--
-- TOC entry 217 (class 1259 OID 16617)
-- Name: mst_pegawai_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.mst_pegawai_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4957 (class 0 OID 0)
-- Dependencies: 217
-- Name: mst_pegawai_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.mst_pegawai_id_seq OWNED BY public.mst_pegawai.id;


--
-- TOC entry 216 (class 1259 OID 16611)
-- Name: mst_pelanggan; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.mst_pelanggan (
    id integer NOT NULL,
    nama_pelanggan character varying(255),
    nomor_hp character varying(15)
);


--
-- TOC entry 215 (class 1259 OID 16610)
-- Name: mst_pelanggan_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.mst_pelanggan_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4958 (class 0 OID 0)
-- Dependencies: 215
-- Name: mst_pelanggan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.mst_pelanggan_id_seq OWNED BY public.mst_pelanggan.id;


--
-- TOC entry 222 (class 1259 OID 16632)
-- Name: transaksi; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transaksi (
    id integer NOT NULL,
    id_pelanggan integer,
    id_pegawai integer,
    tanggal_masuk timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    tanggal_keluar date,
    status_pembayaran character varying(50)
);


--
-- TOC entry 221 (class 1259 OID 16631)
-- Name: transaksi_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.transaksi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4959 (class 0 OID 0)
-- Dependencies: 221
-- Name: transaksi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.transaksi_id_seq OWNED BY public.transaksi.id;


--
-- TOC entry 227 (class 1259 OID 16680)
-- Name: waiting_list; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.waiting_list (
    username text NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    status text NOT NULL,
    hit_chance text NOT NULL
);


--
-- TOC entry 4772 (class 2604 OID 16653)
-- Name: detail_transaksi id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.detail_transaksi ALTER COLUMN id SET DEFAULT nextval('public.detail_transaksi_id_seq'::regclass);


--
-- TOC entry 4769 (class 2604 OID 16628)
-- Name: layanan id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.layanan ALTER COLUMN id SET DEFAULT nextval('public.layanan_id_seq'::regclass);


--
-- TOC entry 4768 (class 2604 OID 16621)
-- Name: mst_pegawai id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.mst_pegawai ALTER COLUMN id SET DEFAULT nextval('public.mst_pegawai_id_seq'::regclass);


--
-- TOC entry 4767 (class 2604 OID 16614)
-- Name: mst_pelanggan id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.mst_pelanggan ALTER COLUMN id SET DEFAULT nextval('public.mst_pelanggan_id_seq'::regclass);


--
-- TOC entry 4770 (class 2604 OID 16635)
-- Name: transaksi id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi ALTER COLUMN id SET DEFAULT nextval('public.transaksi_id_seq'::regclass);


--
-- TOC entry 4947 (class 0 OID 16673)
-- Dependencies: 226
-- Data for Name: admin; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.admin VALUES ('yusharwz', '$2y$10$CUyK0BBLF2lDDOyW3x/FM.1SZ0MODB2n0jE8ynZ9tLWdUHwBUr4le');


--
-- TOC entry 4946 (class 0 OID 16666)
-- Dependencies: 225
-- Data for Name: auth; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.auth VALUES ('user1', 'password1', 5000, 'free', NULL, NULL);
INSERT INTO public.auth VALUES ('user2', 'password2', 5000, 'free', NULL, NULL);
INSERT INTO public.auth VALUES ('user3', 'password3', 5000, 'free', NULL, NULL);
INSERT INTO public.auth VALUES ('user4', 'password4', 5000, 'free', NULL, NULL);
INSERT INTO public.auth VALUES ('user5', 'password5', 5000, 'free', NULL, NULL);
INSERT INTO public.auth VALUES ('dindaaf', '$2y$10$p2aTGEnu.FwXhIQQCYxGiuldhtHWIeoCOEmbYtumz.VsxaL8F731i', 998, 'free', 'Dinda Afrilia', 'dindaafrilia26@gmail.com');


--
-- TOC entry 4945 (class 0 OID 16650)
-- Dependencies: 224
-- Data for Name: detail_transaksi; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.detail_transaksi VALUES (1, 1, 1, 2);
INSERT INTO public.detail_transaksi VALUES (2, 1, 3, 1);
INSERT INTO public.detail_transaksi VALUES (3, 1, 4, 3);
INSERT INTO public.detail_transaksi VALUES (4, 2, 2, 1);
INSERT INTO public.detail_transaksi VALUES (5, 2, 3, 2);
INSERT INTO public.detail_transaksi VALUES (6, 2, 4, 1);
INSERT INTO public.detail_transaksi VALUES (7, 3, 1, 3);
INSERT INTO public.detail_transaksi VALUES (8, 3, 2, 2);


--
-- TOC entry 4941 (class 0 OID 16625)
-- Dependencies: 220
-- Data for Name: layanan; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.layanan VALUES (1, 'Cuci', 'KG', 5000);
INSERT INTO public.layanan VALUES (2, 'Cuci dan setrika', 'KG', 7000);
INSERT INTO public.layanan VALUES (3, 'Cuci bedcover', 'Buah', 50000);
INSERT INTO public.layanan VALUES (4, 'Cuci boneka', 'Buah', 25000);


--
-- TOC entry 4939 (class 0 OID 16618)
-- Dependencies: 218
-- Data for Name: mst_pegawai; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.mst_pegawai VALUES (1, 'Diah');
INSERT INTO public.mst_pegawai VALUES (2, 'Calista');
INSERT INTO public.mst_pegawai VALUES (3, 'Kevin');


--
-- TOC entry 4937 (class 0 OID 16611)
-- Dependencies: 216
-- Data for Name: mst_pelanggan; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.mst_pelanggan VALUES (1, 'Yushar', '081234567891');
INSERT INTO public.mst_pelanggan VALUES (2, 'Dinda', '082345678902');
INSERT INTO public.mst_pelanggan VALUES (3, 'Annisa', '083456789013');
INSERT INTO public.mst_pelanggan VALUES (4, 'Rizal', '084567890124');
INSERT INTO public.mst_pelanggan VALUES (5, 'Tegar', '085678901235');
INSERT INTO public.mst_pelanggan VALUES (6, 'Alifia', '086789012346');
INSERT INTO public.mst_pelanggan VALUES (7, 'Jed', '087890123457');
INSERT INTO public.mst_pelanggan VALUES (8, 'Krisna', '088901234568');
INSERT INTO public.mst_pelanggan VALUES (9, 'Clarissa', '089012345679');
INSERT INTO public.mst_pelanggan VALUES (10, 'Farid', '080123456780');


--
-- TOC entry 4943 (class 0 OID 16632)
-- Dependencies: 222
-- Data for Name: transaksi; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.transaksi VALUES (1, 1, 1, '2024-03-25 00:00:00', '2024-03-27', 'Lunas');
INSERT INTO public.transaksi VALUES (2, 2, 2, '2024-03-26 00:00:00', '2024-03-28', 'Belum Lunas');
INSERT INTO public.transaksi VALUES (3, 3, 3, '2024-03-27 00:00:00', '2024-03-29', 'Belum Lunas');


--
-- TOC entry 4948 (class 0 OID 16680)
-- Dependencies: 227
-- Data for Name: waiting_list; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.waiting_list VALUES ('yusharwz', 'Yushar Wahidin Zamzam', 'yusharw@gmail.com', '$2y$10$LSkiqFwKUy0wi7NdY9NA0O575RImbFs2WySBgn92oVDAmRnl.V2um', 'free', '1000');


--
-- TOC entry 4960 (class 0 OID 0)
-- Dependencies: 223
-- Name: detail_transaksi_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.detail_transaksi_id_seq', 8, true);


--
-- TOC entry 4961 (class 0 OID 0)
-- Dependencies: 219
-- Name: layanan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.layanan_id_seq', 4, true);


--
-- TOC entry 4962 (class 0 OID 0)
-- Dependencies: 217
-- Name: mst_pegawai_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.mst_pegawai_id_seq', 3, true);


--
-- TOC entry 4963 (class 0 OID 0)
-- Dependencies: 215
-- Name: mst_pelanggan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.mst_pelanggan_id_seq', 10, true);


--
-- TOC entry 4964 (class 0 OID 0)
-- Dependencies: 221
-- Name: transaksi_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.transaksi_id_seq', 3, true);


--
-- TOC entry 4786 (class 2606 OID 16679)
-- Name: admin admin_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.admin
    ADD CONSTRAINT admin_pkey PRIMARY KEY (username);


--
-- TOC entry 4784 (class 2606 OID 16672)
-- Name: auth auth_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_pkey PRIMARY KEY (username);


--
-- TOC entry 4782 (class 2606 OID 16655)
-- Name: detail_transaksi detail_transaksi_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.detail_transaksi
    ADD CONSTRAINT detail_transaksi_pkey PRIMARY KEY (id);


--
-- TOC entry 4778 (class 2606 OID 16630)
-- Name: layanan layanan_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.layanan
    ADD CONSTRAINT layanan_pkey PRIMARY KEY (id);


--
-- TOC entry 4776 (class 2606 OID 16623)
-- Name: mst_pegawai mst_pegawai_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.mst_pegawai
    ADD CONSTRAINT mst_pegawai_pkey PRIMARY KEY (id);


--
-- TOC entry 4774 (class 2606 OID 16616)
-- Name: mst_pelanggan mst_pelanggan_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.mst_pelanggan
    ADD CONSTRAINT mst_pelanggan_pkey PRIMARY KEY (id);


--
-- TOC entry 4780 (class 2606 OID 16638)
-- Name: transaksi transaksi_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi
    ADD CONSTRAINT transaksi_pkey PRIMARY KEY (id);


--
-- TOC entry 4788 (class 2606 OID 16686)
-- Name: waiting_list waiting_list_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.waiting_list
    ADD CONSTRAINT waiting_list_pkey PRIMARY KEY (username);


--
-- TOC entry 4791 (class 2606 OID 16661)
-- Name: detail_transaksi detail_transaksi_id_layanan_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.detail_transaksi
    ADD CONSTRAINT detail_transaksi_id_layanan_fkey FOREIGN KEY (id_layanan) REFERENCES public.layanan(id);


--
-- TOC entry 4792 (class 2606 OID 16656)
-- Name: detail_transaksi detail_transaksi_id_transaksi_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.detail_transaksi
    ADD CONSTRAINT detail_transaksi_id_transaksi_fkey FOREIGN KEY (id_transaksi) REFERENCES public.transaksi(id);


--
-- TOC entry 4789 (class 2606 OID 16644)
-- Name: transaksi transaksi_id_pegawai_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi
    ADD CONSTRAINT transaksi_id_pegawai_fkey FOREIGN KEY (id_pegawai) REFERENCES public.mst_pegawai(id);


--
-- TOC entry 4790 (class 2606 OID 16639)
-- Name: transaksi transaksi_id_pelanggan_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi
    ADD CONSTRAINT transaksi_id_pelanggan_fkey FOREIGN KEY (id_pelanggan) REFERENCES public.mst_pelanggan(id);


-- Completed on 2025-02-11 09:28:24

--
-- PostgreSQL database dump complete
--

