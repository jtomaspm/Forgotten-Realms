CREATE TABLE realms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL,
    api VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE account_realms (
    account_id UUID NOT NULL,
    realm_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, realm_id),
    FOREIGN KEY (realm_id) REFERENCES realms(id) ON DELETE CASCADE
);

CREATE INDEX idx_account_realms_account_id ON account_realms(account_id);
CREATE INDEX idx_account_realms_realm_id ON account_realms(realm_id);

CREATE TRIGGER trigger_set_updated_at_realms
BEFORE UPDATE ON realms
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();