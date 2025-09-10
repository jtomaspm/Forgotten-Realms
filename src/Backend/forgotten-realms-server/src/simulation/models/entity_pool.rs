use std::sync::{Arc, RwLock};

use rayon::iter::{IntoParallelRefMutIterator, ParallelIterator};
use colored::Colorize;

use crate::{game::entities::village::Village, simulation::models::{server_tick::ServerTick, simulation_entity::SimulationEntity}};

pub struct EntityPool {
    pub villages: Vec<Arc<RwLock<Village>>>,
}

impl EntityPool {
    pub fn new() -> Self {
        return EntityPool {
            villages: Vec::new(),
        };
    }

    pub fn update(&mut self, tick: &ServerTick) {
        self.villages.par_iter_mut().for_each(|entity| {
            match entity.try_write() {
                Err(e) => {
                    println!("{} {}", "[ERROR] Failed to lock entity:".bold().red(), e.to_string().bold().red());
                },
                Ok(ref mut e) => {
                    e.update(tick);
                },  
            }
        });
    }
}