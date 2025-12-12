DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        CREATE TYPE role AS ENUM ('admin', 'moderator', 'npc', 'player', 'guest', 'sitter');
    END IF;
END $$;

CREATE TABLE auth.accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id VARCHAR(255) NOT NULL,
    source VARCHAR(50) NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    role role NOT NULL,  
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE auth.account_properties (
    account_id UUID PRIMARY KEY,
    email_verified BOOLEAN DEFAULT FALSE,
    send_email_notifications BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES auth.accounts(id) ON DELETE CASCADE
);

CREATE TABLE auth.account_verification (
    account_id UUID PRIMARY KEY,
    verification_token UUID DEFAULT gen_random_uuid(),
    token_expires_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP + INTERVAL '7 days',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES auth.accounts(id) ON DELETE CASCADE
);

CREATE TRIGGER auth.trigger_set_updated_at_account_properties
BEFORE UPDATE ON auth.account_properties
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER auth.trigger_set_updated_at_accounts
BEFORE UPDATE ON auth.accounts
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();