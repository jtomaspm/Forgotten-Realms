CREATE TABLE villages (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (coord_x, coord_y) REFERENCES locations (coord_x, coord_y),
    PRIMARY KEY(coord_x, coord_y)
);

CREATE TABLE village_buildings (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    headquarters SMALLINT NOT NULL DEFAULT 1,
    forest SMALLINT NOT NULL DEFAULT 0,
    quarry SMALLINT NOT NULL DEFAULT 0,
    mine SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (coord_x, coord_y) REFERENCES villages (coord_x, coord_y),
    PRIMARY KEY(coord_x, coord_y)
);

CREATE TABLE village_resources (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    wood INTEGER NOT NULL DEFAULT 0,
    stone INTEGER NOT NULL DEFAULT 0,
    metal INTEGER NOT NULL DEFAULT 0,
    wood_hour INTEGER NOT NULL DEFAULT 3600,
    stone_hour INTEGER NOT NULL DEFAULT 3600,
    metal_hour INTEGER NOT NULL DEFAULT 3600,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (coord_x, coord_y) REFERENCES villages (coord_x, coord_y),
    PRIMARY KEY(coord_x, coord_y)
);

CREATE TABLE village_troops (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (coord_x, coord_y) REFERENCES villages (coord_x, coord_y),
    PRIMARY KEY(coord_x, coord_y)
);

CREATE TRIGGER trigger_set_updated_at_villages
BEFORE UPDATE ON villages
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
