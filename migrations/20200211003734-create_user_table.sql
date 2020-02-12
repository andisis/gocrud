-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id          INTEGER       AUTO_INCREMENT  NOT NULL,
    fullname    VARCHAR(255)                  NOT NULL,
    email       VARCHAR(255)                  NOT NULL,
    username    VARCHAR(255)                  NOT NULL,
    created_at  TIMESTAMP                     NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                     NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE users;