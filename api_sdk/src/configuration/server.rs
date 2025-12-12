use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct ServerConfig {
    pub name: String,
    pub host: String,
    pub port: u16,
}