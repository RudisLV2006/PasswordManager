package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"server/api/model"

	"github.com/gin-gonic/gin"
)

var site []model.Website

func getWebsite(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, site)
}
func postWebsite(c *gin.Context) {
	var newSite = *model.CreateWebsite()

	if err := c.BindJSON(&newSite); err != nil {
		return
	}

	site = append(site, newSite)
	fmt.Printf("DEBUG NewSite Addr: %p, Data: %+v\n", &newSite, newSite)

	fmt.Printf("Current slice: %+v\n", site)

	c.IndentedJSON(http.StatusCreated, newSite)
}
func debug(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, len(site))
}
func main() {
	site = append(site, *model.CreateWebsite())
	site = append(site, *model.CreateWebsite())
	site[0].Site = "Test"
	site[0].Url = "tet"
	site[1].Site = "sdsd"
	site[1].Url = "sfasf"

	fmt.Printf("DEBUG: %+v\n", site)

	router := gin.Default()
	router.GET("/website", getWebsite)
	router.GET("/debug", debug)
	router.POST("/website", postWebsite)

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
