mod configuration;
mod database;

use api_sdk::configuration::{auth_server::AuthServerConfig};
use sqlx::{Pool, Postgres};
use tokio::{net::TcpListener};

use crate::{api::setup_api_router, server::{configuration::get_configuration, database::setup_database_pool}};


pub struct Server {
    config: AuthServerConfig,
    pool: Pool<Postgres>,
}

impl Server {
    pub async fn new() -> Self {
        let config = get_configuration().await;
        let pool = setup_database_pool(&config).await;
        Server {
            config: config,
            pool: pool,
        }
    }

    pub async fn run(&self) {
        println!("Server starting...");

        let listener = TcpListener::bind("0.0.0.0:3000").await.expect("Failed to bind tcp socket");
        axum::serve(listener, setup_api_router(self.pool.clone())).await.expect("Failed to serve api on binded tcp socket");
    }
}