-- +goose Up
-- +goose StatementBegin
CREATE TABLE units
(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    singular_name TEXT check ( singular_name is null or length(singular_name) > 0 ),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products
(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    unit_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (unit_id) REFERENCES units (id) ON DELETE CASCADE
);

CREATE TABLE inventory
(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inventory;
DROP TABLE products;
DROP TABLE units;
-- +goose StatementEnd
