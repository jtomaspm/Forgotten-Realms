use serde::Serialize;

#[derive(Serialize)]
pub struct CreatePlayer {
    pub player_id: u32,
}

