CREATE table users (
		id int primary key,
		name varchar(255) not null,
		email varchar(255) not null,
		password varchar(255) not null,
		created_at timestamp not null default NOW(),
		updated_at timestamp not null default NOW() 
);

INSERT into users (id, name, email) 
values (1, 'John Doe', 'john@email.com');	
INSERT into users (id, name, email, password) 
values (1, 'Jane Doe', 'jane@email.com', 'password1');	

--valid sql
-- failed to insert because id is primary key and already exist
-- response: ERROR:  duplicate key value violates unique constraint "users_pkey"
-- in postgres use ON CONFLICT UPDATE
-- in mysql use ON DUPLICATE KEY UPDATE
-- is msql use Merge INTO

--correct:INSERT into users (id, name, email, password) values (2, 'Jane Doe', 'jane@email.com', 'password1');	

-- follow up question: should another value be unique?
-- follow up question: what can we do to make sure that the id is unique?
-- answer: use sequence, serial, or auto increment
