use std::{collections::HashMap, fs};

use sqlx::{Pool, Postgres, postgres::PgPoolOptions, Row};

use crate::models::{DbUser, Migration, MigrationConfig};


pub struct Migrator {
    migrations: HashMap<String, Migration>,
    users: Vec<DbUser>,
}

impl Migrator {
    pub fn new(base_path: String) -> Result<Self, String> {
        match fs::read_dir(base_path) {
            Ok(entries) => {
                let mut migrations: HashMap<String, Migration> = HashMap::new();

                for entry in entries {
                    if let Ok(entry) = entry && entry.path().is_dir() {
                        let path = entry.path();
                        let files_result = fs::read_dir(path.clone());
                        let mut sql_files: Vec<String> = Vec::new();
                        if let Ok(files) = files_result {
                            for file in files {
                                if let Ok(file) = file {
                                    let file_path = file.path();
                                    if file_path.is_file() {
                                        if let Some(file_str) = file_path.to_str() {
                                            sql_files.push(file_str.to_string());
                                        }
                                    }
                                }
                            }
                        }
                        let migration = Migration {
                            sql_files: sql_files,
                            connection_string: None,
                        };
                        let name = entry.file_name();
                        migrations.insert(name.into_string().unwrap(), migration);
                    }
                }

                return Ok(Migrator { 
                    migrations: migrations,
                    users: Vec::new(),
                });
            }
            Err(e) => {
                return Err(format!("Failed to read migrations directory: {}", e));
            }
        }
    }

    pub fn configure_migrations(&mut self, configs: Vec<MigrationConfig>) {
        for config in configs {
            if !self.migrations.contains_key(config.name.as_str()) {
                continue;
            }
            let migration_request = self.migrations.get_mut(config.name.as_str());
            match migration_request {
                Some(migration) => {
                    migration.connection_string = Some(config.connection_string);
                }
                None => continue,
            }
        }
    }

    pub fn configure_users(&mut self, configs: Vec<DbUser>) {
        for config in configs {
            if let Some(user) = self.users.iter_mut().find(|u| u.name == config.name) {
                user.password = config.password;
            }
            else {
                self.users.push(config);
            }
        }
    }

    pub async fn run_migrations(&self) -> Result<(), String> {
        for (name, migration) in &self.migrations {
            match &migration.connection_string {
                Some(connection_string) => {
                    // extract database name from connection string and create database if it doesn't exist
                    let db_name_start = connection_string.rfind('/').unwrap() + 1;
                    let db_name_end = connection_string[db_name_start..].find('?').map_or(connection_string.len(), |e| db_name_start + e);
                    let db_name = &connection_string[db_name_start..db_name_end];
                    let admin_connection_string = format!("{}?dbname=postgres", &connection_string[..db_name_start - 1]);
                    let admin_pool = PgPoolOptions::new()
                        .max_connections(1)
                        .connect(&admin_connection_string)
                        .await;
                    if admin_pool.is_err() {
                        return Err(format!("Failed to connect to admin database: {}", admin_pool.err().unwrap()));
                    }
                    let admin_pool = admin_pool.unwrap();
                    let db_exists_result = sqlx::query("SELECT 1 FROM pg_database WHERE datname = $1")
                        .bind(db_name)
                        .fetch_one(&admin_pool)
                        .await;
                    if db_exists_result.is_err() {
                        let create_db_result = sqlx::query(&format!("CREATE DATABASE {}", db_name))
                            .execute(&admin_pool)
                            .await;
                        if create_db_result.is_err() {
                            return Err(format!("Failed to create database {}: {}", db_name, create_db_result.err().unwrap()));
                        }
                    }
                    let pool = PgPoolOptions::new()
                        .max_connections(1)
                        .connect(connection_string)
                        .await;
                    if pool.is_err() {
                        return Err(format!("Failed to connect to database for migration {}: {}", name, pool.err().unwrap()));
                    }
                    let pool = pool.unwrap();

                    for user in &self.users {
                        let result = self.create_db_user(&pool, user).await;
                        if result.is_err() {
                            return Err(format!("Failed to create DB user {}: {}", user.name, result.err().unwrap()));
                        }
                    }

                    self.create_migrations_table(&pool).await;
                    for sql_file in &migration.sql_files {
                        let migration_id = format!("{}/{}", name, sql_file.split(&format!("/{}/", name).to_string()).last().unwrap());
                        if self.has_migration_been_applied(&pool, &migration_id).await {
                            continue;
                        }
                        let sql_content = match fs::read_to_string(sql_file) {
                            Ok(content) => content,
                            Err(e) => { return Err(format!("Failed to read SQL file {}: {}", sql_file, e)); }
                        };
                        let result = sqlx::raw_sql(&sql_content)
                            .execute(&pool)
                            .await;
                        match result {
                            Ok(_) => {
                                let _ = sqlx::query("INSERT INTO migrations.database_migrations (id) VALUES ($1)")
                                    .bind(&migration_id)
                                    .execute(&pool)
                                    .await;
                            },
                            Err(e) => {
                                return Err(format!("Failed to execute migration {}: {}", migration_id, e));
                            }
                        };
                    }
                },
                None => {}
            };
        }
        return Ok(());
    }

    pub async fn create_migrations_table(&self, pool: &Pool<Postgres>) {
        let _ = sqlx::query(
            "CREATE SCHEMA IF NOT EXISTS migrations",
        )
        .execute(pool)
        .await;
        let _ = sqlx::query(
            "CREATE TABLE IF NOT EXISTS migrations.\"database_migrations\" (
                \"id\" VARCHAR(255) PRIMARY KEY,
                \"created_at\" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
            )",
        )
        .execute(pool)
        .await;
    }

    pub async fn has_migration_been_applied(&self, pool: &Pool<Postgres>, migration_id: &str) -> bool {
        let result = sqlx::query("SELECT COUNT(*) as count FROM migrations.database_migrations WHERE id = $1")
            .bind(migration_id)
            .fetch_one(pool)
            .await;

        match result {
            Ok(row) => {
                let count: i64 = row.get("count");
                count > 0
            }
            Err(_) => false,
        }
    }

    pub async fn create_db_user(&self, pool: &Pool<Postgres>, user: &DbUser) -> Result<(), String> {
        let user_exists_result = sqlx::query("SELECT 1 FROM pg_roles WHERE rolname = $1")
            .bind(&user.name)
            .fetch_one(pool)
            .await;
        if user_exists_result.is_ok() {
            return Ok(());
        }
        let create_user_query = format!("CREATE USER {} WITH PASSWORD '{}';", user.name, user.password);
        let result = sqlx::query(&create_user_query)
            .execute(pool)
            .await;
        return match result {
            Ok(_) => Ok(()),
            Err(e) => Err(format!("Failed to create user {}: {}", user.name, e)),
        };
    }
}