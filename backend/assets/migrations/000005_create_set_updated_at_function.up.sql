CREATE
OR REPLACE FUNCTION core.set_updated_at () RETURNS TRIGGER AS '
BEGIN
    NEW.updated_at := NOW();
    RETURN NEW;
END;
' LANGUAGE plpgsql;

CREATE TRIGGER set_users_updated_at_trigger BEFORE
UPDATE ON core.users FOR EACH ROW EXECUTE FUNCTION core.set_updated_at ();

CREATE TRIGGER set_forms_updated_at_trigger BEFORE
UPDATE ON core.forms FOR EACH ROW EXECUTE FUNCTION core.set_updated_at ();