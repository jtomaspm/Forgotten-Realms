use axum::{
    routing::get,
    response::Html,
    Router,
};
use tower_http::services::ServeDir;
use askama::Template;

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(index))
        .nest_service("/static", ServeDir::new("static"));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000")
        .await
        .unwrap();

    axum::serve(listener, app).await.unwrap();
}

#[derive(Template)]
#[template(path = "index.html")]
struct IndexTemplate;

async fn index() -> impl axum::response::IntoResponse {
    let tpl = IndexTemplate;
    Html(tpl.render().unwrap())
}