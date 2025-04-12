CREATE TABLE locations (
    coord_x SMALLINT NOT NULL,
    coord_y SMALLINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(coord_x, coord_y)
);

CREATE TRIGGER trigger_set_updated_at_locations
BEFORE UPDATE ON locations
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();