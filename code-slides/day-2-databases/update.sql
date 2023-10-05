CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY, 
	name TEXT, 
	email TEXT UNIQUE, 
	password TEXT NOT NULL);

INSERT INTO users (name, email, password) VALUES 
	('John Doe', ' john.doe@email.com', 'opensesame')
INSERT INTO users (name, email, password) VALUES 
('Jane Smith', 'smithjane@email.com', 'opensesame')

UPDATE users SET name = 'Alfred Doe' WHERE id = 1;

SELECT * FROM users;
