package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"server/api/data_access"
	"server/api/model"
)

func main() {
	dbFile := "sql/PassMangerDB.db"
	scanner := bufio.NewScanner(os.Stdin)

	if _, err := os.Stat(dbFile); errors.Is(err, os.ErrNotExist) {
		data_access.CreateTables(dbFile)
	}
	for {
		fmt.Print("Enter your choise:\n", "1:Insert website \n", "2:Insert Account: ")
		if scanner.Scan() {
			output := scanner.Text()

			// Handle the user's input based on choice
			switch output {
			case "1":
				// Ask for website details
				website := model.CreateWebsite()
				/*
					website := model.CreateWebsite()
					fmt.Print(&website)
				*/

				fmt.Println("Enter site name: ")
				if scanner.Scan() {
					website.SetSite(scanner.Text()) // Read site name
				}

				// Ask for the website URL
				fmt.Println("Enter site URL (optional):")
				if scanner.Scan() {
					website.SetURL(scanner.Text()) // Read site URL
				}

				// Insert website into database
				data_access.InsertWebsite(website, dbFile)

			case "2":
				// fmt.Println("I will be implemented")
				account := model.CreateAccount()

				fmt.Println("Enter account username")
				if scanner.Scan() {
					account.SetUsername(scanner.Text())
				}
				account.SetAccountName("Kristaps")
				fmt.Println("Enter account password")
				if scanner.Scan() {
					account.SetPassword(scanner.Text())
				}
				fmt.Print("I am implemented")

			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}
	}
}
