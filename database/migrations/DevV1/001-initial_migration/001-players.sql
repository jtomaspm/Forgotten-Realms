CREATE TABLE `Players` (
    `Id` CHAR(36) PRIMARY KEY,        -- guid
    `AccountId` CHAR(36) NOT NULL,    -- guid
    `Name` VARCHAR(50) NOT NULL       -- string
);