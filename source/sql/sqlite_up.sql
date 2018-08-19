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
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username text NOT NULL UNIQUE,
    password text,
    email text,
    api_key text,
    api_secret text,
    administrator bool,
    created_at DATETIME,
    UNIQUE (username, email)
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
    order_id integer default 0,
    timeout integer default 30,
    created_at DATETIME
);

CREATE TABLE hits (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    latency float,
    created_at DATETIME
);

CREATE TABLE failures (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    issue text,
    method text,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at DATETIME
);

CREATE TABLE checkins (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    check_interval integer,
    api text,
    created_at DATETIME
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
    created_at DATETIME
);


CREATE INDEX idx_hits ON hits(service);
CREATE INDEX idx_failures ON failures(service);
CREATE INDEX idx_checkins ON checkins(service);