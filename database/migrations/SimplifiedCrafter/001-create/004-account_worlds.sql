CREATE TABLE `AccountWorlds` (
    `AccountId`     CHAR(36)        NOT NULL,
    `WorldId`       CHAR(36)        NOT NULL,
    `CreatedAt`     TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`AccountId`)   REFERENCES `Accounts`(`Id`) ON DELETE CASCADE,
    FOREIGN KEY (`WorldId`)     REFERENCES `Worlds`(`Id`) ON DELETE CASCADE,
    PRIMARY KEY (`AccountId`, `WorldId`)
);