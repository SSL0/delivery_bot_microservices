BEGIN;
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price MONEY,
    description TEXT,
    type VARCHAR(255)
);
COMMIT;