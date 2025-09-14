use std::{collections::HashMap, sync::{Arc, RwLock}};

use rand::Rng;

use crate::{game::{entities::village::Village, models::{coordinates::Coordinates, direction::Direction}}, simulation::models::server_tick::ServerTick};

pub struct MapManager {
    villages: HashMap<Coordinates, Arc<RwLock<Village>>>,
    tick: ServerTick
}

impl MapManager {
    pub fn new(tick: ServerTick) -> Self {
        return MapManager {
            villages: HashMap::new(),
            tick,
        };
    }

    fn spawn_village_inner(&mut self, direction: Direction, entities: &mut Vec<Arc<RwLock<Village>>>, tries: i32) -> Result<Arc<RwLock<Village>>, String> {
        if tries <= 0 {
            return Err("Failed to spawn village: too many tries".to_string());
        }

        let mut rng = rand::rng();
        let coords = match direction {
            Direction::NorthEast => Coordinates::new(1, 1),
            Direction::NorthWest => Coordinates::new(-1, 1),
            Direction::SouthEast => Coordinates::new(1, -1),
            Direction::SouthWest => Coordinates::new(-1, -1),
            Direction::Random => {
                match rng.random_range(0..4) {
                    0 => return self.spawn_village_inner(Direction::NorthEast, entities, tries),
                    1 => return self.spawn_village_inner(Direction::NorthWest, entities, tries),
                    2 => return self.spawn_village_inner(Direction::SouthEast, entities, tries),
                    3 => return self.spawn_village_inner(Direction::SouthWest, entities, tries),
                    _ => unreachable!(),
                }
            },
        };

        let coords = Coordinates::new(
            coords.x * rng.random_range(0..=500),
            coords.y * rng.random_range(0..=500),
        );

        if self.villages.contains_key(&coords) {
            return self.spawn_village_inner(direction, entities, tries - 1);
        }

        let arc = Arc::new(RwLock::new(Village::new(coords, &self.tick)));
        entities.push(arc.clone());
        self.villages.insert(coords, arc.clone());
        return Ok(arc);
    }

    pub fn spawn_village(&mut self, direction: Direction, entities: &mut Vec<Arc<RwLock<Village>>>) -> Result<Arc<RwLock<Village>>, String> {
        return self.spawn_village_inner(direction, entities, 100);
    }

    pub fn fill_map_with_villages(&mut self, entities: &mut Vec<Arc<RwLock<Village>>>) {
        for x in -500..=500 {
            for y in -500..=500 {
                let coords = Coordinates::new(x, y);
                if !self.villages.contains_key(&coords) {
                    let arc = Arc::new(RwLock::new(Village::new(coords, &self.tick)));
                    entities.push(arc.clone());
                    self.villages.insert(coords, arc);
                }
            }
        }       
    }

    pub fn get_village(&self, coords: Coordinates) -> Option<Arc<RwLock<Village>>> {
        return self.villages.get(&coords).cloned();
    }
}