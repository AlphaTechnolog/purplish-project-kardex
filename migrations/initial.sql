CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(36) PRIMARY KEY,
    item_id VARCHAR(36),
    warehouse_id VARCHAR(36),
    kind VARCHAR(10) CHECK (kind IN ('input', 'output')),
    amount INT NOT NULL,
    created_at DATETIME,
    description VARCHAR(255),
    cost DECIMAL(10, 2)
);

CREATE INDEX idx_transaction_item_warehouse ON transactions (item_id, warehouse_id);

CREATE INDEX idx_transaction_created_at ON transactions (created_at);

CREATE VIEW IF NOT EXISTS stock_view AS
SELECT
    item_id,
    warehouse_id,
    SUM(amount) AS stock
FROM
    transactions
GROUP BY
    item_id,
    warehouse_id;

CREATE VIEW IF NOT EXISTS transactions_cost_view AS
SELECT
    item_id,
    warehouse_id,
    SUM(amount * cost) AS total_cost,
    SUM(amount) as amount
FROM
    transactions
GROUP BY
    item_id,
    warehouse_id;
