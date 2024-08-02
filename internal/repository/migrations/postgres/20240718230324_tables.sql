-- +goose Up
-- +goose StatementBegin
CREATE TABLE list(
                     id SERIAL PRIMARY KEY,
                     fullname VARCHAR(255) NOT NULL,
                     passport int UNIQUE NOT NULL,
                     is_employee boolean,
                     actual_company varchar,
                     reason varchar not null,
                     added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                     update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE IF EXISTS list OWNER to "admin";

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS list CASCADE;

-- +goose StatementEnd


