CREATE table transactions(
	id int primary key,
	user_id int not null,
	amount int not null,
	created_at timestamp not null default NOW(),
	updated_at timestamp not null default NOW()
	foreign key (user_id) references users(id)
);

INSERT into transactions (id, user_id, amount) values (1, 1, 1000);
INSERT into transactions (id, user_id, amount) values (2, 2, 2000);
INSERT into transactions (id, user_id, amount) values (3, 2, 3000);
INSERT into transactions (id, user_id, amount) values (4, 1, 4000);

UPDATE transactions set amount = 5000 where user_id = 1;

select * from transactions;

-- response 
-- 2	2	2000	2023-06-27 20:22:33.719055	2023-06-27 20:22:33.719055
-- 3	2	3000	2023-06-27 20:22:33.764045	2023-06-27 20:22:33.764045
-- 4	1	5000	2023-06-27 20:22:33.767405	2023-06-27 20:22:33.767405
-- 1	1	5000	2023-06-27 20:26:08.17272	2023-06-27 20:26:08.17272

-- what is the problem with this query?
-- answer: it updates all the rows with user_id = 1
-- how can we fix this?
-- answer: add where clause
