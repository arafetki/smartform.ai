CREATE TABLE
    IF NOT EXISTS core.forms (
        id uuid PRIMARY KEY DEFAULT core.gen_random_uuid (),
        user_id uuid NOT NULL,
        title varchar(255) NOT NULL,
        description text,
        fields jsonb NOT NULL DEFAULT '[]',
        view_count bigint NOT NULL DEFAULT 0,
        published boolean NOT NULL DEFAULT false,
        created_at timestamptz NOT NULL DEFAULT now (),
        updated_at timestamptz,
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES core.users (id) ON DELETE CASCADE
    )