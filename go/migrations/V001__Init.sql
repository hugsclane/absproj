CREATE TABLE data_set (
    id PRIMARY KEY,
    "key" VARCHAR(64) NOT NULL,
    "data" BLOB
);

CREATE UNIQUE INDEX data_set_key ON TABLE data_set ("key");