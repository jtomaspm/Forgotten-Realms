CREATE TABLE `Roles` (
    `Name` VARCHAR(50) PRIMARY KEY
);


CREATE TABLE `Accounts` (
    `Id`            CHAR(36)        PRIMARY KEY DEFAULT (UUID()),
    `ExternalId`    VARCHAR(255)    UNIQUE,
    `Name`          VARCHAR(50)     UNIQUE NOT NULL,
    `Email`         VARCHAR(255)    UNIQUE NOT NULL,
    `Role`          VARCHAR(50)     NOT NULL,
    `CreatedAt`     TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    `UpdatedAt`     TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`Role`) REFERENCES `Roles`(`Name`) ON DELETE CASCADE
);


CREATE TABLE `AccountProperties` (
    `AccountId`                 CHAR(36)        PRIMARY KEY,
    `VerificationToken`         CHAR(36)        UNIQUE,
    `TokenExpiresAt`            TIMESTAMP,
    `EmailVerified`             BOOLEAN         DEFAULT FALSE,
    `SendEmailNotifications`    BOOLEAN         DEFAULT FALSE,
    `CreatedAt`                 TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    `UpdatedAt`                 TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`AccountId`) REFERENCES `Accounts`(`Id`) ON DELETE CASCADE
);