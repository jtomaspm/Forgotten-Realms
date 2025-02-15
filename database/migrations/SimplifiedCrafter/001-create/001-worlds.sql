CREATE TABLE `Worlds` (
    `Id` CHAR(36) PRIMARY KEY DEFAULT (UUID()),      -- guid
    `Name` VARCHAR(50) NOT NULL,    -- string
    `Database` VARCHAR(255) NOT NULL,    -- string
    `Speed` FLOAT,                  -- float
    `GameVersion` VARCHAR(20),       -- string
    `CreatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP -- datetime
);