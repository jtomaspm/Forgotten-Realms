pub mod player;
pub mod village;
use tokio::sync::mpsc::Sender;
use uuid::Uuid;

use crate::event_hub::event::{player::CreatePlayerEvent, village::UpgradeBuildingEvent};

pub struct Event {
    pub id: Uuid,
    pub response_channel: Option<Sender<Event>>,
    pub payload: EventPayload,
}

pub enum EventPayload {
    CreatePlayer(CreatePlayerEvent),
    UpgradeBuilding(UpgradeBuildingEvent),
}