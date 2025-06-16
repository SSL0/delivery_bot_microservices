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

CREATE TABLE carts ( id SERIAL PRIMARY KEY, user_id BIGINT );

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INT REFERENCES carts (id) ON DELETE CASCADE,
    item_id INT,
    type TEXT NOT NULL CHECK (
        type IN ('product', 'topping')
    ),
    price MONEY,
    quantity INT CHECK (quantity > 0)
);

COMMIT;