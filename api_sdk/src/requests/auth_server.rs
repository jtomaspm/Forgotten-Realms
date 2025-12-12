use serde::Deserialize;

#[derive(Deserialize)]
pub struct LoginLocal {
    pub user: String,
    pub password: String,
}