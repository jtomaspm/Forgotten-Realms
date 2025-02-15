CREATE TABLE `Accounts` (
    `Id` CHAR(36) PRIMARY KEY DEFAULT (UUID()),      -- guid
    `Name` VARCHAR(50) NOT NULL,    -- string
    `Email` VARCHAR(255) NOT NULL,    -- string
    `CreatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP -- datetime
);