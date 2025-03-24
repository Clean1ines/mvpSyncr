// config.rs
use std::env;

/// Структура конфигурации приложения.
pub struct Config {
    pub telegram_token: String,
}

impl Config {
    /// Загружает конфигурацию из переменных окружения.
    pub fn from_env() -> Self {
        let telegram_token = env::var("TELEGRAM_TOKEN")
            .expect("TELEGRAM_TOKEN must be set in environment variables");
        Self { telegram_token }
    }
}