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
) ENGINE=INNODB;
CREATE TABLE users (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, PRIMARY KEY(id),
    username VARCHAR(50) NOT NULL UNIQUE,
    password text,
    email VARCHAR (50),
    api_key VARCHAR(50),
    api_secret VARCHAR(50),
    administrator BOOL NOT NULL DEFAULT '0',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (username, email)
) ENGINE=INNODB;
CREATE TABLE services (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, PRIMARY KEY(id),
    name VARCHAR(50),
    domain text,
    check_type text,
    method VARCHAR(50),
    port INT(6),
    expected text,
    expected_status INT(6),
    check_interval int(11),
    post_data text,
    order_id integer default 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    timeout INT(6) DEFAULT 30
) ENGINE=INNODB;
CREATE TABLE hits (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, PRIMARY KEY(id),
    service BIGINT(20) UNSIGNED NOT NULL,
    latency float,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service) REFERENCES services(id) ON DELETE CASCADE
) ENGINE=INNODB;
CREATE TABLE failures (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, PRIMARY KEY(id),
    issue text,
    method text,
    service BIGINT(20) UNSIGNED NOT NULL,
    created_at TIMESTAMP,
    FOREIGN KEY (service) REFERENCES services(id) ON DELETE CASCADE
) ENGINE=INNODB;
CREATE TABLE checkins (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, PRIMARY KEY(id),
    service BIGINT(20) UNSIGNED NOT NULL,
    check_interval integer,
    api text,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service) REFERENCES services(id) ON DELETE CASCADE
) ENGINE=INNODB;
CREATE TABLE communication (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, PRIMARY KEY(id),
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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;
