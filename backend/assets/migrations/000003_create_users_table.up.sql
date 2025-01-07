CREATE TABLE
    IF NOT EXISTS core.users (
        id uuid PRIMARY KEY,
        email varchar(255) UNIQUE NOT NULL,
        name varchar(255) NOT NULL,
        phone_number varchar(25) UNIQUE NOT NULL,
        is_verified boolean NOT NULL,
        avatar_url text,
        created_at timestamptz NOT NULL,
        updated_at timestamptz
    )