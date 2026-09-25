BEGIN;
ALTER TABLE users ADD COLUMN role VARCHAR(20) NOT NULL DEFAULT 'user' CONSTRAINT chk_users_role CHECK (role IN ('user', 'admin'));
COMMIT;
