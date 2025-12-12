use axum::{Json, http::StatusCode, response::IntoResponse};
use serde::Serialize;


#[derive(Serialize)]
pub struct ErrorResponse {
    pub error: String, 
}
impl IntoResponse for ErrorResponse {
    fn into_response(self) -> axum::response::Response {
        let status = StatusCode::BAD_REQUEST;
        (status, Json(self)).into_response()
    }
}