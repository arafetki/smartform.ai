CREATE TABLE
    IF NOT EXISTS core.forms (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id varchar(36) NOT NULL,
        title varchar(255) NOT NULL,
        description text,
        fields jsonb NOT NULL DEFAULT '[]',
        published boolean NOT NULL DEFAULT false,
        created_at timestamptz NOT NULL DEFAULT now (),
        updated_at timestamptz,
        FOREIGN KEY (user_id) REFERENCES core.users (id) ON DELETE CASCADE
    )