DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'faction') THEN
        CREATE TYPE faction AS ENUM ('caldari', 'varnak', 'dawnhold', 'forgotten');
    END IF;
END $$;

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'building') THEN
        CREATE TYPE building AS ENUM ('headquarters', 'farm', 'warehouse', 'forest', 'quarry', 'mine');
    END IF;
END $$;

CREATE TABLE realm_settings (
    speed REAL NOT NULL,
    created_at TIMESTAMPTZ PRIMARY KEY DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE building_levels (
    faction faction NOT NULL,
    building building NOT NULL,
    level SMALLINT NOT NULL,
    wood INTEGER NOT NULL,
    stone INTEGER NOT NULL,
    metal INTEGER NOT NULL,
    population SMALLINT NOT NULL,
    special INTEGER,
    time_seconds INTEGER NOT NULL,
    created_at TIMESTAMPTZ PRIMARY KEY DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(faction, building, level)
);

CREATE TRIGGER trigger_set_updated_at_realm_settings
BEFORE UPDATE ON realm_settings
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trigger_set_updated_at_building_levels
BEFORE UPDATE ON building_levels
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();