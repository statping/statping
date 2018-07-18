=========================================== 1531891670
ALTER TABLE services ALTER COLUMN order_id SET DEFAULT 0;
ALTER TABLE services ADD COLUMN timeout integer DEFAULT 30;
=========================================== 1530841150
ALTER TABLE core ADD COLUMN use_cdn bool DEFAULT FALSE;
=========================================== 1
ALTER TABLE core ADD COLUMN migration_id integer default 0 NOT NULL;