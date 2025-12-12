use axum::Router;

pub mod player;

pub fn mount_controllers(router: Router) -> Router {
    let router = player::mount(router);
    return router;
}