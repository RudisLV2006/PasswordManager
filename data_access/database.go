package data_access

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"server/api/model"

	"golang.org/x/crypto/pbkdf2"

	_ "github.com/glebarez/go-sqlite"
)

func CreateTables(dbFile string) {
	db, err := sql.Open("sqlite", dbFile)

	if err != nil {
		log.Fatal("Can't make connect")
	}
	defer db.Close()

	sqlStatment := `
		CREATE TABLE IF NOT EXISTS users(
	user_id INTEGER PRIMARY KEY,
	username CHAR(16) NOT NULL,
	name VARCHAR(20),
	password TEXT NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts(
	account_id INTEGER PRIMARY KEY,
	username CHAR(24),
	encrypted_password TEXT NOT NULL,
	salt TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES user(user_id)
);
CREATE TABLE IF NOT EXISTS websites (
	site_id INTEGER PRIMARY KEY,
	site CHAR(30) NOT NULL,
	url TEXT
);

CREATE TABLE IF NOT EXISTS account_site(
	account_id INTEGER NOT NULL,
	site_id INTEGER NOT NULL,
	PRIMARY KEY (account_id, site_id),
	FOREIGN KEY (account_id) REFERENCES account(account_id) ON DELETE CASCADE,
	FOREIGN KEY (site_id) REFERENCES website(site_id)
);
		`
	_, err = db.Exec(sqlStatment)
	if err != nil {
		log.Fatal("Error creating tables:")
	}
	fmt.Println("Database and table created successfully!")
}

func InsertWebsite(website *model.Website, dbFile string) {
	db, err := sql.Open("sqlite", dbFile)

	if err != nil {
		log.Fatal("Can't make connect")
	}
	defer db.Close()
	var insertStatement string = `INSERT INTO websites (site,url) VALUES (?,?);`

	_, err = db.Exec(insertStatement, toNullString(website.GetSite()), toNullString(website.GetURL()))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Website inserted successfully!")
}

func InsertAccount(account *model.Account, dbFile string) {
	db, err := sql.Open("sqlite", dbFile)

	if err != nil {
		log.Fatal("Can't make connect")
	}

	key := DeriveEncryptionKey(account.GetKey(), account.GetSalt())

	defer db.Close()
	var insertStatement string = `INSERT INTO accounts (username,encrypted_password,salt,user_id) VALUES (?,?,?,?);`
	_, err = db.Exec(insertStatement, toNullString(account.GetUsername()),
		toNullString(base64.StdEncoding.EncodeToString(encryptIt([]byte(account.GetPassword()), key))),
		toNullString(base64.StdEncoding.EncodeToString(account.GetSalt())), 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Website inserted successfully!")
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func DeriveEncryptionKey(sk string, salt []byte) []byte {
	const iteration int = 10000
	return pbkdf2.Key([]byte(sk), salt, iteration, 32, sha1.New)
}

func encryptIt(value []byte, encryptedKey []byte) []byte {
	aesBlock, err := aes.NewCipher(encryptedKey)
	if err != nil {
		return nil
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return nil
	}

	nonce := make([]byte, gcmInstance.NonceSize())

	if _, err := rand.Read(nonce); err != nil {
		return nil
	}

	ciphertext := gcmInstance.Seal(nil, nonce, value, nil)

	return append(nonce, ciphertext...)
}

func DecryptIt(ciphered []byte, encryptedKey []byte) string {
	aesBlock, err := aes.NewCipher(encryptedKey)
	if err != nil {
		return ""
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return ""
	}

	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	plaintext, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		return ""
	}

	plaintextStr := string(plaintext)
	return plaintextStr
}
