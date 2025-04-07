SET TIME ZONE 'UTC';

CREATE EXTENSION IF NOT EXISTS pgcrypto;

DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_proc 
        JOIN pg_namespace ON pg_proc.pronamespace = pg_namespace.oid
        WHERE proname = 'set_updated_at' 
        AND nspname = 'public'
    ) THEN
        EXECUTE 'CREATE FUNCTION set_updated_at() 
        RETURNS TRIGGER AS $func$
        BEGIN
            NEW.updated_at = NOW();
            RETURN NEW;
        END;
        $func$ LANGUAGE plpgsql;';
    END IF;
END $$;