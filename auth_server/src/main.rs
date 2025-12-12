use crate::server::Server;

mod server;
mod api;

#[tokio::main]
async fn main() {
    let server = Server::new().await;
    server.run().await;
}
