use crate::simulation::ecs::entities::{Entity, Player, Village};

pub mod entities;
pub mod models;

pub struct ECS {
    pub players: Vec<Entity<Player>>,
    pub villages: Vec<Entity<Village>>,
}

impl ECS {
    pub fn new() -> Self {
        ECS {
            players: Vec::new(),
            villages: Vec::new(),
        }
    }

    pub fn update(&mut self) {
        
    }

    
}