--
-- PostgreSQL database dump
--

\restrict f6lVthRo5M9nt5Zqia6ESiTk6vYP3Aa6Npet7Q8nedNhAd4fzTOA5XlQtqyylik

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.8 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: book_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.book_status AS ENUM (
    'available',
    'borrowed',
    'lost'
);


ALTER TYPE public.book_status OWNER TO postgres;

--
-- Name: fine_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.fine_status AS ENUM (
    'unpaid',
    'paid'
);


ALTER TYPE public.fine_status OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    id bigint NOT NULL,
    title character varying(150) NOT NULL,
    author character varying(100) NOT NULL,
    isbn character varying(13) NOT NULL,
    status character varying(255) DEFAULT 'available'::character varying NOT NULL,
    created_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    created_by bigint,
    updated_by bigint,
    deleted_by bigint,
    CONSTRAINT books_status_check CHECK (((status)::text = ANY (ARRAY[('available'::character varying)::text, ('borrowed'::character varying)::text, ('lost'::character varying)::text])))
);


ALTER TABLE public.books OWNER TO postgres;

--
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.books_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.books_id_seq OWNER TO postgres;

--
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;


--
-- Name: cache; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cache (
    key character varying(255) NOT NULL,
    value text NOT NULL,
    expiration integer NOT NULL
);


ALTER TABLE public.cache OWNER TO postgres;

--
-- Name: cache_locks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cache_locks (
    key character varying(255) NOT NULL,
    owner character varying(255) NOT NULL,
    expiration integer NOT NULL
);


ALTER TABLE public.cache_locks OWNER TO postgres;

--
-- Name: countries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.countries (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    iso_code character(2) NOT NULL,
    phone_code smallint,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    created_by bigint,
    updated_by bigint,
    deleted_by bigint
);


ALTER TABLE public.countries OWNER TO postgres;

--
-- Name: countries_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.countries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.countries_id_seq OWNER TO postgres;

--
-- Name: countries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.countries_id_seq OWNED BY public.countries.id;


