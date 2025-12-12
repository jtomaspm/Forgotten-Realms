use tokio::sync::mpsc;

use crate::event_hub::event::Event;

pub struct SimulationIO {
    pub input_channel: mpsc::Receiver<Event>,
    pub output_channel: mpsc::Sender<Event>,
}

impl SimulationIO {
    pub fn new(input_channel: mpsc::Receiver<Event>, output_channel: mpsc::Sender<Event>) -> Self {
        SimulationIO {
            input_channel,
            output_channel,
        }
    }
}