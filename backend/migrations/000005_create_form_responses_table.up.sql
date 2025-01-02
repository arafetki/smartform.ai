CREATE TABLE
    IF NOT EXISTS core.form_responses (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        form_id uuid NOT NULL,
        data jsonb NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now (),
        CONSTRAINT fk_form_response FOREIGN KEY (form_id) REFERENCES core.forms (id) ON DELETE CASCADE
    )