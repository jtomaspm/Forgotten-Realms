use tokio::sync::mpsc::error::TryRecvError;

use crate::{event_hub::{event::EventPayload, simulation::SimulationIO}, simulation::{ecs::ECS, tick::Tick}};

pub mod ecs;
pub mod tick;
mod event_handlers;

pub struct Simulation {
    ecs: ECS,
    io: SimulationIO,
}

impl Simulation {
    pub fn new(io: SimulationIO) -> Self {
        Simulation {
            ecs: ECS::new(),
            io: io,
        }
    }

    pub fn run(&mut self) {
        println!("Simulation running...");

        let mut tick = Tick::new();
        loop {
            //ecs update
            self.ecs.update();

            //process incoming events
            let result = self.process_events(100, tick);
            if result.is_err() {
                return;
            }

            //tick
            tick.tick(100);
        }
    }

    fn process_events(&mut self, max_amount: usize, tick: Tick) -> Result<(), TryRecvError> {
        for _ in 0..max_amount
        {
            match self.io.input_channel.try_recv() {
                Ok(event) => {
                    println!("Simulation received event with id: {}", event.id);
                    match event.payload {
                        EventPayload::CreatePlayer(payload) => {
                            println!("Simulation handling CreatePlayer event {}", payload.name);
                            event_handlers::player::create_player(&mut self.ecs, payload, tick);
                        },
                        _ => {
                            println!("Simulation received unhandled event payload");
                        }
                    }
                }
                Err(err) => {
                    match err {
                        TryRecvError::Disconnected => {
                            println!("Input channel disconnected, shutting down Simulation.");
                            return Err(err);
                        }
                        TryRecvError::Empty => {
                            return Ok(());
                        }
                    }
                }
            }
        }
        return Ok(());
    }
}