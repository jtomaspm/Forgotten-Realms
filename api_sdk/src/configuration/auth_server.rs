use serde::{Deserialize, Serialize};

use crate::configuration::{database::{DatabaseConfig, MigrationConfig}, server::ServerConfig};

#[derive(Serialize, Deserialize)]
pub struct AuthServerConfig {
    pub server: ServerConfig,
    pub database: DatabaseConfig,
    pub migration: MigrationConfig,
}

impl AuthServerConfig {
    pub fn default() -> Self {
        AuthServerConfig { 
            server: ServerConfig {
                name: "auth_server".to_string(),
                host: "0.0.0.0".to_string(),
                port: 8080,
            },
            database: DatabaseConfig {
                user: "as_auth".to_string(),
                password: "test1234".to_string(),
                database: "auth".to_string(),
                host: "localhost".to_string(),
                port: 5432,
                max_connections: 10,
            },
            migration: MigrationConfig {
                user: "postgres".to_string(),
                password: "test1234".to_string(),
            },
        }
    }
}
