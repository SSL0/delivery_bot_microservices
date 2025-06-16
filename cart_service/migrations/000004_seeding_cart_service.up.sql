BEGIN;

-- Очищаем таблицы и сбрасываем sequence (опционально)
TRUNCATE TABLE cart_items, carts RESTART IDENTITY CASCADE;

-- Вставляем тестовые корзины для пользователей
INSERT INTO
    carts (user_id)
VALUES (1001), -- Корзина 1
    (1002), -- Корзина 2
    (1003), -- Корзина 3
    (1004); -- Корзина 4 (пустая)

-- Вставляем элементы корзины для пользователя 1001 (корзина 1)
INSERT INTO
    cart_items (
        cart_id,
        item_id,
        type,
        price,
        quantity
    )
VALUES (1, 101, 'product', 10.99, 2), -- 2 продукта с ID 101
    (1, 102, 'product', 8.50, 1), -- 1 продукт с ID 102
    (1, 201, 'topping', 1.50, 3), -- 3 топпинга с ID 201
    (1, 202, 'topping', 0.75, 2); -- 2 топпинга с ID 202

-- Вставляем элементы корзины для пользователя 1002 (корзина 2)
INSERT INTO
    cart_items (
        cart_id,
        item_id,
        type,
        price,
        quantity
    )
VALUES (2, 103, 'product', 12.99, 1),
    (2, 203, 'topping', 1.25, 1),
    (2, 204, 'topping', 0.50, 4);

-- Вставляем элементы корзины для пользователя 1003 (корзина 3)
INSERT INTO
    cart_items (
        cart_id,
        item_id,
        type,
        price,
        quantity
    )
VALUES (3, 104, 'product', 15.00, 3),
    (3, 205, 'topping', 2.00, 1);

-- Корзина 4 (ID 4) оставляем пустой (для демонстрации)

COMMIT;