-- =========================
-- Players
-- =========================
CREATE TABLE IF NOT EXISTS game.players (
    account_id INT PRIMARY KEY,
    id INT NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL UNIQUE,
    faction settings.faction NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_Players_Factions"
        FOREIGN KEY (faction_id) REFERENCES game.factions(id)
);

-- Index for faster lookups on Player ID
-- use unqualified index name to avoid schema-qualification issues
CREATE INDEX IF NOT EXISTS "IDX_Players_PlayerID"
ON game.players (id);

CREATE TRIGGER trigger_set_updated_at_players
BEFORE UPDATE ON game.players
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- =========================
-- Unique Requests (Create Player, etc)
-- =========================
CREATE TABLE IF NOT EXISTS game.unique_requests (
    id BIGSERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    hash VARCHAR(64) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);