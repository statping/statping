=========================================== 1530986614
ALTER TABLE users ADD COLUMN pushover_key text default '';
=========================================== 1530841150
ALTER TABLE core ADD COLUMN use_cdn bool DEFAULT FALSE;
=========================================== 1
ALTER TABLE core ADD COLUMN migration_id integer default 0 NOT NULL;