use crate::simulation::models::server_tick::{ServerTick};

pub trait SimulationEntity: Send + Sync {
    fn update(&mut self, server_tick: &ServerTick);
}