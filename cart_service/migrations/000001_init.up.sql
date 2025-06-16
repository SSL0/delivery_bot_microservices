BEGIN;
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price MONEY,
    description TEXT,
    type VARCHAR(255)
);

CREATE TABLE toppings (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products (id) ON DELETE CASCADE,
    name VARCHAR(255),
    price MONEY
);

COMMIT;