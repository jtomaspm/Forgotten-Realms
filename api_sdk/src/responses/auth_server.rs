use serde::Serialize;

#[derive(Serialize)]
pub struct LoginLocal {
    pub token: String,
}