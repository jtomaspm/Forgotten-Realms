pub trait SimulationEntity: Send + Sync {
    fn update(&mut self);
}