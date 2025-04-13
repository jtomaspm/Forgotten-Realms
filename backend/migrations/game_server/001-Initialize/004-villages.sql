CREATE TABLE villages (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    player_id UUID NOT NULL,
    faction faction NOT NULL, -- Still stored for quick access / indexing
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(coord_x, coord_y),
    FOREIGN KEY (player_id) REFERENCES players (id)
);

CREATE OR REPLACE FUNCTION set_village_faction()
RETURNS TRIGGER AS $$
BEGIN
    SELECT p.faction INTO NEW.faction
    FROM players p
    WHERE p.id = NEW.player_id;

    IF NEW.faction IS NULL THEN
        RAISE EXCEPTION 'Player ID % not found or faction is NULL', NEW.player_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_village_faction
BEFORE INSERT ON villages
FOR EACH ROW
EXECUTE FUNCTION set_village_faction();

CREATE INDEX idx_villages_player_id ON villages(player_id);

CREATE TRIGGER trigger_set_updated_at_villages
BEFORE UPDATE ON villages
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE village_buildings (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    headquarters SMALLINT NOT NULL DEFAULT 1,
    warehouse SMALLINT NOT NULL DEFAULT 0,
    farm SMALLINT NOT NULL DEFAULT 0,
    forest SMALLINT NOT NULL DEFAULT 0,
    quarry SMALLINT NOT NULL DEFAULT 0,
    mine SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (coord_x, coord_y) REFERENCES villages (coord_x, coord_y),
    PRIMARY KEY(coord_x, coord_y)
);

CREATE TRIGGER trigger_set_updated_at_village_buildings
BEFORE UPDATE ON village_buildings
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE village_resources (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    wood INTEGER NOT NULL DEFAULT 0,
    stone INTEGER NOT NULL DEFAULT 0,
    metal INTEGER NOT NULL DEFAULT 0,
    wood_hour INTEGER NOT NULL DEFAULT 0,
    stone_hour INTEGER NOT NULL DEFAULT 0,
    metal_hour INTEGER NOT NULL DEFAULT 0,
    capacity INTEGER NOT NULL DEFAULT 0,
    population INTEGER NOT NULL DEFAULT 0,
    maximum_population INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (coord_x, coord_y) REFERENCES villages (coord_x, coord_y),
    PRIMARY KEY(coord_x, coord_y)
);

CREATE TRIGGER trigger_set_updated_at_village_resources
BEFORE UPDATE ON village_resources
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
