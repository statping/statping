CREATE TABLE core (
    name VARCHAR(50),
    description text,
    config VARCHAR(50),
    api_key VARCHAR(50),
    api_secret VARCHAR(50),
    style text,
    footer text,
    domain text,
    version VARCHAR(50),
    migration_id INT(6) NOT NULL DEFAULT 0,
    use_cdn BOOL NOT NULL DEFAULT '0'
);
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password text,
    email VARCHAR (50),
    api_key VARCHAR(50),
    api_secret VARCHAR(50),
    administrator BOOL NOT NULL DEFAULT '0',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    pushover_key VARCHAR (32) NOT NULL DEFAULT '',
    INDEX (id),
    UNIQUE (username, email)
);
CREATE TABLE services (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    domain text,
    check_type text,
    method VARCHAR(50),
    port INT(6),
    expected text,
    expected_status INT(6),
    check_interval int(11),
    post_data text,
    order_id integer,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX (id)
);
CREATE TABLE hits (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    service INTEGER NOT NULL,
    latency float,
    created_at DATETIME,
    INDEX (id, service),
    FOREIGN KEY (service) REFERENCES services(id) ON DELETE CASCADE
);
CREATE TABLE failures (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    issue text,
    method text,
    service INTEGER NOT NULL,
    created_at DATETIME,
    INDEX (id, service),
    FOREIGN KEY (service) REFERENCES services(id) ON DELETE CASCADE
);
CREATE TABLE checkins (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    service INTEGER NOT NULL,
    check_interval integer,
    api text,
    created_at DATETIME,
    INDEX (id, service),
    FOREIGN KEY (service) REFERENCES services(id) ON DELETE CASCADE
);
CREATE TABLE communication (
    id SERIAL PRIMARY KEY,
    method text,
    host text,
    port integer,
    username text,
    password text,
    var1 text,
    var2 text,
    api_key text,
    api_secret text,
    enabled BOOL NOT NULL DEFAULT '0',
    removable BOOL NOT NULL DEFAULT '0',
    limits integer,
    created_at DATETIME
);