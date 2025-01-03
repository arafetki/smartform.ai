ALTER TABLE core.forms
DROP CONSTRAINT IF EXISTS fk_form_settings;

ALTER TABLE core.forms
DROP COLUMN IF EXISTS settings_id;