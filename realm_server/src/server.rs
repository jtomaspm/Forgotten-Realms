mod configuration;
mod database;

use api_sdk::configuration::realm_server::RealmServerConfig;
use sqlx::{Pool, Postgres};
use tokio::{net::TcpListener, sync::mpsc};

use crate::{api::{api_router}, event_hub::{EventHub, event::Event, simulation::SimulationIO}, server::{configuration::get_configuration, database::setup_database_pool}, simulation::Simulation};


pub struct Server {
    config: RealmServerConfig,
    pool: Pool<Postgres>,
}

impl Server {
    pub async fn new(config_path: Option<String>) -> Self {
        let config = get_configuration(config_path).await;
        let pool = setup_database_pool(&config).await;
        Server {
            config: config,
            pool: pool,
        }
    }

    pub async fn run(&self) {
        println!("Server starting...");

        let (sender_simulation, reciever_simulation) = mpsc::channel::<Event>(32);
        let (sender_hub, reciever_hub) = mpsc::channel::<Event>(32);
        let api_hub_channel = sender_hub.clone();
        let hub_pool =  self.pool.clone();

        tokio::task::spawn(async move {
            EventHub::new(reciever_hub, sender_simulation, hub_pool).run();
        });

        tokio::task::spawn_blocking(move || {
            Simulation::new(SimulationIO::new(reciever_simulation, sender_hub.clone())).run();
        });

        let listener = TcpListener::bind("0.0.0.0:3000").await.unwrap();
        axum::serve(listener, api_router(api_hub_channel, self.pool.clone())).await.unwrap();
    }
}