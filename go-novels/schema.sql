CREATE TABLE IF NOT EXISTS site (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    creation_date INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS novel (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    creation_date INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS novel_site (
    id INTEGER PRIMARY KEY,
    id_novel INTEGER NOT NULL,
    id_site INTEGER NOT NULL,
    url TEXT NOT NULL,
    language TEXT NOT NULL,
    creation_date INTEGER not null,
    FOREIGN KEY(id_site) REFERENCES site(id) ON DELETE CASCADE,
    FOREIGN KEY(id_novel) REFERENCES novel(id) ON DELETE CASCADE
);
