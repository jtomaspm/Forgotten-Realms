use crate::simulation::{tick::Tick};

pub mod village;
pub mod player;
pub mod realm_map;

pub struct Entity<T>{
    pub id: u32,
    pub deleted: bool,
    pub last_updated: Tick,
    pub data: T,
}