--
-- Name: failed_jobs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.failed_jobs (
    id bigint NOT NULL,
    uuid character varying(255) NOT NULL,
    connection text NOT NULL,
    queue text NOT NULL,
    payload text NOT NULL,
    exception text NOT NULL,
    failed_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.failed_jobs OWNER TO postgres;

--
-- Name: failed_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.failed_jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.failed_jobs_id_seq OWNER TO postgres;

--
-- Name: failed_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.failed_jobs_id_seq OWNED BY public.failed_jobs.id;


--
-- Name: fines; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fines (
    id bigint NOT NULL,
    loan_id bigint NOT NULL,
    amount numeric(10,2) NOT NULL,
    reason character varying(255) NOT NULL,
    status character varying(255) DEFAULT 'unpaid'::character varying NOT NULL,
    created_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    created_by bigint,
    updated_by bigint,
    deleted_by bigint,
    CONSTRAINT fines_status_check CHECK (((status)::text = ANY (ARRAY[('unpaid'::character varying)::text, ('paid'::character varying)::text])))
);


ALTER TABLE public.fines OWNER TO postgres;

--
-- Name: fines_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fines_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.fines_id_seq OWNER TO postgres;

--
-- Name: fines_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fines_id_seq OWNED BY public.fines.id;


--
-- Name: job_batches; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.job_batches (
    id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    total_jobs integer NOT NULL,
    pending_jobs integer NOT NULL,
    failed_jobs integer NOT NULL,
    failed_job_ids text NOT NULL,
    options text,
    cancelled_at integer,
    created_at integer NOT NULL,
    finished_at integer
);


ALTER TABLE public.job_batches OWNER TO postgres;

--
-- Name: jobs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.jobs (
    id bigint NOT NULL,
    queue character varying(255) NOT NULL,
    payload text NOT NULL,
    attempts smallint NOT NULL,
    reserved_at integer,
    available_at integer NOT NULL,
    created_at integer NOT NULL
);


ALTER TABLE public.jobs OWNER TO postgres;

--
-- Name: jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.jobs_id_seq OWNER TO postgres;

--
-- Name: jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.jobs_id_seq OWNED BY public.jobs.id;


--
-- Name: members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.members (
    id bigint NOT NULL,
    full_name character varying(150) NOT NULL,
    phone character varying(15) NOT NULL,
    email character varying(150) NOT NULL,
    membership_no character varying(20) NOT NULL,
    created_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    created_by bigint,
    updated_by bigint,
    deleted_by bigint
);


ALTER TABLE public.members OWNER TO postgres;

--
-- Name: members_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.members_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.members_id_seq OWNER TO postgres;

--
-- Name: members_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.members_id_seq OWNED BY public.members.id;


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.migrations (
    id integer NOT NULL,
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);


ALTER TABLE public.migrations OWNER TO postgres;

--
-- Name: migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.migrations_id_seq OWNER TO postgres;

--
-- Name: migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.migrations_id_seq OWNED BY public.migrations.id;


--
-- Name: password_reset_tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.password_reset_tokens (
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp(0) without time zone
);


ALTER TABLE public.password_reset_tokens OWNER TO postgres;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(255),
    created_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    created_by bigint,
    updated_by bigint,
    deleted_by bigint
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.roles_id_seq OWNER TO postgres;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id character varying(255) NOT NULL,
    user_id bigint,
    ip_address character varying(45),
    user_agent text,
    payload text NOT NULL,
    last_activity integer NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- Name: staff; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.staff (
    id bigint NOT NULL,
    full_name character varying(150) NOT NULL,
    email character varying(150) NOT NULL,
    phone character varying(15),
    username character varying(50) NOT NULL,
    password_hash character varying(255) NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    created_by bigint,
    updated_by bigint,
    deleted_by bigint
);


ALTER TABLE public.staff OWNER TO postgres;

--
-- Name: staff_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.staff_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.staff_id_seq OWNER TO postgres;

--
-- Name: staff_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.staff_id_seq OWNED BY public.staff.id;


--
-- Name: universities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universities (
    id bigint NOT NULL,
    name character varying(150) NOT NULL,
    abbreviation character varying(20) NOT NULL,
    email character varying(150),
    website character varying(255),
    established_year smallint,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    created_by bigint,
    updated_by bigint,
    deleted_by bigint
);


ALTER TABLE public.universities OWNER TO postgres;

--
-- Name: universities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.universities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.universities_id_seq OWNER TO postgres;

--
-- Name: universities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.universities_id_seq OWNED BY public.universities.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    email_verified_at timestamp(0) without time zone,
    password character varying(255) NOT NULL,
    remember_token character varying(100),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: books id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


--
-- Name: countries id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.countries ALTER COLUMN id SET DEFAULT nextval('public.countries_id_seq'::regclass);


--
-- Name: failed_jobs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.failed_jobs ALTER COLUMN id SET DEFAULT nextval('public.failed_jobs_id_seq'::regclass);


--
-- Name: fines id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fines ALTER COLUMN id SET DEFAULT nextval('public.fines_id_seq'::regclass);


--
-- Name: jobs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jobs ALTER COLUMN id SET DEFAULT nextval('public.jobs_id_seq'::regclass);


--
-- Name: members id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members ALTER COLUMN id SET DEFAULT nextval('public.members_id_seq'::regclass);


--
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id SET DEFAULT nextval('public.migrations_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: staff id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.staff ALTER COLUMN id SET DEFAULT nextval('public.staff_id_seq'::regclass);


--
-- Name: universities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universities ALTER COLUMN id SET DEFAULT nextval('public.universities_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.books (id, title, author, isbn, status, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by) FROM stdin;
1	Introduction to Computer Science	John Smith	9780123456789	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
2	Advanced Mathematics	Mary Johnson	9780123456790	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
3	Physics for Engineers	David Wilson	9780123456806	borrowed	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
4	Chemistry Fundamentals	Sarah Brown	9780123456813	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
5	Biology: Life Sciences	Michael Davis	9780123456820	lost	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
6	Economics Principles	Emily Miller	9780123456837	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
7	Business Management	James Garcia	9780123456844	borrowed	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
8	Accounting Basics	Jennifer Martinez	9780123456851	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
9	Marketing Strategies	Robert Anderson	9780123456868	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
10	Human Resource Management	Linda Taylor	9780123456875	borrowed	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
11	Software Engineering	William Thomas	9780123456882	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
12	Data Structures	Patricia Jackson	9780123456899	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
13	Artificial Intelligence	Christopher White	9780123456905	borrowed	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
14	Machine Learning	Nancy Harris	9780123456912	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
15	Web Development	Daniel Martin	9780123456929	lost	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
16	Database Systems	Karen Thompson	9780123456936	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
17	Network Security	Paul Garcia	9780123456943	borrowed	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
18	Cloud Computing	Susan Martinez	9780123456950	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
19	Mobile App Development	Joseph Robinson	9780123456967	available	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
20	DevOps Practices	Lisa Clark	9780123456974	borrowed	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
\.


--
-- Data for Name: cache; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cache (key, value, expiration) FROM stdin;
\.


--
-- Data for Name: cache_locks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cache_locks (key, owner, expiration) FROM stdin;
\.


--
-- Data for Name: countries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.countries (id, name, iso_code, phone_code, is_active, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by) FROM stdin;
1	Tanzania	TZ	255	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
2	Kenya	KE	254	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
3	Uganda	UG	256	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
4	Rwanda	RW	250	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
5	Burundi	BI	257	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
6	South Africa	ZA	27	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
7	Nigeria	NG	234	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
8	Ghana	GH	233	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
9	Egypt	EG	20	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
10	Morocco	MA	212	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
11	Ethiopia	ET	251	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
12	Sudan	SD	249	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
13	Libya	LY	218	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
14	Algeria	DZ	213	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
15	Tunisia	TN	216	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
16	Cameroon	CM	237	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
17	Ivory Coast	CI	225	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
18	Senegal	SN	221	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
19	Mali	ML	223	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
20	Burkina Faso	BF	226	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
\.


--
-- Data for Name: failed_jobs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.failed_jobs (id, uuid, connection, queue, payload, exception, failed_at) FROM stdin;
\.


--
-- Data for Name: fines; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fines (id, loan_id, amount, reason, status, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by) FROM stdin;
1	1	50.00	Late return - overdue by 15 days	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
2	2	25.00	Book damage - torn pages	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
3	3	100.00	Lost book - replacement cost	paid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
4	4	75.00	Late return - overdue by 30 days	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
5	5	35.00	Book damage - water damage	paid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
6	6	60.00	Late return - overdue by 20 days	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
7	7	150.00	Lost book - rare edition	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
8	8	40.00	Book damage - broken spine	paid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
9	9	80.00	Late return - overdue by 25 days	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
10	10	45.00	Book damage - missing cover	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
11	11	90.00	Lost book - textbook	paid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
12	12	55.00	Late return - overdue by 18 days	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
13	13	30.00	Book damage - pen marks	paid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
14	14	120.00	Lost book - reference book	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
15	15	65.00	Late return - overdue by 22 days	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
16	16	85.00	Book damage - torn cover	paid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
17	17	95.00	Late return - overdue by 28 days	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
18	18	70.00	Lost book - journal	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
19	19	110.00	Book damage - water stains	paid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
20	20	125.00	Lost book - research material	unpaid	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
\.


--
-- Data for Name: job_batches; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.job_batches (id, name, total_jobs, pending_jobs, failed_jobs, failed_job_ids, options, cancelled_at, created_at, finished_at) FROM stdin;
\.


--
-- Data for Name: jobs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.jobs (id, queue, payload, attempts, reserved_at, available_at, created_at) FROM stdin;
\.


--
-- Data for Name: members; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.members (id, full_name, phone, email, membership_no, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by) FROM stdin;
1	Ali Hassan	+255712345001	ali.hassan@email.com	LIB001	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
2	Fatuma Mohamed	+255712345002	fatuma.mohamed@email.com	LIB002	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
3	John Mwangi	+254712345003	john.mwangi@email.com	LIB003	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
4	Grace Nakato	+256712345004	grace.nakato@email.com	LIB004	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
5	Jean Pierre	+250712345005	jean.pierre@email.com	LIB005	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
6	Sarah Niyonzima	+257712345006	sarah.niyonzima@email.com	LIB006	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
7	Michael Smith	+27712345007	michael.smith@email.com	LIB007	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
8	Amina Abubakar	+234712345008	amina.abubakar@email.com	LIB008	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
9	Kwame Asante	+233712345009	kwame.asante@email.com	LIB009	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
10	Ahmed Khalil	+20712345010	ahmed.khalil@email.com	LIB010	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
11	Rachida Benani	+212712345011	rachida.benani@email.com	LIB011	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
12	Bekele Tadesse	+251712345012	bekele.tadesse@email.com	LIB012	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
13	Omar Hassan	+249712345013	omar.hassan@email.com	LIB013	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
14	Mariam Tripoli	+218712345014	mariam.tripoli@email.com	LIB014	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
15	Karim Algiers	+213712345015	karim.algiers@email.com	LIB015	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
16	Sonia Tunis	+216712345016	sonia.tunis@email.com	LIB016	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
17	Pierre Douala	+237712345017	pierre.douala@email.com	LIB017	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
18	Marie Abidjan	+225712345018	marie.abidjan@email.com	LIB018	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
19	Ibrahim Dakar	+221712345019	ibrahim.dakar@email.com	LIB019	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
20	Awa Bamako	+223712345020	awa.bamako@email.com	LIB020	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
\.


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.migrations (id, migration, batch) FROM stdin;
1	0001_01_01_000000_create_users_table	1
2	0001_01_01_000001_create_cache_table	1
3	0001_01_01_000002_create_jobs_table	1
4	2026_02_12_120601_create_books_table	1
5	2026_02_12_121050_create_members_table	1
6	2026_02_12_121246_create_fines_table	1
7	2026_02_12_121344_create_staff_table	1
8	2026_02_12_121452_create_roles_table	1
9	2026_02_12_121535_create_universities_table	1
10	2026_02_12_121712_create_countries_table	1
\.


--
-- Data for Name: password_reset_tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.password_reset_tokens (email, token, created_at) FROM stdin;
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, name, description, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by) FROM stdin;
1	Super Admin	System administrator with full access	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
2	Admin	Library administrator with management access	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
3	Librarian	Library staff with book management access	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
4	Accountant	Financial staff with fine management access	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
5	Senior Librarian	Experienced librarian with supervisory access	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
6	Library Assistant	Junior library staff with limited access	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
7	Circulation Manager	Manages book loans and returns	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
8	Acquisition Manager	Manages book acquisition and procurement	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
9	Catalog Manager	Manages book cataloging and classification	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
10	Reference Librarian	Provides reference and research assistance	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
11	Systems Librarian	Manages library systems and technology	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
12	Archive Manager	Manages library archives and special collections	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
13	Digital Resources Manager	Manages digital library resources	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
14	Interlibrary Loan Manager	Manages interlibrary loan services	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
15	Collection Development Manager	Manages collection development and weeding	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
16	Library Director	Head of library with strategic oversight	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
17	Branch Manager	Manages specific library branch	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
18	Technical Services Manager	Manages technical library services	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
19	Public Services Manager	Manages public library services	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
20	Research Support Manager	Supports research and academic services	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, user_id, ip_address, user_agent, payload, last_activity) FROM stdin;
\.


--
-- Data for Name: staff; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.staff (id, full_name, email, phone, username, password_hash, is_active, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by) FROM stdin;
1	John Smith	john.smith@library.com	+255712345678	jsmith	$2y$12$X9Si7n2w7oBcTPkNrhioA.p2/50xs5LtCFn1NN8.oRANaxp88CLBe	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
2	Mary Johnson	mary.johnson@library.com	+255712345679	mjohnson	$2y$12$1scY2CgiyMzOSiaH5uAVrePXmdPeB0jpfdi062aE4tkiRPjaIsT.C	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
3	David Wilson	david.wilson@library.com	+255712345680	dwilson	$2y$12$lD5D4kLBJ0Ss8.meTLllUupLylnGsmi.4jaGe4HOrWUf2LQIWbDjC	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
4	Sarah Brown	sarah.brown@library.com	+255712345681	sbrown	$2y$12$kMDVFZKYG.QZIMZW5OuA1uKk84STwP.VZJLmBBr4Crb65.lcaQk2W	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
5	Michael Davis	michael.davis@library.com	+255712345682	mdavis	$2y$12$0GvD5UPplEVsyRgsTTanTO2HCggAZDNIdEz6eBl42HMjNEisCSVny	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
6	Emily Miller	emily.miller@library.com	+255712345683	emiller	$2y$12$XvS1JfpDfllAfT2uLt4dreJWJDmhsJJHrEEwDglDzAbVzzuRSNDC.	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
7	James Garcia	james.garcia@library.com	+255712345684	jgarcia	$2y$12$BUCb1lzp/u1R5WNSEA8SaOJzqnrPPC15xFtgFQZVRo2RuSbdGKBHm	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
8	Jennifer Martinez	jennifer.martinez@library.com	+255712345685	jmartinez	$2y$12$wUQwEMTjvBHUDxX.EafOKeVmaFpuLBGkVe5OK1pdfo8Xww6XpGvMm	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
9	Robert Anderson	robert.anderson@library.com	+255712345686	randerson	$2y$12$mqurRSmWpK/nTn9hVpc06unJRvER61j5CprIrckoKZIr3WC/HZHt6	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
10	Linda Taylor	linda.taylor@library.com	+255712345687	ltaylor	$2y$12$1M24yiuyQI.9hxhw9HermeN7V02pURBChJD9gzLkn1R5TNiBtEFQi	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
11	William Thomas	william.thomas@library.com	+255712345688	wthomas	$2y$12$FL.DaxkADN0Ee.zLvcnIDuRRqt.sjcC92uVZL7IwnoZbUi/Q9pCE2	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
12	Patricia Jackson	patricia.jackson@library.com	+255712345689	pjackson	$2y$12$Uaz5Vqvu0NjYAR3rklAvrupbEu/Tz10IL8afpWRtMzfUNNcySZNP6	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
13	Christopher White	christopher.white@library.com	+255712345690	cwhite	$2y$12$rXivBGsrWMz.zUEOldpd1u/YpFIW9Ostc7IvRZyi3loVu8d9oOOKu	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
14	Nancy Harris	nancy.harris@library.com	+255712345691	nharris	$2y$12$EBAgsW6KknBPK3QmMLuFiuPtsbtaq9SGB7VhFaHHyv1PMJM3VIQCO	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
15	Daniel Martin	daniel.martin@library.com	+255712345692	dmartin	$2y$12$UpHZobisExqApAqPOwm/tudDt9piDJzRK5EQSHa2XY52k203Qq7XC	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
16	Karen Thompson	karen.thompson@library.com	+255712345693	kthompson	$2y$12$VutCJAJpsGh0EuI1xnDJOOOv/Ic/3LyXas9bcsSAO5woZ2NLpXZey	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
17	Paul Garcia	paul.garcia@library.com	+255712345694	pgarcia	$2y$12$1Csp3O8OKCUNjCF7x9idHu2xgP4d4LH9rZ07EsXKVHuzagLf3fpP2	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
18	Susan Martinez	susan.martinez@library.com	+255712345695	smartinez	$2y$12$4oNmRKpZ.REVTbWkDW0v8eX.aprXrL4TH65sBwAB7rNt4hvPsjE52	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
19	Joseph Robinson	joseph.robinson@library.com	+255712345696	jrobinson	$2y$12$OvfumMuwZEC/V/xaOGfepuBEGObNg0YTyBJ8IolzF5k.L2hrGqcr2	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
20	Lisa Clark	lisa.clark@library.com	+255712345697	lclark	$2y$12$WzvO4aW5ZM.vPaDAPAVCfetuiULdRu9N1T9DzVRnnX8WvTPu/jfA2	t	2026-02-12 19:25:35	2026-02-12 19:25:35	\N	\N	\N	\N
\.


--
-- Data for Name: universities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universities (id, name, abbreviation, email, website, established_year, is_active, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by) FROM stdin;
1	University of Dar es Salaam	UDSM	info@udsm.ac.tz	www.udsm.ac.tz	1961	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
2	Muhimbili University of Health and Allied Sciences	MUHAS	info@muhimbili.ac.tz	www.muhimbili.ac.tz	1991	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
3	Sokoine University of Agriculture	SUA	info@sua.ac.tz	www.sua.ac.tz	1984	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
4	University of Dodoma	UDOM	info@udom.ac.tz	www.udom.ac.tz	2007	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
5	Mzumbe University	MU	info@mzumbe.ac.tz	www.mzumbe.ac.tz	2001	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
6	Open University of Tanzania	OUT	info@out.ac.tz	www.out.ac.tz	1992	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
7	Catholic University of Health and Allied Sciences	CUHAS	info@cuhas.ac.tz	www.cuhas.ac.tz	2006	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
8	Ardhi University	ARU	info@aru.ac.tz	www.aru.ac.tz	2007	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
9	Nelson Mandela African Institution of Science and Technology	NM-AIST	info@nm-aist.ac.tz	www.nm-aist.ac.tz	2011	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
10	St. Augustine University of Tanzania	SAUT	info@saut.ac.tz	www.saut.ac.tz	1998	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
11	Moshi Co-operative University	MoCU	info@mocu.ac.tz	www.mocu.ac.tz	2005	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
12	Hubert Kairuki Memorial University	HKMU	info@hkmu.ac.tz	www.hkmu.ac.tz	1997	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
13	Mbeya University of Science and Technology	MUST	info@must.ac.tz	www.must.ac.tz	2012	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
14	International Medical and Technological University	IMTU	info@imtu.ac.tz	www.imtu.ac.tz	1995	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
15	University of Iringa	UI	info@iringa.ac.tz	www.iringa.ac.tz	2010	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
16	St John's University of Tanzania	SJUT	info@sjut.ac.tz	www.sjut.ac.tz	2006	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
17	Tumaini University Makumira	TUMA	info@tumaini.ac.tz	www.tumaini.ac.tz	2004	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
18	Zanzibar University	ZU	info@zanzibaruniversity.ac.tz	www.zanzibaruniversity.ac.tz	2002	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
19	Muslim University of Morogoro	MUM	info@mum.ac.tz	www.mum.ac.tz	2004	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
20	The University of Arusha	UOA	info@uoa.ac.tz	www.uoa.ac.tz	2005	t	2026-02-12 19:25:31	2026-02-12 19:25:31	\N	\N	\N	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, email_verified_at, password, remember_token, created_at, updated_at) FROM stdin;
\.


