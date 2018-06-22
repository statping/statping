CREATE TABLE core (
    name text,
    description text,
    config text,
    api_key text,
    api_secret text,
    style text,
    footer text,
    version text
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username text,
    password text,
    email text,
    api_key text,
    api_secret text,
    created_at TIMESTAMP
);

CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name text,
    domain text,
    check_type text,
    method text,
    port integer,
    expected text,
    expected_status integer,
    check_interval integer,
    created_at TIMESTAMP
);

CREATE TABLE hits (
    id SERIAL PRIMARY KEY,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    latency float,
    created_at TIMESTAMP WITHOUT TIME zone
);

CREATE TABLE failures (
    id SERIAL PRIMARY KEY,
    issue text,
    method text,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMP WITHOUT TIME zone
);

CREATE TABLE checkins (
    id SERIAL PRIMARY KEY,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    check_interval integer,
    api text,
    created_at TIMESTAMP
);


CREATE INDEX idx_hits ON hits(service);
CREATE INDEX idx_failures ON failures(service);
CREATE INDEX idx_checkins ON checkins(service);