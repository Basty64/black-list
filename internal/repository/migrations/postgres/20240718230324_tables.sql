-- +goose Up
-- +goose StatementBegin
CREATE TABLE list(
                     id SERIAL PRIMARY KEY,
                     fullname VARCHAR(255) NOT NULL,
                     passport_id int NOT NULL,
                     is_employee boolean,
                     actual_company varchar,
                     reason varchar not null,
                     added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                     update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE passport(
                         id serial primary key,
                         passport_series int not null,
                         passport_number int not null,
                         previous_passport_series int,
                         previous_passport_number int
);

ALTER TABLE list
    ADD CONSTRAINT fk_passport
        FOREIGN KEY (passport_id)
            REFERENCES passport(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS list CASCADE;
DROP TABLE IF EXISTS passport CASCADE;

-- +goose StatementEnd


