CREATE TABLE IF NOT EXISTS users (
    id         BIGSERIAL PRIMARY KEY,
    email      VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS apps (
    id        BIGSERIAL PRIMARY KEY,
    name      VARCHAR(255) NOT NULL,
    owner_id  BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_apps_owner_id ON apps(owner_id);

CREATE TABLE IF NOT EXISTS api_keys (
    id          BIGSERIAL PRIMARY KEY,
    app_id      BIGINT NOT NULL REFERENCES apps(id) ON DELETE CASCADE,
    access_key  VARCHAR(64) NOT NULL UNIQUE,
    secret_key  VARCHAR(128) NOT NULL,
    permissions TEXT NOT NULL DEFAULT 'read,write',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_api_keys_app_id ON api_keys(app_id);

CREATE TABLE IF NOT EXISTS buckets (
    id         BIGSERIAL PRIMARY KEY,
    app_id     BIGINT NOT NULL REFERENCES apps(id) ON DELETE CASCADE,
    name       VARCHAR(63) NOT NULL,
    visibility VARCHAR(10) NOT NULL DEFAULT 'private',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE(app_id, name)
);

CREATE INDEX idx_buckets_app_id ON buckets(app_id);

CREATE TABLE IF NOT EXISTS objects (
    id           BIGSERIAL PRIMARY KEY,
    bucket_id    BIGINT NOT NULL REFERENCES buckets(id) ON DELETE CASCADE,
    key          VARCHAR(1024) NOT NULL,
    size         BIGINT NOT NULL,
    mime         VARCHAR(255) NOT NULL,
    storage_path TEXT NOT NULL,
    checksum     VARCHAR(128) NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE(bucket_id, key)
);

CREATE INDEX idx_objects_bucket_id ON objects(bucket_id);
