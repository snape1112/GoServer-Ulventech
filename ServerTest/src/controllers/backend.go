package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ithunter101/GoServer-Ulventech/ServerTest/src/modules"
)

func Parsebook(c *gin.Context) {
	// Validate input
	text := c.PostForm("book_script")
	fmt.Println(text)

	topTenWords := modules.ParseText(text)

	fmt.Println(topTenWords[0].Word)
	//j, _ := json.Marshal(topTenWords)
	//log.Print(j)
	c.JSON(http.StatusOK, topTenWords)
}
