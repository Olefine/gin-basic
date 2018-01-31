package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	payload := gin.H{
		"title":   "Home Page",
		"payload": articles,
	}
	render(c, payload, "index.html")
}

func getArticle(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("article_id"))
	article, err := getArticleByID(articleID)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		payload := gin.H{
			"title":   article.Title,
			"payload": article,
		}
		render(c, payload, "article.html")
	}
}

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)
	router.GET("/articles/:article_id", getArticle)

	router.Run()
}
