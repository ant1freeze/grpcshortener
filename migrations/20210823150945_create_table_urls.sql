-- +goose Up
CREATE TABLE IF NOT EXISTS urls (
	id serial, 
	longurl varchar(255), 
	shorturl varchar(255)
);

-- +goose Down
DROP TABLE urls;
