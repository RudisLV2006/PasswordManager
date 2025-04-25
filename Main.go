package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"server/api/data_access"
	"server/api/model"

	"github.com/gin-gonic/gin"
)

var site []model.Website
var db = data_access.MakeConnection()

func getWebsite(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, site)
}
func postWebsite(c *gin.Context) {
	var newSite = *model.CreateWebsite()

	if err := c.BindJSON(&newSite); err != nil {
		return
	}

	site = append(site, newSite)

	data_access.InsertWebsite(&newSite, db)

	c.IndentedJSON(http.StatusCreated, newSite)
}
func postAccount(c *gin.Context) {
	var newAccount = *model.CreateAccount()

	if err := c.BindJSON(&newAccount); err != nil {
		return
	}

	newAccount.AccountName = "Kristaps"
	newAccount.SetSalt(makeSalt())

	err := data_access.CreateAccountAndLinkSite(&newAccount, db)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newAccount)
}
func debug(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, len(site))
}
func main() {
	data_access.ApplyMigrations()

	site = data_access.SelectSite(db)

	fmt.Printf("DEBUG: %+v\n", site)

	router := gin.Default()
	router.GET("/website", getWebsite)
	router.GET("/debug", debug)
	router.POST("/website", postWebsite)
	router.POST("/account", postAccount)

	router.Run("localhost:8080")
}

func makeSalt() []byte {
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal("Cant create salt")
	}
	return salt
}
