use serde::Deserialize;

#[derive(Deserialize)]
pub struct CreatePlayer {
    pub name: String,
    pub faction: String,
}

