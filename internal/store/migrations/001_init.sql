CREATE TABLE IF NOT EXISTS sessions (
    platform      TEXT     NOT NULL PRIMARY KEY,
    access_token  BLOB     NOT NULL,  -- AES-GCM encrypted
    refresh_token BLOB,               -- AES-GCM encrypted, nullable
    expires_at    DATETIME NOT NULL,
    scope         TEXT,
    raw_session   BLOB,               -- AES-GCM encrypted; Kick: serialized cookie map
    created_at    DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at    DATETIME NOT NULL DEFAULT (datetime('now'))
);

-- future tables go below
