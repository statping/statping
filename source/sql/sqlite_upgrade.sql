=========================================== 1530986614
ALTER TABLE users ADD COLUMN pushover_key TEXT NOT NULL DEFAULT '';
=========================================== 1530841150
ALTER TABLE core ADD COLUMN use_cdn bool DEFAULT FALSE;
=========================================== 1
ALTER TABLE core ADD COLUMN migration_id integer NOT NULL DEFAULT 0;