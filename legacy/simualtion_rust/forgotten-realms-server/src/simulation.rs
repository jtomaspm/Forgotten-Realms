use crate::{game::models::direction::Direction, simulation::models::{server_tick::ServerTick, simulation_memory::SimulationMemory}};

pub mod models;
pub struct Simulation {
    memory: SimulationMemory,
    tick: ServerTick,
}

impl Simulation {
    pub fn new() -> Self {
        let tick = ServerTick::new(100);
        return Simulation {
            memory: SimulationMemory::new(tick),       
            tick,
        };
    }

    pub fn run(&mut self) {
        self.memory.map.fill_map_with_villages(&mut self.memory.entity_pool.villages);
        //for i in 1..=1 {
            //self.memory.map.spawn_village(Direction::Random, &mut self.memory.entity_pool.villages);
        //}
        
        loop {
            self.memory.entity_pool.update(&self.tick);           

            self.tick.wait_next_tick();
        }    
    }
}