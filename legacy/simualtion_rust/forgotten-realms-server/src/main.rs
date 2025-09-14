use crate::simulation::Simulation;

mod simulation;
mod game;

fn main() {
    let mut sim = Simulation::new();
    sim.run();
}
