use axum::{Extension, Router};
use sqlx::{Pool, Postgres};
use tokio::sync::mpsc::Sender;

use crate::{api::controller::mount_controllers, event_hub::event::Event};

pub mod controller;
pub mod service;
pub mod models;

pub fn api_router(hub_channel: Sender<Event>, pool: Pool<Postgres>) -> Router {
    mount_controllers(Router::new())
        .layer(Extension(hub_channel))
        .layer(Extension(pool))
}