use crate::simulation::{ecs::entities::{Entity, player::Player, realm_map::RealmMap, village::Village}, tick::Tick};

pub mod entities;
pub mod models;

pub struct ECS {
    pub players: Vec<Entity<Player>>,
    pub villages: Vec<Entity<Village>>,
    pub map: RealmMap,
}

impl ECS {
    pub fn new() -> Self {
        ECS {
            players: Vec::new(),
            villages: Vec::new(),
            map: RealmMap::default(),
        }
    }

    pub fn update(&mut self, tick: &Tick) {

    }

    
}