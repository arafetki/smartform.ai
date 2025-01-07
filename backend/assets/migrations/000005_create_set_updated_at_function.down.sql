DROP TRIGGER IF EXISTS set_users_updated_at_trigger ON core.users;

DROP TRIGGER IF EXISTS set_forms_updated_at_trigger ON core.users;

DROP FUNCTION IF EXISTS core.set_updated_at () CASCADE;