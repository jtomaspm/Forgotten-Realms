use api_sdk::configuration::{auth_server::AuthServerConfig};
use sqlx::{Pool, Postgres, postgres::PgPoolOptions};

pub async fn setup_database_pool(config: &AuthServerConfig) -> Pool<Postgres> {
    match PgPoolOptions::new()
        .max_connections(config.database.max_connections)
        .connect(
            &format!(
                "postgres://{}:{}@{}:{}/{}", 
                config.database.user.to_string(), 
                config.database.password.to_string(), 
                config.database.host.to_string(), 
                config.database.port.to_string(), 
                config.database.database.to_string()))
        .await 
    {
        Ok(pool) => pool,
        Err(e) => panic!("failed to connect to database: {}", e.to_string()),
    }
}