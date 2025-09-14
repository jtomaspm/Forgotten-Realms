use std::{thread, time::{Duration, Instant, SystemTime}};
use colored::Colorize;

#[derive(Clone, Copy)]
pub struct ServerTick {
    pub current_tick: SystemTime,
    pub wait_time: Duration,
    pub delta_secs: f32,
}

impl ServerTick {
    pub fn new(ticks_per_second: u32) -> Self {
        let wait_time = Duration::from_secs_f32(1.0 / ticks_per_second as f32);
        ServerTick {
            current_tick: SystemTime::now(),
            wait_time: wait_time,
            delta_secs: wait_time.as_secs_f32(),
        }
    }

    pub fn wait_next_tick(&mut self) {
        let sleep_time = self.wait_time.checked_sub(self.current_tick.elapsed().unwrap_or(Duration::from_secs(0)));
        match sleep_time {
            Some(duration) => {
                print!("Free time: {:?}\n", duration);
                thread::sleep(duration);
            },
            None => { println!("{}", "[WARNING] Frame took longer than tick duration!!!".bold().red()) }
        }
        self.current_tick = SystemTime::now();
    }
}