

CREATE TABLE users(
name VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL,
password VARCHAR(255) NOT NULL, -- Hash of password
acct_created TIMESTAMP NOT NULL,
last_login TIMESTAMP NOT NULL,
user_type INTEGER NOT NULL,  -- How much power the user has
id SERIAL PRIMARY KEY
);


CREATE TABLE posts(
title VARCHAR(255) NOT NULL,
content TEXT NOT NULL,
user_id INTEGER REFERENCES user(id) 
id SERIAL PRIMARY KEY
);

INSERT INTO users (name, email, password, acct_created, last_login, user_type) 
VALUES 
('Derek Banas', 'derekbanas@aol.com', 'abcdefghij', current_timestamp, current_timestamp, 3);

