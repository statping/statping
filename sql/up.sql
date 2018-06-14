CREATE TABLE core (
    name text,
    config text,
    api_key text,
    api_secret text,
    version text,
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username text,
    password text,
    key text,
    secret text,
    created_at TIMESTAMP WITHOUT TIME zone
);

CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name text,
    domain text,
    method text,
    port integer,
    expected text,
    expected_status integer,
    interval integer,
    created_at TIMESTAMP WITHOUT TIME zone
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
    service INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMP WITHOUT TIME zone
);

CREATE INDEX idx_hits ON hits(service);
CREATE INDEX idx_failures ON failures(service);

INSERT INTO users (id, username, password, created_at) VALUES (1, 'admin', '$2a$14$sBO5VDKiGPNUa3IUSMRX.OJNIbw/VM5dXOzTjlsjvG6qA987Lfzga', NOW());

INSERT INTO services (id, name, domain, method, port, expected, expected_status, interval, created_at) VALUES (1, 'Statup Demo', 'https://demo.statup.io', 'https', 0, '', 200, 5, NOW());

INSERT INTO services (id, name, domain, method, port, expected, expected_status, interval, created_at) VALUES (2, 'Github', 'https://github.com', 'https', 0, '', 200, 10, NOW());

INSERT INTO services (id, name, domain, method, port, expected, expected_status, interval, created_at) VALUES (3, 'Santa Monica', 'https://www.santamonica.com', 'https', 0, '', 200, 30, NOW());

INSERT INTO services (id, name, domain, method, port, expected, expected_status, interval, created_at) VALUES (4, 'Example JSON', 'https://jsonplaceholder.typicode.com/posts/42', 'https', 0, 'userId', 200, 5, NOW());

INSERT INTO services (id, name, domain, method, port, expected, expected_status, interval, created_at) VALUES (5, 'Token Balance API', 'https://api.tokenbalance.com/token/0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0/0x5c98dc37c0b3ef75476eb6b8ef0d564f7c6af6ae', 'https', 0, 'broken', 200, 8, NOW());

INSERT INTO services (id, name, domain, method, port, expected, expected_status, interval, created_at) VALUES (6, 'Example JSON 2', 'https://jsonplaceholder.typicode.com/posts/42', 'https', 0, 'commodi ullam sint et excepturi error explicabo praesentium voluptas', 200, 13, NOW());

