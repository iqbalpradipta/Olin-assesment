SELECT u.name, o.amount, o.created_at
FROM dblink('dbname=first user=postgres password=mbangg12', 'SELECT id, name FROM users') AS u(id INT, name TEXT)
JOIN dblink('dbname=second user=postgres password=mbangg12', 'SELECT user_id, amount, created_at FROM orders') AS o(user_id INT, amount NUMERIC, created_at TIMESTAMP)
ON u.id = o.user_id;