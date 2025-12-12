mod configuration;
mod database;

use api_sdk::configuration::{auth_server::AuthServerConfig};
use axum::{Extension, Router};
use sqlx::{Pool, Postgres};
use tokio::{net::TcpListener};

use crate::{api::controller::mount_controllers, server::{configuration::get_configuration, database::setup_database_pool}};


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

        let mut api = Router::new();
        api = mount_controllers(api)
            .layer(Extension(self.pool.clone()));
        let api = api;

        let listener = TcpListener::bind("0.0.0.0:3000").await.unwrap();
        axum::serve(listener, api).await.unwrap();
    }
}