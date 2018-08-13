=========================================== 1534178020
UPDATE services SET order_id=0 WHERE order_id IS NULL;
=========================================== 1532068515
ALTER TABLE services ALTER COLUMN order_id integer DEFAULT 0;
ALTER TABLE services ADD COLUMN timeout integer DEFAULT 30;
=========================================== 1530841150
ALTER TABLE core ADD COLUMN use_cdn bool DEFAULT FALSE;
=========================================== 1
ALTER TABLE core ADD COLUMN migration_id integer default 0 NOT NULL;