-- Ensure schema exists
CREATE SCHEMA IF NOT EXISTS game;

-- =========================
-- Factions
-- =========================
CREATE TABLE IF NOT EXISTS game.factions (
    id INT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO game.factions (id, name) VALUES
(1, 'caldari'),
(2, 'varnak'),
(3, 'dawnhold')
ON CONFLICT DO NOTHING;


-- =========================
-- Players
-- =========================
CREATE TABLE IF NOT EXISTS game.players (
    account_id INT PRIMARY KEY,
    id INT NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL UNIQUE,
    faction_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_Players_Factions"
        FOREIGN KEY (faction_id) REFERENCES game.factions(id)
);

-- Index for faster lookups on Player ID
-- use unqualified index name to avoid schema-qualification issues
CREATE INDEX IF NOT EXISTS "IDX_Players_PlayerID"
ON game.players (id);


-- =========================
-- Trigger Function
-- =========================
CREATE OR REPLACE FUNCTION game.update_players_updatedat()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger
-- NOTE: trigger name must NOT be schema-qualified
CREATE TRIGGER "trg_UpdatePlayersUpdatedAt"
BEFORE UPDATE ON game.players
FOR EACH ROW
EXECUTE FUNCTION game.update_players_updatedat();


-- =========================
-- Unique Requests (Create Player, etc)
-- =========================
CREATE TABLE IF NOT EXISTS game.unique_requests (
    id BIGSERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    hash VARCHAR(64) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);