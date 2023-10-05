create table users (
		id int primary key,
		name varchar(255) not null,
		email varchar(255) not null,
		password varchar(255) not null,
		created_at timestamp not null default now(),
		updated_at timestamp not null default now() 
);

INSERT into users (id, name, email) 
values (1, 'John Doe', 'john@email.com');	

--valid sql
-- failed to insert because password is not null
--response: ERROR: ERROR:  null value in column "password" violates not-null constraint
--correct:
-- INSERT into users (id, name, email, password) values (1, 'John Doe', 'john@email.com', 'password');	
-- Query 1 OK: INSERT 0 1, 1 row affected
