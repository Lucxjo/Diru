use thiserror::Error;
use tracing::*;
use tracing_subscriber::EnvFilter;

#[derive(Error, Debug)]
pub enum BotError {
    #[error("A database error occurred")]
    Database(#[from] sea_orm::DbErr),
}

async fn run_bot() -> Result<(), BotError> {
    let backup_conn_uri = format!(
        "sqlite://{}/diru/diru.db?mode=rwc",
        dirs::data_dir().unwrap().to_str().unwrap()
    );
    let conn_uri: &str = &std::env::var("DATABASE_URI").unwrap_or(backup_conn_uri);
    let db = sea_orm::Database::connect(conn_uri).await?;

    Ok(())
}

fn main() {
    tracing_subscriber::fmt()
        .with_env_filter(EnvFilter::from_default_env())
        .pretty()
        .init();
    let rt = tokio::runtime::Builder::new_multi_thread()
        .enable_all()
        .build().unwrap();
    let bot = rt.spawn(run_bot());

    // Ensure tasks are finished.
    loop {
        if bot.is_finished() {
            return;
        }
    }
}
