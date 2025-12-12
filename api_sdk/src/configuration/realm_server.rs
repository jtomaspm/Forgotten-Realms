use serde::{Deserialize, Serialize};

use crate::configuration::{database::{DatabaseConfig, MigrationConfig}, server::ServerConfig};

#[derive(Serialize, Deserialize)]
pub struct RealmServerConfig {
    pub server: ServerConfig,
    pub auth_server: String,
    pub database: DatabaseConfig,
    pub migration: MigrationConfig,
}

impl RealmServerConfig {
    pub fn default() -> Self {
        RealmServerConfig { 
            server: ServerConfig {
                name: "realm_dev".to_string(),
                host: "0.0.0.0".to_string(),
                port: 8090,
            },
            auth_server: "http://localhost:8080".to_string(),
            database: DatabaseConfig {
                user: "rs_realm_dev".to_string(),
                password: "test1234".to_string(),
                database: "realm_dev".to_string(),
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