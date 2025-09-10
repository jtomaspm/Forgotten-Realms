use std::cmp::Ordering;
use std::time::SystemTime;

pub struct Event {
    pub name: String,
    pub execute_date: SystemTime,
}

impl PartialEq for Event {
    fn eq(&self, other: &Self) -> bool {
        self.execute_date.eq(&other.execute_date)
    }
}

impl Eq for Event {}

impl PartialOrd for Event {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.execute_date.cmp(&other.execute_date))
    }
}

impl Ord for Event {
    fn cmp(&self, other: &Self) -> Ordering {
        self.execute_date.cmp(&other.execute_date)
    }
}

pub struct EventPool {
    pub events: Vec<Event>,
}

impl EventPool {
    pub fn new() -> Self {
        return EventPool {
            events: Vec::new(),
        };
    }

    fn add_event(&mut self, new_event: Event) {
        let pos = self.events.binary_search(&new_event).unwrap_or_else(|e| e);
        self.events.insert(pos, new_event);
    }

    pub fn take_ready_events(&mut self, before: SystemTime) -> Vec<Event> {
        let cutoff = self
            .events
            .partition_point(|e| e.execute_date <= before);
        self.events.drain(..cutoff).collect()
    } 
}