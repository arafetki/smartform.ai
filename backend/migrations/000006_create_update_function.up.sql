CREATE
OR REPLACE FUNCTION core.set_updated_at_now () RETURNS TRIGGER AS '
BEGIN
    NEW.updated_at := NOW();
    RETURN NEW;
END;
' LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_now_users_trigger BEFORE
UPDATE ON core.users FOR EACH ROW EXECUTE FUNCTION core.set_updated_at_now ();

CREATE TRIGGER set_updated_at_now_forms_trigger BEFORE
UPDATE ON core.forms FOR EACH ROW EXECUTE FUNCTION core.set_updated_at_now ();