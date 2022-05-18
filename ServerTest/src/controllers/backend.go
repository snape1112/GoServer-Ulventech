package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Parsebook(c *gin.Context) {
	// Validate input
	text := c.PostForm("book_script")
	fmt.Println(text)

	topTenWords := bookparser.ParseText(text)

	fmt.Println(topTenWords[0].Word)
	//j, _ := json.Marshal(topTenWords)
	//log.Print(j)
	c.JSON(http.StatusOK, topTenWords)
}
