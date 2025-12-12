use std::env;

use api_sdk::configuration::auth_server::AuthServerConfig;
use tokio::fs;

pub async fn get_configuration() -> AuthServerConfig {
    let config_path = format!(
        "{}/.config/forgotten_realms/auth", 
        env::home_dir().expect("Could not determine home directory").to_str().unwrap());

    if fs::create_dir_all(&config_path).await.is_err() {
        panic!("Could not create configuration directory at {}", config_path);
    }

    if fs::try_exists(format!("{}/config.json", config_path)).await.unwrap_or(false) {
        let config_content = fs::read_to_string(format!("{}/config.json", config_path)).await.expect("Could not read configuration file");
        let config: AuthServerConfig = serde_json::from_str(&config_content).expect("Could not parse configuration file");
        return config;
    } else {
        let default_config = AuthServerConfig::default();
        let config_content = serde_json::to_string_pretty(&default_config).expect("Could not serialize default configuration");
        fs::write(format!("{}/config.json", config_path), config_content).await.expect("Could not write default configuration file");
        return default_config;
    }
}