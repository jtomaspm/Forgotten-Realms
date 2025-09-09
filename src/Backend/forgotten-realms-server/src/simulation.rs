use rand::Rng;
use rayon::iter::{IntoParallelRefMutIterator, ParallelIterator};

use crate::{game::models::village::Village, simulation::models::{server_tick::ServerTick, simulation_entity::SimulationEntity}};

pub mod models;
pub struct Simulation {
    entities: Vec<Box<dyn SimulationEntity>>,
    tick: ServerTick,
}

impl Simulation {
    pub fn new() -> Self {
        return Simulation {
            entities: Vec::new(),
            tick: ServerTick::new(60),
        };
    }

    pub fn run(&mut self) {
        for i in 1..1000*1000 {
            let mut rng = rand::rng();
            self.entities.push(Box::new(Village::new(rng.random_range(-500..501), rng.random_range(-500..501))));
        }

        loop {
            self.entities.par_iter_mut().for_each(|entity| entity.update());

            self.tick.wait_next_tick();
        }    
    }
}