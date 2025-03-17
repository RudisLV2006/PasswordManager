CREATE DATABASE IF NOT EXISTS PassManagment;

CREATE TABLE IF NOT EXISTS user(
	user_id INT,
	username CHAR(16) NOT NULL,
	name VARCHAR(20),
	password TEXT NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS account(
	account_id INT,
	username CHAR(24),
	encrypted_password TEXT NOT NULL,
	salt TEXT NOT NULL,
	user_id INT,
	PRIMARY KEY (account_id),
	FOREIGN KEY (user_id) REFERENCES user(user_id)
);
CREATE TABLE IF NOT EXISTS website (
	site_id INT,
	site CHAR(30) NOT NULL,
	url TEXT,
	PRIMARY KEY (site_id)
);

CREATE TABLE IF NOT EXISTS account_site(
	account_id INT,
	site_id INT,
	PRIMARY KEY (account_id, site_id),
	FOREIGN KEY (account_id) REFERENCES account(account_id) ON DELETE CASCADE,
	FOREIGN KEY (site_id) REFERENCES website(site_id)
);