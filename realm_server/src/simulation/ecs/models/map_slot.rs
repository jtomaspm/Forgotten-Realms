use std::fmt::Display;

use crate::simulation::ecs::{models::coords::Coords};

#[derive(Clone, Copy)]
pub struct MapSlot {
    pub coords: Coords,
    pub village: Option<usize>,
}

impl Display for MapSlot {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        return f.write_str(format!("MapSlot -> coords:[{}|{}], village:{} ", self.coords.x, self.coords.y, self.village.is_some()).as_str());
    }
}

impl MapSlot {
    pub fn default(x: usize, y: usize) -> MapSlot {
        return MapSlot { coords: Coords { x, y }, village: None };
    }
}