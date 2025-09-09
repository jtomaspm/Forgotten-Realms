use std::{thread, time::{Duration, Instant}};

pub struct ServerTick {
    current_tick: Instant,
    wait_time: Duration,
}

impl ServerTick {
    pub fn new(ticks_per_second: u32) -> Self {
        ServerTick {
            current_tick: Instant::now(),
            wait_time: Duration::from_nanos(1000000000 / ticks_per_second as u64),
        }
    }

    pub fn wait_next_tick(&mut self) {
        let sleep_time = self.wait_time.checked_sub(self.current_tick.elapsed());
        match sleep_time {
            Some(duration) => {
                print!("Time left to use this frame: {:?}\n", duration);
                thread::sleep(duration);
            },
            None => { /* We're behind schedule, skip sleeping */ }
        }
        self.current_tick = Instant::now();
    }
}