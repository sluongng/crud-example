CREATE TABLE IF NOT EXISTS user
(
    id int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    username varchar(32) NOT NULL,
    email varchar(32) NOT NULL,
    website varchar(255)
);
CREATE UNIQUE INDEX user_id_uindex ON user (id);
CREATE UNIQUE INDEX user_username_uindex ON user (username);
CREATE UNIQUE INDEX user_email_uindex ON user (email)