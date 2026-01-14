use rand::Rng;

#[derive(Clone, Copy)]
pub struct Coords {
    pub x: usize,
    pub y: usize,
}

impl Coords {
    pub fn random() -> Coords {
        let mut rng = rand::rng();
        let x = rng.random_range(0..1000);
        let y = rng.random_range(0..1000);
        return Coords { x, y };
    }
}
