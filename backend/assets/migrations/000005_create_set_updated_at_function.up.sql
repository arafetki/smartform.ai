CREATE
OR REPLACE FUNCTION set_updated_at () RETURNS TRIGGER AS '
BEGIN
    NEW.updated_at := NOW();
    RETURN NEW;
END;
' LANGUAGE plpgsql;

CREATE TRIGGER set_users_updated_at_trigger BEFORE
UPDATE ON users FOR EACH ROW EXECUTE FUNCTION set_updated_at ();

CREATE TRIGGER set_forms_updated_at_trigger BEFORE
UPDATE ON forms FOR EACH ROW EXECUTE FUNCTION set_updated_at ();