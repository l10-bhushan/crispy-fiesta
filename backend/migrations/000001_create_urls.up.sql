CREATE TABLE
    urls (
        id BIGSERIAL PRIMARY KEY,
        short_code VARCHAR(15) NOT NULL UNIQUE,
        original_url TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        expires_at TIMESTAMPTZ NULL
    )