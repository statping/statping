=========================================== 1530986614
ALTER TABLE users ADD COLUMN pushover_key VARCHAR (32) NOT NULL DEFAULT '';
=========================================== 1530841150
ALTER TABLE core ADD COLUMN use_cdn BOOL NOT NULL DEFAULT '0';
=========================================== 1
ALTER TABLE core ADD COLUMN migration_id INT(6) NOT NULL DEFAULT 0;