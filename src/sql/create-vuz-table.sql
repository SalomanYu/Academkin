CREATE TABLE vuz (
    id SERIAL,
    shortname TEXT NOT NULL,
    fullname TEXT NOT NULL,
    logo TEXT,
    city TEXT,
    locality TEXT,
    PRIMARY KEY (fullname)
);