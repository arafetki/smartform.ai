CREATE TABLE
    IF NOT EXISTS core.form_views (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        form_id uuid NOT NULL,
        ip_address INET NOT NULL,
        viewed_at timestamptz DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (form_id) REFERENCES core.forms (id) ON DELETE CASCADE,
        UNIQUE (form_id, ip_address)
    );

CREATE INDEX idx_views_form_id ON core.form_responses (form_id);