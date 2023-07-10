INSERT into users (id, name, email) 
values (1, 'John Doe', 'john@email.com');	

--valid sql
-- failed to insert because password is not null
--response: ERROR: ERROR:  null value in column "password" violates not-null constraint
--correct:
-- INSERT into users (id, name, email, password) values (1, 'John Doe', 'john@email.com', 'password');	
-- Query 1 OK: INSERT 0 1, 1 row affected
