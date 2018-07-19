CREATE EXTENSION IF NOT EXISTS "hstore";

CREATE TABLE endpoint
(
    id serial,

    ep text,
    d text,
    lt integer DEFAULT 90000 CHECK (lt >= 60 AND lt <= 4294967295),
    base text,

    PRIMARY KEY(id)
);

CREATE TABLE resource
(
    id serial,

    path text,
    ct integer,
    anchor text,
    extra_attrs hstore,

    endpoint integer NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(endpoint) REFERENCES endpoint(id)
);

-- vim: ai ts=4 sw=4 et sts=4 ft=sql