--
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.books_id_seq', 20, true);


--
-- Name: countries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.countries_id_seq', 20, true);


--
-- Name: failed_jobs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.failed_jobs_id_seq', 1, false);


--
-- Name: fines_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fines_id_seq', 20, true);


--
-- Name: jobs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.jobs_id_seq', 1, false);


--
-- Name: members_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.members_id_seq', 20, true);


--
-- Name: migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.migrations_id_seq', 10, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_seq', 20, true);


--
-- Name: staff_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.staff_id_seq', 20, true);


--
-- Name: universities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.universities_id_seq', 20, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- Name: cache_locks cache_locks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cache_locks
    ADD CONSTRAINT cache_locks_pkey PRIMARY KEY (key);


--
-- Name: cache cache_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cache
    ADD CONSTRAINT cache_pkey PRIMARY KEY (key);


--
-- Name: countries countries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT countries_pkey PRIMARY KEY (id);


--
-- Name: failed_jobs failed_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_pkey PRIMARY KEY (id);


--
-- Name: failed_jobs failed_jobs_uuid_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_uuid_unique UNIQUE (uuid);


--
-- Name: fines fines_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fines
    ADD CONSTRAINT fines_pkey PRIMARY KEY (id);


--
-- Name: job_batches job_batches_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.job_batches
    ADD CONSTRAINT job_batches_pkey PRIMARY KEY (id);


