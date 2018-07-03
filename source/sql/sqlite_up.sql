CREATE TABLE core (
    name text,
    description text,
    config text,
    api_key text,
    api_secret text,
    style text,
    footer text,
    domain text,
    version text
);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username text,
    password text,
    email text,
    api_key text,
    api_secret text,
    administrator bool,
    created_at TIMESTAMP
);

CREATE TABLE services (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name text,
    domain text,
    check_type text,
    method text,
    port integer,
    expected text,
    expected_status integer,
    check_interval integer,
    post_data text,
    order_id integer,
    created_at TIMESTAMP
);

CREATE TABLE hits (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    latency float,
    created_at TIMESTAMP
);

CREATE TABLE failures (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    issue text,
    method text,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMP
);

CREATE TABLE checkins (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    check_interval integer,
    api text,
    created_at TIMESTAMP
);

CREATE TABLE communication (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
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