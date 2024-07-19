DROP TABLE IF EXISTS mails;
CREATE TABLE IF NOT EXISTS mails (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    content TEXT NOT NULL DEFAULT '',
    hkey BIGINT NOT NULL DEFAULT 0
);
INSERT INTO mails (title, content) VALUES ('flag', '2GIS.CTF{sh4_256_47t4ck_d0n3_7h1s_15_w4s_v3ry_34sy_f0r_y0u}');