CREATE TABLE url
(
    id         VARCHAR(36) PRIMARY KEY,
    slug       VARCHAR(128) NOT NULL UNIQUE,
    url        TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);