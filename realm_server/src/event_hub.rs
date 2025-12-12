
use sqlx::{Pool, Postgres};
use tokio::sync::mpsc::{Receiver, Sender};

use crate::event_hub::event::{Event, EventPayload};

pub mod simulation;
pub mod event;

pub struct EventHub {
    input_channel: Receiver<Event>,
    simulation_channel: Sender<Event>,
    pool: Pool<Postgres>,
}

impl EventHub {
    pub fn new(input_channel: Receiver<Event>, simulation_channel: Sender<Event>, pool: Pool<Postgres>) -> Self {
        EventHub {
            input_channel,
            simulation_channel,
            pool,
        }
    }

    pub fn run(&mut self) {
        println!("EventHub running...");
        loop {
            match self.input_channel.try_recv() {
                Ok(event) => {
                    self.handle_event(event);
                }
                Err(err) => {
                    match err {
                        tokio::sync::mpsc::error::TryRecvError::Disconnected => {
                            println!("Input channel disconnected, shutting down EventHub.");
                            break;
                        },
                        tokio::sync::mpsc::error::TryRecvError::Empty => {
                            // No event available, continue
                        }
                    }
                }
            }
        }
    }

    fn handle_event(&self, event: Event) {
        println!("Received event with id: {}", event.id);
        match event.payload {
            EventPayload::CreatePlayer(payload) => {
                println!("Handling CreatePlayer event {}", payload.name);
                self.simulation_channel.try_send(Event { id: event.id, response_channel: None, payload: EventPayload::CreatePlayer(payload)}).unwrap();
            },
            EventPayload::UpgradeBuilding(payload) => {
                println!("Handling UpgradeBuilding event {}", payload.village_id);
            },
        }
    }
}