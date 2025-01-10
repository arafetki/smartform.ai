DROP TRIGGER IF EXISTS set_users_updated_at_trigger ON users;

DROP TRIGGER IF EXISTS set_forms_updated_at_trigger ON users;

DROP FUNCTION IF EXISTS set_updated_at () CASCADE;