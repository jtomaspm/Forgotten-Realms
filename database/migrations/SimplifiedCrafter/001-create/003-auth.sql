CREATE TABLE `Sessions` (
    `Id`            CHAR(36)        PRIMARY KEY DEFAULT (UUID()),
    `AccountId`     CHAR(36)        NOT NULL,
    `Token`         VARCHAR(255)    UNIQUE NOT NULL,
    `CreatedAt`     TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    `ExpiresAt`     TIMESTAMP       NOT NULL,
    FOREIGN KEY (`AccountId`) REFERENCES `Accounts`(`Id`) ON DELETE CASCADE
);

CREATE TABLE `Logins` (
    `Id`            CHAR(36)        PRIMARY KEY DEFAULT (UUID()),
    `AccountId`     CHAR(36)        NOT NULL,
    `CreatedAt`     TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    `IpAddress`     VARCHAR(45),
    FOREIGN KEY (`AccountId`) REFERENCES `Accounts`(`Id`) ON DELETE CASCADE
);