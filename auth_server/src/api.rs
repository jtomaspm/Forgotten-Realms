use axum::{Extension, Router};
use sqlx::{Pool, Postgres};

use crate::api::controller::mount_controllers;

pub mod controller;

pub fn setup_api_router(pool: Pool<Postgres>) -> Router {
    mount_controllers(Router::new())
        .layer(Extension(pool))
}