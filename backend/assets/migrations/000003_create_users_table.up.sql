CREATE TABLE
    IF NOT EXISTS core.users (
        id varchar(36) PRIMARY KEY,
        is_verified boolean NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now (),
        updated_at timestamptz
    );