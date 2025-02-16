CREATE TABLE `Worlds` (
    `Id`            CHAR(36)        PRIMARY KEY DEFAULT (UUID()),
    `Name`          VARCHAR(50)     NOT NULL,
    `Database`      VARCHAR(255)    NOT NULL,
    `Speed`         FLOAT           NOT NULL,
    `GameVersion`   VARCHAR(20)     NOT NULL,
    `CreatedAt`     TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);