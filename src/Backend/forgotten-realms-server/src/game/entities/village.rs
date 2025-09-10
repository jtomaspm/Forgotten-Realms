use crate::{game::models::coordinates::Coordinates, simulation::models::{server_tick::{ServerTick}, simulation_entity::SimulationEntity}};

pub struct Village {
    pub coordinates: Coordinates,
    pub wood: f32,
    pub clay: f32,
    pub iron: f32,
    pub base_resource_production_per_hour: f32,
    pub wood_per_tick: f32,
    pub clay_per_tick: f32,
    pub iron_per_tick: f32,
    pub wood_per_hour_from_buildings: f32,
    pub clay_per_hour_from_buildings: f32,
    pub iron_per_hour_from_buildings: f32,
    pub wood_production_multiplier: f32,
    pub clay_production_multiplier: f32,
    pub iron_production_multiplier: f32,
    pub population: u32,
    pub max_population: u32,
}

impl Village {
    pub fn new(coordinates: Coordinates, tick: &ServerTick) -> Self {
        let mut village = Village {
            coordinates,
            wood: 0.0,
            clay: 0.0,
            iron: 0.0,
            base_resource_production_per_hour: 3600.0,
            wood_per_tick: 0.0,
            clay_per_tick: 0.0,
            iron_per_tick: 0.0,
            wood_per_hour_from_buildings: 0.0,
            clay_per_hour_from_buildings: 0.0,
            iron_per_hour_from_buildings: 0.0,
            wood_production_multiplier: 1.0,
            clay_production_multiplier: 1.0,
            iron_production_multiplier: 1.0,
            population: 1,
            max_population: 100,
        };
        village.update_production_per_tick(tick);
        return village;
    }
    pub fn update_production_per_tick(&mut self, tick: &ServerTick) {
        self.wood_per_tick = ((self.wood_per_hour_from_buildings+self.base_resource_production_per_hour) * self.wood_production_multiplier) / 3600.0 * tick.delta_secs;
        self.clay_per_tick = ((self.clay_per_hour_from_buildings+self.base_resource_production_per_hour) * self.clay_production_multiplier) / 3600.0 * tick.delta_secs;
        self.iron_per_tick = ((self.iron_per_hour_from_buildings+self.base_resource_production_per_hour) * self.iron_production_multiplier) / 3600.0 * tick.delta_secs;
    }
}

impl SimulationEntity for Village {
    fn update(&mut self, tick: &ServerTick) {
        self.wood += self.wood_per_tick;
        self.clay += self.clay_per_tick;
        self.iron += self.iron_per_tick;
        //print!("Village at ({}, {}) has {:.0} wood, {:.0} clay, {:.0} iron\n", self.coordinates.x, self.coordinates.y, self.wood, self.clay, self.iron);
    }
}