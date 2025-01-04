ALTER TABLE core.forms
ADD COLUMN settings_id SMALLINT NOT NULL;

ALTER TABLE core.forms ADD CONSTRAINT fk_form_settings FOREIGN KEY (settings_id) REFERENCES core.form_settings (id) ON DELETE CASCADE;