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
-- Name: chart; Type: TABLE; Schema: public; Owner: payam
--

CREATE TABLE chart (
    id integer NOT NULL,
    market_id integer NOT NULL,
    "interval" text NOT NULL
);


ALTER TABLE chart OWNER TO payam;

--
-- Name: chart_id_seq; Type: SEQUENCE; Schema: public; Owner: payam
--

CREATE SEQUENCE chart_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE chart_id_seq OWNER TO payam;

--
-- Name: chart_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: payam
--

ALTER SEQUENCE chart_id_seq OWNED BY chart.id;


--
-- Name: chart id; Type: DEFAULT; Schema: public; Owner: payam
--

ALTER TABLE ONLY chart ALTER COLUMN id SET DEFAULT nextval('chart_id_seq'::regclass);


--
-- Name: chart chart_pkey; Type: CONSTRAINT; Schema: public; Owner: payam
--

ALTER TABLE ONLY chart
    ADD CONSTRAINT chart_pkey PRIMARY KEY (id);


--
-- Name: interval_idx; Type: INDEX; Schema: public; Owner: payam
--

CREATE INDEX interval_idx ON chart USING btree ("interval");


--
-- Name: market_id_idx; Type: INDEX; Schema: public; Owner: payam
--

CREATE INDEX market_id_idx ON chart USING btree (market_id);


--
-- Name: market_interval_uniq_idx; Type: INDEX; Schema: public; Owner: payam
--

CREATE UNIQUE INDEX market_interval_uniq_idx ON chart USING btree (market_id, "interval");


--
-- Name: chart fk_market_id; Type: FK CONSTRAINT; Schema: public; Owner: payam
--

ALTER TABLE ONLY chart
    ADD CONSTRAINT fk_market_id FOREIGN KEY (market_id) REFERENCES market(id);


--
-- PostgreSQL database dump complete
--

