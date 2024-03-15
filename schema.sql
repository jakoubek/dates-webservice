CREATE TABLE IF NOT EXISTS request (
    id INTEGER PRIMARY KEY,
    request TEXT,
    lang TEXT,
    format TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
)
