DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'faction') THEN
        CREATE TYPE faction AS ENUM ('caldari', 'varnak', 'dawnhold', 'forgotten');
    END IF;
END $$;

CREATE TABLE settings_realm (
    speed REAL NOT NULL,
    unit_speed REAL NOT NULL,
    created_at TIMESTAMPTZ PRIMARY KEY DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trigger_set_updated_at_settings_realm
BEFORE UPDATE ON settings_realm
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE settings_farm_levels (
    faction faction NOT NULL,
    level SMALLINT NOT NULL,
    wood INTEGER NOT NULL,
    stone INTEGER NOT NULL,
    metal INTEGER NOT NULL,
    population SMALLINT NOT NULL,
    maximum_population SMALLINT NOT NULL,
    time_seconds INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(faction, level)
);

CREATE TRIGGER trigger_set_updated_at_settings_farm_levels
BEFORE UPDATE ON settings_farm_levels
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE settings_forest_levels (
    faction faction NOT NULL,
    level SMALLINT NOT NULL,
    wood INTEGER NOT NULL,
    stone INTEGER NOT NULL,
    metal INTEGER NOT NULL,
    population SMALLINT NOT NULL,
    wood_hour SMALLINT NOT NULL,
    time_seconds INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(faction, level)
);

CREATE TRIGGER trigger_set_updated_at_settings_forest_levels
BEFORE UPDATE ON settings_forest_levels
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE settings_headquarters_levels (
    faction faction NOT NULL,
    level SMALLINT NOT NULL,
    wood INTEGER NOT NULL,
    stone INTEGER NOT NULL,
    metal INTEGER NOT NULL,
    population SMALLINT NOT NULL,
    build_speed_multi_x1000 SMALLINT NOT NULL,
    time_seconds INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(faction, level)
);

CREATE TRIGGER trigger_set_updated_at_settings_headquarters_levels
BEFORE UPDATE ON settings_headquarters_levels
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE settings_mine_levels (
    faction faction NOT NULL,
    level SMALLINT NOT NULL,
    wood INTEGER NOT NULL,
    stone INTEGER NOT NULL,
    metal INTEGER NOT NULL,
    population SMALLINT NOT NULL,
    metal_hour SMALLINT NOT NULL,
    time_seconds INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(faction, level)
);

CREATE TRIGGER trigger_set_updated_at_settings_mine_levels
BEFORE UPDATE ON settings_mine_levels
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE settings_quarry_levels (
    faction faction NOT NULL,
    level SMALLINT NOT NULL,
    wood INTEGER NOT NULL,
    stone INTEGER NOT NULL,
    metal INTEGER NOT NULL,
    population SMALLINT NOT NULL,
    stone_hour SMALLINT NOT NULL,
    time_seconds INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(faction, level)
);

CREATE TRIGGER trigger_set_updated_at_settings_quarry_levels
BEFORE UPDATE ON settings_quarry_levels
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE settings_warehouse_levels (
    faction faction NOT NULL,
    level SMALLINT NOT NULL,
    wood INTEGER NOT NULL,
    stone INTEGER NOT NULL,
    metal INTEGER NOT NULL,
    population SMALLINT NOT NULL,
    capacity INTEGER NOT NULL,
    time_seconds INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(faction, level)
);

CREATE TRIGGER trigger_set_updated_at_settings_warehouse_levels
BEFORE UPDATE ON settings_warehouse_levels
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();