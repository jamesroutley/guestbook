CREATE TABLE visits (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url VARCHAR(64) NOT NULL,
    referrer VARCHAR(64) NULL,
    ip VARCHAR(64) NOT NULL,
    created DATE NOT NULL
);
