CREATE TABLE
    IF NOT EXISTS forms (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id varchar(36) NOT NULL,
        title varchar(255) NOT NULL,
        description text,
        fields jsonb NOT NULL DEFAULT '[]',
        is_published boolean NOT NULL DEFAULT false,
        created_at timestamptz NOT NULL DEFAULT now (),
        updated_at timestamptz,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE INDEX idx_forms_user_id ON forms (user_id);