CREATE TABLE
    IF NOT EXISTS core.form_responses (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        form_id uuid NOT NULL,
        content jsonb NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now (),
        FOREIGN KEY (form_id) REFERENCES core.forms (id) ON DELETE CASCADE
    );

CREATE INDEX idx_responses_form_id ON core.form_responses (form_id);