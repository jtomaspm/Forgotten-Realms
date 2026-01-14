#[derive(Clone, Copy)]
pub struct Resources {
    pub wood: u64,
    pub stone: u64,
    pub iron: u64,
}

impl Resources {
    pub fn empty() -> Resources {
        return Resources { wood: 0, stone: 0, iron: 0 };
    }
    pub fn new_village_resources() -> Resources {
        return Resources { wood: 750, stone: 750, iron: 750 };
    }
    pub fn village_base_production() -> Resources {
        return Resources { wood: 500, stone: 500, iron: 500 };
    }
    pub fn village_base_capacity() -> Resources {
        return Resources { wood: 1000, stone: 1000, iron: 1000 };
    }
}
