CREATE TABLE IF NOT EXISTS historico (
    id SERIAL PRIMARY KEY,
    origin TEXT NOT NULL,
    response_status_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_historico_origin_created_at ON historico (origin, created_at);
CREATE INDEX IF NOT EXISTS idx_historico_response_status_id ON historico (response_status_id);