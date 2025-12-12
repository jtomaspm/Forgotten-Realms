DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'faction') THEN
        CREATE TYPE settings.faction AS ENUM ('caldari', 'varnak', 'dawnhold', 'forgotten');
    END IF;
END $$;

CREATE TABLE settings.realm (
    name VARCHAR(50) UNIQUE PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);