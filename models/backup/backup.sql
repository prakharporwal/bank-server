--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4
-- Dumped by pg_dump version 14.4

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: trigger_set_timestamp(); Type: FUNCTION; Schema: public; Owner: admin
--

CREATE FUNCTION public.trigger_set_timestamp() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$;


ALTER FUNCTION public.trigger_set_timestamp() OWNER TO admin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: account_transactions_entries; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.account_transactions_entries (
    transaction_id bigint NOT NULL,
    account_id bigint NOT NULL,
    other_account bigint NOT NULL,
    amount bigint NOT NULL,
    type character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    uid uuid DEFAULT public.uuid_generate_v4() NOT NULL
);


ALTER TABLE public.account_transactions_entries OWNER TO admin;

--
-- Name: COLUMN account_transactions_entries.amount; Type: COMMENT; Schema: public; Owner: admin
--

COMMENT ON COLUMN public.account_transactions_entries.amount IS 'must be positive';


--
-- Name: accounts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.accounts (
    id bigint NOT NULL,
    owner_email character varying NOT NULL,
    balance bigint NOT NULL,
    currency character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.accounts OWNER TO admin;

--
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_id_seq OWNER TO admin;

--
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO admin;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.transactions (
    transaction_id bigint NOT NULL,
    from_account_id bigint NOT NULL,
    to_account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    uid uuid DEFAULT public.uuid_generate_v4() NOT NULL
);


ALTER TABLE public.transactions OWNER TO admin;

--
-- Name: COLUMN transactions.amount; Type: COMMENT; Schema: public; Owner: admin
--

COMMENT ON COLUMN public.transactions.amount IS 'can be negative depending on debit or credit';


--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    user_id bigint NOT NULL,
    username character varying NOT NULL,
    user_email character varying NOT NULL,
    password_hash character varying NOT NULL,
    is_verified boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.users OWNER TO admin;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_user_id_seq OWNER TO admin;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: accounts id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: account_transactions_entries; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.account_transactions_entries (transaction_id, account_id, other_account, amount, type, created_at, uid) FROM stdin;
1662476215826135	3	0	-3000		2022-09-06 14:56:55.826298+00	d94d6ac3-ce97-4bd9-b9e2-2c12776c8625
1662476215826135	2	0	3000		2022-09-06 14:56:55.826298+00	b0df687b-754a-4b9b-a8d5-69af4b39ebbe
1662476238447138	3	0	-3000		2022-09-06 14:57:18.447308+00	a335597c-13d1-4181-9942-a678d835e1c5
1662476238447138	2	0	3000		2022-09-06 14:57:18.447308+00	fe151381-c535-434e-969d-514de9ee4528
1662476243670778	3	0	-3000		2022-09-06 14:57:23.670957+00	54692d56-9d78-47a9-9162-75142d6dfbb0
1662476243670778	2	0	3000		2022-09-06 14:57:23.670957+00	e7525610-d84e-4875-b16c-360e96eed4a8
1662481987757207	1	0	-3000	DEBIT	2022-09-06 16:33:07.758381+00	e0e492ae-53e3-4e04-ae55-d876d61b8719
1662481987757207	3	0	3000	CREDIT	2022-09-06 16:33:07.758381+00	075888c2-20ff-4681-a646-73f0516a084f
\.


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.accounts (id, owner_email, balance, currency, created_at, updated_at) FROM stdin;
2	prakhar@gmail.com	9000	INR	2022-09-06 14:24:54.387932+00	2022-09-06 14:57:23.670957+00
1	prakharporwal@gmail.com	1000	INR	2022-09-06 14:24:42.168492+00	2022-09-06 16:33:07.758381+00
3	porwal@gmail.com	194000	INR	2022-09-06 14:24:59.036725+00	2022-09-06 16:33:07.758381+00
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.transactions (transaction_id, from_account_id, to_account_id, amount, created_at, uid) FROM stdin;
1662476215826135	3	2	3000	2022-09-06 14:56:55.826298+00	d6af91a2-c279-40b6-8889-a6e24fb28058
1662476238447138	3	2	3000	2022-09-06 14:57:18.447308+00	626bb265-effd-47cb-a6ab-9dbe3b53ffe4
1662476243670778	3	2	3000	2022-09-06 14:57:23.670957+00	fc6c4dc3-d1e6-44a7-ae8a-94eab1418e82
1662481987757207	1	3	3000	2022-09-06 16:33:07.758381+00	21164bf3-f5b6-4582-9a49-40cffbd82587
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.users (user_id, username, user_email, password_hash, is_verified, created_at, updated_at) FROM stdin;
\.


--
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.accounts_id_seq', 3, true);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.users_user_id_seq', 1, false);


--
-- Name: account_transactions_entries account_transactions_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.account_transactions_entries
    ADD CONSTRAINT account_transactions_entries_pkey PRIMARY KEY (uid);


--
-- Name: accounts accounts_owner_email_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_owner_email_key UNIQUE (owner_email);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (uid);


--
-- Name: transactions transactions_transaction_id_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_transaction_id_key UNIQUE (transaction_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: account_transactions_entries_account_id_idx; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX account_transactions_entries_account_id_idx ON public.account_transactions_entries USING btree (account_id);


--
-- Name: accounts_owner_email_idx; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX accounts_owner_email_idx ON public.accounts USING btree (owner_email);


--
-- Name: transactions_from_account_id_idx; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX transactions_from_account_id_idx ON public.transactions USING btree (from_account_id);


--
-- Name: transactions_from_account_id_to_account_id_idx; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX transactions_from_account_id_to_account_id_idx ON public.transactions USING btree (from_account_id, to_account_id);


--
-- Name: transactions_to_account_id_idx; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX transactions_to_account_id_idx ON public.transactions USING btree (to_account_id);


--
-- Name: accounts set_timestamp; Type: TRIGGER; Schema: public; Owner: admin
--

CREATE TRIGGER set_timestamp BEFORE UPDATE ON public.accounts FOR EACH ROW EXECUTE FUNCTION public.trigger_set_timestamp();


--
-- Name: users set_timestamp; Type: TRIGGER; Schema: public; Owner: admin
--

CREATE TRIGGER set_timestamp BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.trigger_set_timestamp();


--
-- Name: account_transactions_entries account_transactions_entries_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.account_transactions_entries
    ADD CONSTRAINT account_transactions_entries_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.accounts(id);


--
-- Name: transactions transactions_from_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_from_account_id_fkey FOREIGN KEY (from_account_id) REFERENCES public.accounts(id);


--
-- Name: transactions transactions_to_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_to_account_id_fkey FOREIGN KEY (to_account_id) REFERENCES public.accounts(id);


--
-- PostgreSQL database dump complete
--

