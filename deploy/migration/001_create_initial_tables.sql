CREATE TABLE eulabs.product
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255)                                                   NOT NULL,
    description TEXT                                                           NULL,
    image_url   VARCHAR(255)                                                   NULL,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NULL,
    INDEX idx_product_name (name)
);

CREATE TABLE eulabs.product_price
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    product_id  INT                                                            NOT NULL,
    value       INT                                                            NOT NULL COMMENT 'Price in cents',
    currency    CHAR(3)                                                        NOT NULL COMMENT 'ISO 4217 currency code',
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NULL,
    CONSTRAINT uidx_product_price_product_id_currency UNIQUE (product_id, currency),
    INDEX idx_product_price_product_id_currency (product_id, currency),
    INDEX idx_product_price_product_id (product_id),
    INDEX idx_product_price_currency (currency),
    FOREIGN KEY (product_id) REFERENCES eulabs.product (id)
);
