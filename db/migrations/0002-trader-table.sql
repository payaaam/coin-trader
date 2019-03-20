-- +migrate Up notransaction

CREATE TABLE IF NOT EXISTS "trades" (
    "id" integer NOT NULL
);


-- +migrate Down notransaction

DROP TABLE IF EXISTS "trades" CASCADE;