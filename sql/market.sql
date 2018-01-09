--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.4
-- Dumped by pg_dump version 9.6.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: market; Type: TABLE; Schema: public; Owner: payam
--

CREATE TABLE market (
    id integer NOT NULL,
    exchange_name text NOT NULL,
    base_currency text NOT NULL,
    base_currency_name text,
    market_currency text NOT NULL,
    market_currency_name text,
    market_key text NOT NULL
);


ALTER TABLE market OWNER TO payam;

--
-- Name: market_id_seq; Type: SEQUENCE; Schema: public; Owner: payam
--

CREATE SEQUENCE market_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE market_id_seq OWNER TO payam;

--
-- Name: market_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: payam
--

ALTER SEQUENCE market_id_seq OWNED BY market.id;


--
-- Name: market id; Type: DEFAULT; Schema: public; Owner: payam
--

ALTER TABLE ONLY market ALTER COLUMN id SET DEFAULT nextval('market_id_seq'::regclass);


--
-- Name: market market_pkey; Type: CONSTRAINT; Schema: public; Owner: payam
--

ALTER TABLE ONLY market
    ADD CONSTRAINT market_pkey PRIMARY KEY (id);


--
-- Name: base_currency_idx; Type: INDEX; Schema: public; Owner: payam
--

CREATE INDEX base_currency_idx ON market USING btree (base_currency);


--
-- Name: exchange_market_key_uniq_idx; Type: INDEX; Schema: public; Owner: payam
--

CREATE UNIQUE INDEX exchange_market_key_uniq_idx ON market USING btree (exchange_name, market_key);


--
-- Name: exchange_name; Type: INDEX; Schema: public; Owner: payam
--

CREATE INDEX exchange_name ON market USING btree (exchange_name);


--
-- Name: market_currency_idx; Type: INDEX; Schema: public; Owner: payam
--

CREATE INDEX market_currency_idx ON market USING btree (market_currency);


--
-- Name: market_key_idx; Type: INDEX; Schema: public; Owner: payam
--

CREATE INDEX market_key_idx ON market USING btree (market_key);


--
-- PostgreSQL database dump complete
--

