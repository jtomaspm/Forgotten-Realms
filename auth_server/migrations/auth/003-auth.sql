CREATE TABLE auth.sessions (
    account_id INT PRIMARY KEY SERIAL,
    token UUID UNIQUE DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP + INTERVAL '7 days',
    FOREIGN KEY (account_id) REFERENCES auth.accounts(id) ON DELETE CASCADE
);

CREATE TABLE auth.logins (
    id INT PRIMARY SERIAL,
    account_id INT NOT NULL, 
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(45),
    FOREIGN KEY (account_id) REFERENCES auth.accounts(id) ON DELETE CASCADE
);

CREATE INDEX idx_sessions_token ON auth.sessions(token);