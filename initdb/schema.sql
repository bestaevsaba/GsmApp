CREATE DATABASE IF NOT EXISTS saba_gsm_db;
USE saba_gsm_db;

CREATE TABLE IF NOT EXISTS gsm_data (
    id INT AUTO_INCREMENT PRIMARY KEY,
    recorded_at DATETIME NOT NULL
);