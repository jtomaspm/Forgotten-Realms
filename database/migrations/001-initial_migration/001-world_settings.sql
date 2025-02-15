CREATE TABLE `Worlds` (
    `Id` CHAR(36) PRIMARY KEY,      -- guid
    `Name` VARCHAR(50) NOT NULL,    -- string
    `Speed` FLOAT,                  -- float
    `GameVersion` VARCHAR(20)       -- string
);