CREATE TABLE `Accounts` (
    `Id`        CHAR(36)        PRIMARY KEY DEFAULT (UUID()),
    `Name`      VARCHAR(50)     NOT NULL,
    `Email`     VARCHAR(255)    NOT NULL,
    `CreatedAt` DATETIME        DEFAULT CURRENT_TIMESTAMP
);