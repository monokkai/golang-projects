CREATE TABLE IF NOT EXISTS urls
(
    id
    BIGINT
    AUTO_INCREMENT
    PRIMARY
    KEY,
    original_url
    TEXT
    NOT
    NULL,
    short_code
    VARCHAR
(
    20
) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NULL
    )