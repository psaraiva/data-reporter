CREATE DATABASE datareporter WITH ENCODING = 'UTF8';
CREATE USER app WITH ENCRYPTED PASSWORD 'app-my-super-pass';
GRANT ALL PRIVILEGES ON DATABASE datareporter TO app;
\c datareporter
CREATE TABLE IF NOT EXISTS articles (
   article_id serial PRIMARY KEY,
   title VARCHAR (150) NOT NULL,
   body TEXT NOT NULL,
   author_code VARCHAR (50) NOT NULL,
   version SMALLINT NOT NULL,
   is_active BOOLEAN NOT NULL DEFAULT TRUE,
   created_on TIMESTAMP NOT NULL,
   updated_on TIMESTAMP NOT NULL
);
INSERT INTO articles (title,body,author_code,version,created_on,updated_on)
VALUES
('PG One to One', 'One to One, two to two...', 'CNV', 1, now(), now()),
('PG More five days', 'One day, two days, tree days, for days and five days', 'CNV', 1, now(), now()),
('PG by all', 'The half all is all by half', 'DNV', 1, now(), now());
GRANT SELECT, INSERT, UPDATE ON articles TO app;
