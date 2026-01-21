CREATE TABLE IF NOT EXISTS tasks(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT now()
)