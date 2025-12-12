use crate::{migrator::Migrator, models::{DbUser, MigrationConfig}};

pub mod models;
pub mod migrator;

pub async fn migrate(migrations_folder: String, connections: Vec<MigrationConfig>, users: Vec<DbUser>) -> Result<(), String> {
    let migrator_result = Migrator::new(migrations_folder);
    if migrator_result.is_err() {
        return Err(migrator_result.err().unwrap());
    }
    let mut migrator = migrator_result.unwrap();
    migrator.configure_migrations(connections);
    migrator.configure_users(users);
    return migrator.run_migrations().await;
}


#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn test_migrate() {
        let migrations_folder = "/home/pop/Code/Forgotten-Realms/db_migration_lib/migrations".to_string();
        let users = vec![DbUser{name: "app_user".to_string(), password: "test1234".to_string()}];
        let connections = vec![MigrationConfig{connection_string: "postgres://postgres:test1234@localhost:5432/test_migrations".to_string(), name: "test".to_string()}];

        let result = migrate(migrations_folder, connections, users).await;
        match result.clone() {
            Ok(_) => println!("Migrations ran successfully"),
            Err(e) => println!("Migration error: {}", e),
        }
        assert!(result.is_ok());
    }
}