-- name: GetUser :one
SELECT * FROM users 
WHERE username = $1;

-- name: UpdateUser :exec
INSERT INTO users 
(username) VALUES ($1)
ON CONFLICT (username) 
DO NOTHING;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;

-- name: GetAUsersGames :many
SELECT * FROM games
join users on users.id = games.user_id
WHERE users.username = $1;
