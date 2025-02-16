CREATE TABLE `Roles` (
    `Name` VARCHAR(50) PRIMARY KEY
);

CREATE TABLE `Accounts` (
    `Id`            CHAR(36)        PRIMARY KEY DEFAULT (UUID()),
    `ExternalId`    VARCHAR(255),
    `Name`          VARCHAR(50)     NOT NULL,
    `Email`         VARCHAR(255)    NOT NULL,
    `Role`          VARCHAR(50)     NOT NULL,
    `CreatedAt`     TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`Role`) REFERENCES `Roles`(`Name`) ON DELETE CASCADE
);