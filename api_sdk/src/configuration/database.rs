use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct DatabaseConfig {
    pub user: String,
    pub password: String,
    pub database: String,
    pub host: String,
    pub port: u16,
    pub max_connections: u32,
}
#[derive(Serialize, Deserialize)]
pub struct MigrationConfig {
    pub user: String,
    pub password: String,
}