use crate::simulation::models::simulation_entity::SimulationEntity;

pub struct Village {
    coord_x: i32,
    coord_y: i32,
    wood: i32,
    clay: i32,
    iron: i32,
    wood_per_hour: i32,
    clay_per_hour: i32,
    iron_per_hour: i32,
}

impl Village {
    pub fn new(x: i32, y: i32) -> Self {
        Village {
            coord_x: x,
            coord_y: y,
            wood: 0,
            clay: 0,
            iron: 0,
            wood_per_hour: 100,
            clay_per_hour: 100,
            iron_per_hour: 100,
        }
    }
}

impl SimulationEntity for Village {
    fn update(&mut self) {
        self.wood += self.wood_per_hour / 3600;
        self.clay += self.clay_per_hour / 3600;
        self.iron += self.iron_per_hour / 3600;
    }
}