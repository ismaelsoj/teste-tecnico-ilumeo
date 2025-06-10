CREATE TABLE IF NOT EXISTS historico (
    id SERIAL PRIMARY KEY,
    origin TEXT NOT NULL,
    response_status_id INTEGER NOT NULL
);