use thiserror::Error;
use tracing_subscriber::EnvFilter;

#[derive(Error, Debug)]
pub enum BotError {
    #[error("A database error occurred")]
    Database(#[from] sea_orm::DbErr),
}

async fn run_bot() -> Result<(), sea_orm::DbErr> {
    println!("Saluton, mondo!");

    Ok(())
}

fn main() {
    tracing_subscriber::fmt()
        .with_env_filter(EnvFilter::from_default_env())
        .pretty()
        .init();
    let rt = tokio::runtime::Builder::new_multi_thread().build().unwrap();
    rt.spawn(run_bot());
}
