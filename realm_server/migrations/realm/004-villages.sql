CREATE TABLE game.villages (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    player_id INT NOT NULL,
    points SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(coord_x, coord_y),
    FOREIGN KEY (player_id) REFERENCES game.players (id)
);

CREATE INDEX idx_villages_player_id ON game.villages(player_id);

CREATE TRIGGER trigger_set_updated_at_villages
BEFORE UPDATE ON game.villages
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE game.village_buildings (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    headquarters SMALLINT NOT NULL DEFAULT 0,
    warehouse SMALLINT NOT NULL DEFAULT 0,
    farm SMALLINT NOT NULL DEFAULT 0,
    forest SMALLINT NOT NULL DEFAULT 0,
    quarry SMALLINT NOT NULL DEFAULT 0,
    mine SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (coord_x, coord_y) REFERENCES game.villages (coord_x, coord_y),
    PRIMARY KEY(coord_x, coord_y)
);

CREATE TRIGGER trigger_set_updated_at_village_buildings
BEFORE UPDATE ON game.village_buildings
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TABLE game.village_resources (
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
BEFORE UPDATE ON game.village_resources
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
