use std::time::Duration;

use tokio::time::Instant;

#[derive(PartialEq, Eq, Clone, Copy)]
pub struct Tick {
    tick_time: Instant
}

impl Tick {
    pub fn new() -> Self {
        Tick {
            tick_time: Instant::now()
        }
    }

    pub fn tick(&mut self, fps: u64) {
        let elapsed = self.tick_time.elapsed();
        let frame_duration = Duration::from_millis(1000 / fps);

        if elapsed < frame_duration {
            std::thread::sleep(frame_duration - elapsed);
        } else {
            println!("Simulation tick took longer than frame duration!");
        }
        self.tick_time = Instant::now();
    }

}