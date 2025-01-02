DROP TRIGGER IF EXISTS set_updated_at_now_trigger ON core.users;

DROP FUNCTION IF EXISTS core.set_updated_at_now ();