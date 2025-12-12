pub struct MigrationConfig {
    pub name: String,
    pub connection_string: String,
}

#[derive(Debug, Clone)]
pub struct Migration {
    pub connection_string: Option<String>,
    pub sql_files: Vec<String>,
}

impl Migration {
    pub fn new(connection_string: String, sql_files: Vec<String>) -> Self {
        let mut files = sql_files.clone();
        files.sort();
        return Migration {
            connection_string: Some(connection_string),
            sql_files: files,
        };
    }
}

#[derive(Debug, Clone)]
pub struct DbUser {
    pub name: String,
    pub password: String,
}
