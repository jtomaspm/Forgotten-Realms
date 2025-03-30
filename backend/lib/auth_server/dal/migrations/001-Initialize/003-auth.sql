CREATE TABLE sessions (
    account_id PRIMARY KEY UUID NOT NULL,  
    token UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP + INTERVAL '7 days',
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
);

CREATE TABLE logins (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
    account_id UUID NOT NULL, 
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(45),
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
);