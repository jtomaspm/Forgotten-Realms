use crate::simulation::{ecs::models::Resources, tick::Tick};

pub struct Entity<T>{
    pub id: u32,
    pub deleted: bool,
    pub last_updated: Tick,
    pub data: T,
}

pub struct Player {
     pub faction_id: u8,
     pub villages: Vec<u32>,
}

pub struct Village {
    pub owner_id: u32,
    pub resources: Resources,
}