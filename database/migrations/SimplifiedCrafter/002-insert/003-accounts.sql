INSERT INTO `Accounts` 
    (`Name`,                `Email`,                `Role`,     `ExternalId`,   `Source`)
VALUES 
    ('${GAME_ADMIN_NAME}',  '${GAME_ADMIN_EMAIL}',  'ADMIN',    '76636822',     'GitHub');

INSERT INTO `Accounts` 
    (`Name`,                `Email`,                `Role`)
VALUES 
    ('Tusks',               'tusks@npc.com',        'NPC');


INSERT INTO `AccountProperties` 
    (`AccountId`, `VerificationToken`, `TokenExpiresAt`, `EmailVerified`)
VALUES 
    ((
        SELECT `Id` 
        FROM `Accounts`
        WHERE `Name`='${GAME_ADMIN_NAME}'), UUID(), CURRENT_TIMESTAMP, TRUE),
    ((
        SELECT `Id` 
        FROM `Accounts`
        WHERE `Name`='Tusks'), UUID(), CURRENT_TIMESTAMP, TRUE);