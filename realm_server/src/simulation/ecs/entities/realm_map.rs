use crate::simulation::ecs::models::{coords::Coords, map_slot::MapSlot};

pub struct RealmMap {
    pub slots: Vec<Vec<MapSlot>>,
    pub villages: Vec<Coords>,
}

impl RealmMap {
    pub fn default() -> RealmMap {
        let mut slots = Vec::with_capacity(1000);
        for y in 0..1000 {
            let mut s_row = Vec::with_capacity(1000);
            for x in 0..1000{
                s_row.push(MapSlot::default(x, y));
            }
            slots.push(s_row);
        } 
        
        return RealmMap { slots , villages: vec![] };
    }

    pub fn slot(&self, coords: Coords) -> &MapSlot {
        return &self.slots[coords.y][coords.x];
    }
    pub fn slot_mut(&mut self, coords: Coords) -> &mut MapSlot {
        return &mut self.slots[coords.y][coords.x];
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_map() {
        let map = RealmMap::default();
        for y in 0..5 {
            for x in 0..5 {
                assert_eq!(x, map.slots[y][x].coords.x);
                assert_eq!(y, map.slots[y][x].coords.y);
            }
        }
    }

}