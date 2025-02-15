CREATE TABLE `Migrations` (
    `Name` VARCHAR(255) NOT NULL,        
    `Database` VARCHAR(255) NOT NULL,    
    `CreatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP,                    
    PRIMARY KEY (`Name`, `Database`)  -- Composite primary key
);