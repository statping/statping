CREATE TABLE core (
    name VARCHAR(50),
    description text,
    config VARCHAR(50),
    api_key VARCHAR(50),
    api_secret VARCHAR(50),
    style text,
    footer text,
    version VARCHAR(50)
);
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password text,
    email text,
    api_key VARCHAR(50),
    api_secret VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX (id)
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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX (id)
);
CREATE TABLE hits (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    service INTEGER NOT NULL,
    latency float,
    created_at TIMESTAMP,
    INDEX (id, service),
    FOREIGN KEY (service) REFERENCES services(id)
);
CREATE TABLE failures (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    issue text,
    service INTEGER NOT NULL,
    created_at TIMESTAMP,
    INDEX (id, service),
    FOREIGN KEY (service) REFERENCES services(id)
);
CREATE TABLE checkins (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    service INTEGER NOT NULL,
    check_interval integer,
    api text,
    created_at TIMESTAMP,
    INDEX (id, service),
    FOREIGN KEY (service) REFERENCES services(id)
);