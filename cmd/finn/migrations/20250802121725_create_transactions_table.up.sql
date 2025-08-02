CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    amount DECIMAL(12,2) NOT NULL,
    category_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE INDEX idx_transactions_category_id ON transactions(category_id);


CREATE TABLE monthly_budget_summary (
    month DATE PRIMARY KEY,
    total_amount DECIMAL(12,2)
);


INSERT INTO monthly_budget_summary (month, total_amount)
SELECT
    DATE_FORMAT(created_at, '%Y-%m-01') AS month,
    SUM(amount)
FROM transactions
GROUP BY month;

DELIMITER //

CREATE EVENT IF NOT EXISTS refresh_monthly_budget_summary
ON SCHEDULE EVERY 1 DAY
DO
BEGIN
    DELETE FROM monthly_budget_summary;
    INSERT INTO monthly_budget_summary (month, total_amount)
    SELECT
        DATE_FORMAT(created_at, '%Y-%m-01') AS month,
        SUM(amount)
    FROM transactions
    GROUP BY month;
END //

DELIMITER ;