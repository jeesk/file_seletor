package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ncruces/zenity"
	"net/http"
	"strings"
)

func main() {

	r := gin.Default()
	r.GET("/ping/:id", func(c *gin.Context) {
		file, err := zenity.SelectFileMultiple(
			zenity.Attach(0),
			zenity.Filename("/Users/mac/Desktop/"),
			zenity.FileFilters{
				{"Go files", []string{"*.go"}, false},
				{"Web files", []string{"*.html", "*.js", "*.css"}, true},
				{"Image files", []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}, true},
			})

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"files": strings.Join(file, ","),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
