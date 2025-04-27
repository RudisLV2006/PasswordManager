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
	"strconv"

	"golang.org/x/crypto/pbkdf2"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var dbFile = "sql/PassManagerDB.db"

func ApplyMigrations() {

	m, err := migrate.New(
		"file://sql/migrations",
		"sqlite://"+dbFile,
	)
	if err != nil {
		log.Fatalf("Could not create migration object: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migrations: %v", err)
	} else if err == migrate.ErrNoChange {
		// You might want to log this case in a less fatal way
		fmt.Println("No new migrations to apply.")
	} else {
		fmt.Println("Migrations applied successfully!")
	}
}

func InsertWebsite(website *model.Website, db *sql.DB) {
	var insertStatement string = `INSERT INTO websites (site,url) VALUES (?,?);`

	_, err := db.Exec(insertStatement, toNullString(website.Site), toNullString(website.Url))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Website inserted successfully!")
}

func InsertAccount(account *model.Account, tx *sql.Tx) (string, error) {
	key := DeriveEncryptionKey(account.Secret_key, account.GetSalt())

	const insertStatement string = `INSERT INTO accounts (username,encrypted_password,salt,user_id) VALUES (?,?,?,?);`
	res, err := tx.Exec(insertStatement, toNullString(account.Username),
		toNullString(base64.StdEncoding.EncodeToString(encryptIt([]byte(account.Password), key))),
		toNullString(base64.StdEncoding.EncodeToString(account.GetSalt())), 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Website inserted successfully!")

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(lastInsertId, 10), nil
}
func SelectSiteID(site string, tx *sql.Tx) (string, error) {
	var siteID string
	query := "SELECT site_id FROM websites WHERE site = ?"

	err := tx.QueryRow(query, site).Scan(&siteID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no site found with name: %s", site)
		}
		return "", fmt.Errorf("query error: %v", err)
	}

	return siteID, nil
}

func SelectSite(db *sql.DB) []model.Website {
	query := "SELECT site,url FROM websites"
	rows, err := db.Query(query)

	if err != nil {
		return nil
	}
	defer rows.Close()
	var website []model.Website

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var web model.Website
		var url sql.NullString

		if err := rows.Scan(&web.Site, &url); err != nil {
			log.Println("Error scanning row:", err) // Log the error instead of returning
			continue                                // Skip this row and move to the next one
		}
		if url.Valid {
			web.Url = url.String
		} else {
			web.Url = ""
		}
		website = append(website, web)
	}
	if err = rows.Err(); err != nil {
		return website
	}
	return website
}
func SearchSite(db *sql.DB, searchTerm string) error {
	query := "SELECT site FROM websites WHERE site LIKE ?"
	siteName := "%" + searchTerm + "%" // assuming searchTerm is the value you're looking for

	rows, err := db.Query(query, siteName)

	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil // No matching site found, return nil (no error)
	}

	var foundSite string
	if err := rows.Scan(&foundSite); err != nil {
		return err // return the error if there was an issue scanning
	}

	// If you reach here, that means a matching site was found
	return fmt.Errorf("site already exists: %s", foundSite)
}

func CreateAccountAndLinkSite(account *model.Account, db *sql.DB) error {
	// Begin a new transaction
	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Insert account within the transaction
	accountID, err := InsertAccount(account, tx)
	if err != nil {
		return fmt.Errorf("error inserting account: %v", err)
	}

	// Select site ID within the transaction
	siteID, err := SelectSiteID(account.Site, tx)
	if err != nil {
		return fmt.Errorf("error selecting site: %v", err)
	}

	// Link the account to the site within the transaction
	err = LinkedTable(accountID, siteID, tx)
	if err != nil {
		return fmt.Errorf("error linking account to site: %v", err)
	}

	// If everything succeeds, commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
func LinkedTable(accID string, sitID string, tx *sql.Tx) error {
	const insertStatement string = `INSERT INTO account_site VALUES (?,?);`
	_, err := tx.Exec(insertStatement, accID, sitID)
	if err != nil {
		return err
	}
	return nil
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
func MakeConnection() *sql.DB {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
