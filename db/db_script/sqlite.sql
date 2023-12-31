.open ./sqlite/datareporter.db
CREATE TABLE articles (
    article_id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    author_code TEXT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    created_on TEXT NOT NULL,
    updated_on TEXT NOT NULL
);

INSERT INTO articles (title,body,author_code,created_on,updated_on) VALUES
('SQLite One to One', 'One to One, two to two...', 'CNV', DATETIME('now'), DATETIME('now')),
('SQLite More five days', 'One day, two days, tree days, for days and five days', 'CNV', DATETIME('now'), DATETIME('now')),
('SQLite by all', 'The half all is all by half', 'DNV', DATETIME('now'), DATETIME('now'));
