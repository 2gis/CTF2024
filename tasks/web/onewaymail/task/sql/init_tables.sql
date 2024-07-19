DROP TABLE IF EXISTS mails;
CREATE TABLE IF NOT EXISTS mails (id SERIAL PRIMARY KEY, title TEXT NOT NULL, content TEXT NOT NULL);
INSERT INTO mails (title, content) VALUES ('flag', '2GIS.CTF{sq1_1nj3ct10n_byp4ss3d}');