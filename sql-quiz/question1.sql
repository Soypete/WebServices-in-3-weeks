CREATE table users (
		id int primary key,
		name varchar(255) not null,
		email varchar(255) not null,
		password varchar(255) not null,
		created_at timestamp not null default NOW(),
		updated_at timestamp not null default NOW() 
);

--valid
--create table with primary key
