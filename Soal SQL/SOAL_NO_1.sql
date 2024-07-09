SELECT name, email, amount AS total_spend
FROM users
LEFT JOIN orders ON users.id = orders.user_id
WHERE created_at >= '2022-01-01'
GROUP BY name, email, amount
HAVING SUM(amount) >= 100;