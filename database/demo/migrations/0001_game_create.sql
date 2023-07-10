-- +goose Up
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS questions (
	id SERIAL PRIMARY KEY,
	question VARCHAR(255) NOT NULL,
	user_id INTEGER NOT NULL references users(id),
	game_id INTEGER NOT NULL references games(id)
);

CREATE TABLE IF NOT EXISTS guesses (
	id SERIAL PRIMARY KEY,
	guess VARCHAR(255) NOT NULL,
	user_id INTEGER NOT NULL references users(id),
	game_id INTEGER NOT NULL references games(id),
	correct BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS games (
	id SERIAL PRIMARY KEY,
	host VARCHAR(255) NOT NULL,
	players VARCHAR(255)[5] NOT NULL references users(id),
	answer VARCHAR(255) NOT NULL,
	questions VARCHAR(255)[] references questions(id) ,
	guesses VARCHAR(255)[] references guesses(id),
	start_time TIMESTAMP NOT NULL DEFAULT NOW(),
	end_time TIMESTAMP,
	ended BOOLEAN DEFAULT FALSE
);

-- +goose Down

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS guesses;
DROP TABLE IF EXISTS games;

