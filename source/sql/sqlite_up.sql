CREATE TABLE core (
    name text,
    description text,
    config text,
    api_key text,
    api_secret text,
    style text,
    footer text,
    domain text,
    version text,
    migration_id integer default 0,
    use_cdn bool default false
);
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username text NOT NULL UNIQUE,
    password text,
    email text,
    api_key text,
    api_secret text,
    administrator bool,
    created_at TIMESTAMP,
    UNIQUE (username, email)
);
CREATE TABLE services (
    id INTEGER PRIMARY KEY,
    name text,
    domain text,
    check_type text,
    method text,
    port integer,
    expected text,
    expected_status integer,
    check_interval integer,
    post_data text,
    order_id integer default 0,
    timeout integer default 30,
    created_at TIMESTAMP
);

CREATE TABLE hits (
    id INTEGER PRIMARY KEY,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    latency float,
    created_at TIMESTAMP
);

CREATE TABLE failures (
    id INTEGER PRIMARY KEY,
    issue text,
    method text,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMP
);

CREATE TABLE checkins (
    id INTEGER PRIMARY KEY,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    check_interval integer,
    api text,
    created_at TIMESTAMP
);

CREATE TABLE communication (
    id INTEGER PRIMARY KEY,
    method text,
    host text,
    port integer,
    username text,
    password text,
    var1 text,
    var2 text,
    api_key text,
    api_secret text,
    enabled boolean,
    removable boolean,
    limits integer,
    created_at TIMESTAMP
);


CREATE INDEX idx_hits ON hits(service);
CREATE INDEX idx_failures ON failures(service);
CREATE INDEX idx_checkins ON checkins(service);
