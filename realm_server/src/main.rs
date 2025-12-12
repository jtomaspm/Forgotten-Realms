use crate::server::Server;

mod simulation;
mod event_hub;
mod server;
mod api;

#[tokio::main]
async fn main() {
    let server = Server::new(None).await;
    server.run().await;
}
