-- +migrate Up notransaction

-- Create Market ID Sequence
CREATE SEQUENCE IF NOT EXISTS "market_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Create Table
CREATE TABLE IF NOT EXISTS "market" (
    "id" integer NOT NULL DEFAULT nextval('market_id_seq'::regclass),
    "exchange_name" text NOT NULL,
    "base_currency" text NOT NULL,
    "base_currency_name" text,
    "market_currency" text NOT NULL,
    "market_currency_name" text,
    "market_key" text NOT NULL
);

-- Primary Keys
ALTER TABLE "market" DROP CONSTRAINT IF EXISTS "market_pkey";
ALTER TABLE "market" ADD CONSTRAINT "market_pkey" PRIMARY KEY ("id");

-- Indexes
CREATE INDEX IF NOT EXISTS "base_currency_idx" ON "market" USING btree ("base_currency");
CREATE UNIQUE INDEX IF NOT EXISTS "exchange_market_key_uniq_idx" ON "market" USING btree ("exchange_name", "market_key");
CREATE INDEX IF NOT EXISTS "exchange_name" ON "market" USING btree ("exchange_name");
CREATE INDEX IF NOT EXISTS "market_currency_idx" ON "market" USING btree ("market_currency");
CREATE INDEX IF NOT EXISTS "market_key_idx" ON "market" USING btree ("market_key");


-- Exchange ID Sequence
CREATE SEQUENCE IF NOT EXISTS "exchange_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "exchange" (
    "id" integer NOT NULL DEFAULT nextval('exchange_id_seq'::regclass),
    "name" text NOT NULL
);

ALTER TABLE "exchange" DROP CONSTRAINT IF EXISTS "exchange_pkey";
ALTER TABLE "exchange" ADD CONSTRAINT "exchange_pkey" PRIMARY KEY ("id");


-- Charts

CREATE SEQUENCE IF NOT EXISTS "chart_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "chart" (
    "id" integer NOT NULL DEFAULT nextval('chart_id_seq'::regclass),
    "market_id" integer NOT NULL,
    "interval" text NOT NULL
);


-- Primary Key Chart
ALTER TABLE  "chart" DROP CONSTRAINT IF EXISTS "chart_pkey";
ALTER TABLE  "chart" ADD CONSTRAINT "chart_pkey" PRIMARY KEY ("id");

-- Chart Indexes
CREATE INDEX IF NOT EXISTS "interval_idx" ON "chart" USING btree ("interval");
CREATE INDEX IF NOT EXISTS "market_id_idx" ON "chart" USING btree ("market_id");
CREATE UNIQUE INDEX IF NOT EXISTS "market_interval_uniq_idx" ON chart USING btree ("market_id", "interval");

-- Foreign Key to Market
ALTER TABLE "chart" ADD CONSTRAINT "fk_market_id" FOREIGN KEY ("market_id") REFERENCES "market"("id");

-- Ticks

CREATE SEQUENCE IF NOT EXISTS "tick_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


CREATE TABLE IF NOT EXISTS "tick" (
    "id" integer NOT NULL DEFAULT nextval('tick_id_seq'::regclass),
    "chart_id" integer NOT NULL,
    "open" text NOT NULL,
    "close" text NOT NULL,
    "high" text NOT NULL,
    "low" text NOT NULL,
    "day" integer NOT NULL,
    "volume" text NOT NULL,
    "timestamp" bigint NOT NULL
);

-- Primary Key Tick
ALTER TABLE "tick" DROP CONSTRAINT IF EXISTS "tick_pkey";
ALTER TABLE "tick" ADD CONSTRAINT "tick_pkey" PRIMARY KEY (id);

-- Indices Tick
CREATE INDEX IF NOT EXISTS "chart_id_idx" ON "tick" USING btree ("chart_id");
CREATE UNIQUE INDEX IF NOT EXISTS "chart_id_timestamp_uniq_idx" ON "tick" USING btree ("chart_id", "timestamp");

-- Foreign Key Tick
ALTER TABLE "tick" ADD CONSTRAINT "chart_fk" FOREIGN KEY ("chart_id") REFERENCES "chart"(id);



-- +migrate Down notransaction
-- Market Tables
DROP TABLE IF EXISTS "market" CASCADE;
DROP TABLE IF EXISTS "exchange" CASCADE;
DROP TABLE IF EXISTS "chart" CASCADE;
DROP TABLE IF EXISTS "tick" CASCADE;