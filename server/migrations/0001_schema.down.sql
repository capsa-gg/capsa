BEGIN;

DROP TABLE logs_links;
DROP TABLE logs_metadata;
DROP TABLE logs_chunks;
DROP TABLE logs;
DROP TYPE log_client_type;
DROP TABLE environments;
DROP TABLE titles;
DROP TABLE users_password_reset;
DROP TABLE users;

COMMIT;
