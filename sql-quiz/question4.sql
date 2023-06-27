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

SELECT COUNT(*) from transactions;

-- valid
-- response: 4
