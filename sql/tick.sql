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
-- Name: tick; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tick (
    id integer NOT NULL,
    chart_id integer NOT NULL,
    open text NOT NULL,
    close text NOT NULL,
    high text NOT NULL,
    low text NOT NULL,
    day integer NOT NULL,
    volume text NOT NULL,
    "timestamp" bigint NOT NULL
);


ALTER TABLE tick OWNER TO postgres;

--
-- Name: tick_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE tick_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE tick_id_seq OWNER TO postgres;

--
-- Name: tick_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE tick_id_seq OWNED BY tick.id;


--
-- Name: tick id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tick ALTER COLUMN id SET DEFAULT nextval('tick_id_seq'::regclass);


--
-- Name: tick tick_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tick
    ADD CONSTRAINT tick_pkey PRIMARY KEY (id);


--
-- Name: chart_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX chart_id_idx ON tick USING btree (chart_id);


--
-- Name: chart_id_timestamp_uniq_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX chart_id_timestamp_uniq_idx ON tick USING btree (chart_id, "timestamp");


--
-- Name: tick chart_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tick
    ADD CONSTRAINT chart_fk FOREIGN KEY (chart_id) REFERENCES chart(id);


--
-- PostgreSQL database dump complete
--

