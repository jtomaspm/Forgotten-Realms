use clap::Parser;

use crate::args::{Args, Commands};

mod args;

#[tokio::main]
async fn main() {
    let cli = Args::parse();

    match cli.command {
        Commands::GenerateConfig { path, server_type } => {
            match server_type {
                args::ServerType::Realm => todo!(),
                args::ServerType::Auth => todo!(),
            }
        }
        Commands::Migrate { migrations_folder, connections, users } => {
            let result = db_migration_lib::migrate(
                migrations_folder, 
                connections.iter().map(|connection| db_migration_lib::models::MigrationConfig{name: connection.name.clone(), connection_string: connection.connection_string.clone()}).collect(),
                users.iter().map(|user| db_migration_lib::models::DbUser{name: user.name.clone(), password: user.password.clone()}).collect())
                .await;
            match result {
                Ok(_) => {
                    println!("Migrations applied successfully.");
                },
                Err(err) => {
                    eprintln!("Migration failed: {}", err);
                },
            };
        }
    }
}
