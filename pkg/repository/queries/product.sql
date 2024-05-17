-- name: create-product
INSERT INTO eulabs.product (name, description, image_url) VALUES (?, ?, ?);

-- name: create-price
INSERT INTO eulabs.product_price (product_id, value, currency) VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE value = VALUES(value);

-- name: get-product
SELECT id, name, description, image_url
FROM eulabs.product
WHERE id = ?;

-- name: get-price
SELECT value, currency
FROM eulabs.product_price
WHERE product_id = ?;

-- name: update-product
UPDATE eulabs.product SET name = ?, description = ?, image_url = ? WHERE id = ?;

-- name: delete-product
DELETE FROM eulabs.product WHERE id = ?;

-- name: delete-price
DELETE FROM eulabs.product_price WHERE product_id = ?;

-- name: delete-price-by-currency
DELETE FROM eulabs.product_price WHERE product_id = ? AND currency = ?;
