use clap::{Parser, Subcommand, ValueEnum};
use std::str::FromStr;
use std::fmt;



#[derive(Parser)]
#[command(name = "fr_cli")]
#[command(about = "Developer utility CLI for Forgotten Realms", long_about = None)]
pub struct Args {
    #[command(subcommand)]
    pub command: Commands,
}

#[derive(Subcommand)]
pub enum Commands {
    // Generate default configuration
    GenerateConfig {
        #[arg(short, long)]
        server_type : ServerType,
        #[arg(short, long)]
        path: String,
    },

    // Run database migrations
    Migrate {
        #[arg(short, long)]
        migrations_folder: String,
        #[arg(short, long)]
        users: Vec<DbUser>,
        #[arg(short, long)]
        connections: Vec<DbConnection>,
    },
}

#[derive(ValueEnum, Clone, Debug)]
pub enum ServerType {
    World,
    Auth,
}

#[derive(Clone, Debug)]
pub struct DbUser {
    pub name: String,
    pub password: String,
}

impl FromStr for DbUser {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts: Vec<&str> = s.splitn(2,':').collect();
        if parts.len() == 2 {
            Ok(DbUser {
                name: parts[0].to_string(),
                password: parts[1].to_string(),
            })
        } else {
            Err("DbUser must be in the format 'username:password'".to_string())
        }
    }
}

impl fmt::Display for DbUser {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}:{}", self.name, self.password)
    }
}

impl FromStr for DbConnection {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts: Vec<&str> = s.splitn(2, ':').collect();
        if parts.len() == 2 {
            Ok(DbConnection {
                name: parts[0].to_string(),
                connection_string: parts[1].to_string(),
            })
        } else {
            Err("DbConnection must be in the format 'name:connection_string'".to_string())
        }
    }
}

impl fmt::Display for DbConnection {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}:{}", self.name, self.connection_string)
    }
}


#[derive(Clone)]
pub struct DbConnection {
    pub name: String,
    pub connection_string: String,
}
