pub mod login;
pub mod refresh_token;
pub mod authenticate;
pub mod profile;

use axum::Router;

use crate::api::controller;

pub fn mount_controllers(router: Router) -> Router {
    let router = controller::login::local::mount(router);
    return router;
}
