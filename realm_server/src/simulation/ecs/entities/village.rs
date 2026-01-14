use crate::simulation::{ecs::models::{coords::Coords, resources::Resources}, tick::Tick};

#[derive(Clone, Copy)]
pub struct Village {
    pub owner_id: usize,
    pub resources: Resources,
    pub resource_production: Resources,
    pub resource_capacity: Resources,
    pub coords: Coords,
}

impl Village {
    pub fn default(owner_id: usize) -> Village {
        return Village { 
            owner_id: owner_id, 
            resources: Resources::new_village_resources(), 
            resource_production: Resources::village_base_production(),
            resource_capacity: Resources::village_base_capacity(),
            coords: Coords::random(),
        };
    }

    pub fn update(&mut self, tick: &Tick) {
    }
}