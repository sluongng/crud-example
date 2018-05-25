CREATE TABLE user_db.user
(
    id int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    username varchar(32) NOT NULL,
    email varchar(32) NOT NULL,
    website varchar(255)
);
INSERT INTO user_db.user (id, username, email, website) VALUES (1, 'Son Luong', 'sluongng@gmail.com', 'https://sluongng.gitlab.io');