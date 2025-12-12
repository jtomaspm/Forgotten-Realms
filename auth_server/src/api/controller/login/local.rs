use api_sdk::{requests, responses::{self, auth_server::LoginLocal, error::ErrorResponse}};
use axum::{Extension, Json, Router, routing::post};
use sqlx::{Pool, Postgres};


pub fn mount(router: Router) -> Router {
    return router.route("/login/local", post(login));
}

async fn login(pool: Extension<Pool<Postgres>>, Json(request): Json<requests::auth_server::LoginLocal>) -> Result<Json<responses::auth_server::LoginLocal>, ErrorResponse> {

    return Ok(Json(LoginLocal{
        token: "hello, world".to_string()
    }));
}