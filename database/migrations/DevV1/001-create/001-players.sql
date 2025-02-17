CREATE TABLE `Players` (
    `Id`        CHAR(36)    PRIMARY KEY,
    `AccountId` CHAR(36)    NOT NULL,
    `Name`      VARCHAR(50) NOT NULL,
    `CreatedAt` TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);