--
-- Name: jobs jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jobs
    ADD CONSTRAINT jobs_pkey PRIMARY KEY (id);


--
-- Name: members members_membership_no_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT members_membership_no_unique UNIQUE (membership_no);


--
-- Name: members members_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT members_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- Name: password_reset_tokens password_reset_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.password_reset_tokens
    ADD CONSTRAINT password_reset_tokens_pkey PRIMARY KEY (email);


--
-- Name: roles roles_name_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_unique UNIQUE (name);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: staff staff_email_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.staff
    ADD CONSTRAINT staff_email_unique UNIQUE (email);


--
-- Name: staff staff_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.staff
    ADD CONSTRAINT staff_pkey PRIMARY KEY (id);


--
-- Name: staff staff_username_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.staff
    ADD CONSTRAINT staff_username_unique UNIQUE (username);


--
-- Name: universities universities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universities
    ADD CONSTRAINT universities_pkey PRIMARY KEY (id);


--
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_unique UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: cache_expiration_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX cache_expiration_index ON public.cache USING btree (expiration);


--
-- Name: cache_locks_expiration_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX cache_locks_expiration_index ON public.cache_locks USING btree (expiration);


--
-- Name: jobs_queue_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX jobs_queue_index ON public.jobs USING btree (queue);


--
-- Name: sessions_last_activity_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX sessions_last_activity_index ON public.sessions USING btree (last_activity);


--
-- Name: sessions_user_id_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX sessions_user_id_index ON public.sessions USING btree (user_id);


--
-- PostgreSQL database dump complete
--

\unrestrict f6lVthRo5M9nt5Zqia6ESiTk6vYP3Aa6Npet7Q8nedNhAd4fzTOA5XlQtqyylik

