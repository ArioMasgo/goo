CREATE TABLE residence_robot
(
    id                CHAR(36) NOT NULL PRIMARY KEY,
    residence_id      CHAR(36) NOT NULL,
    prompt_tokens     INT      NOT NULL,
    completion_tokens INT      NOT NULL,
    total_tokens      INT      NOT NULL,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (residence_id) REFERENCES residencias (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = DYNAMIC;