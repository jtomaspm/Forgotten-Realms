use crate::{game::map::map_manager::MapManager, simulation::models::{entity_pool::EntityPool, server_tick::ServerTick}};

pub struct SimulationMemory {
    pub map: MapManager,
    pub entity_pool: EntityPool,
}

impl SimulationMemory {
    pub fn new(tick: ServerTick) -> Self {
        return SimulationMemory {
            map: MapManager::new(tick),
            entity_pool: EntityPool::new(),
        };
    }
}