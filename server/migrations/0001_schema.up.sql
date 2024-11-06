BEGIN;

-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_uuid uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),
    password_hash varchar,
    password_uuid uuid UNIQUE,
    email varchar UNIQUE NOT NULL,
    first_name varchar NOT NULL,
    last_name varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now())
);

-- Password reset table
CREATE TABLE users_password_reset (
    user_id int UNIQUE NOT NULL REFERENCES users(id),
    reset_token uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),
    valid_until timestamp DEFAULT (now() + '1 day'::interval) NOT NULL
);

-- Title table
CREATE TABLE titles (
    id SERIAL PRIMARY KEY,
    name varchar UNIQUE NOT NULL,
    created_on timestamp NOT NULL DEFAULT now()
);

-- Environments table
CREATE TABLE environments (
    id SERIAL PRIMARY KEY,
    title int UNIQUE NOT NULL REFERENCES titles(id),
    key uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),
    name varchar NOT NULL,
    created_on timestamp NOT NULL DEFAULT now(),

    UNIQUE (title, name) -- Environment names should be unique per title
);

-- Enum for different client types
CREATE TYPE log_client_type AS ENUM ('Client', 'Server');

-- Logs table
CREATE TABLE logs (
    id BIGSERIAL PRIMARY KEY,
    log_uuid uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),
    environment int NOT NULL REFERENCES environments(id),
    platform varchar NOT NULL,
    log_type log_client_type NOT NULL,
    created_on timestamp NOT NULL DEFAULT now(),
    log_start timestamp,
    log_end timestamp
);

-- Logs chunks table
CREATE TABLE logs_chunks (
    chunk SERIAL PRIMARY KEY,
    log bigint NOT NULL REFERENCES logs(id),
    created_on timestamp NOT NULL DEFAULT now(),
    blob_path varchar UNIQUE NOT NULL,
    chunk_start timestamp,
    chunk_end timestamp,
    category_counts jsonb NOT NULL,
    severity_counts jsonb NOT NULL,

    CONSTRAINT log_chunks_category_counts_object CHECK (jsonb_typeof(category_counts) = 'object'),
    CONSTRAINT log_chunks_severity_counts_object CHECK (jsonb_typeof(severity_counts) = 'object')
);

-- Logs metadata table
CREATE TABLE logs_metadata (
    log bigint NOT NULL REFERENCES logs(id),
    saved_on timestamp NOT NULL DEFAULT now(),
    metadata jsonb NOT NULL,

    CONSTRAINT logs_metadata_metadata_object CHECK (jsonb_typeof(metadata) = 'object')
);

-- Log links table
CREATE TABLE logs_links (
    source bigint NOT NULL REFERENCES logs(id),
    link bigint NOT NULL REFERENCES logs(id),
    created_on timestamp NOT NULL DEFAULT now(),
    description varchar NOT NULL,

    UNIQUE (source, link) -- Source/link combination should be unique
);

COMMIT;
