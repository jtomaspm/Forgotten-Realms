use crate::{event_hub::event::player::CreatePlayerEvent, simulation::{ecs::{ECS, entities::{Entity, Player}}, tick::Tick}};

pub fn create_player (ecs: &mut ECS, event: CreatePlayerEvent, tick: Tick) {
    ecs.players.push(Entity::<Player> {
        id: ecs.players.len() as u32,
        deleted: false,
        last_updated: tick,
        data: Player {
            faction_id: event.faction_id,
            villages: Vec::new(),
        }
    });
}