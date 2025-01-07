CREATE TABLE
    IF NOT EXISTS core.users (
        id uuid PRIMARY KEY,
        avatar_url text,
        is_verified boolean NOT NULL,
        created_at timestamptz NOT NULL,
        updated_at timestamptz
    );