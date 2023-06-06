package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type list struct {
	ID          string `json:"id"`
	NextPageKey string `json:"nextPageKey"`
}

type page struct {
	ID          string `json:"id"`
	Articles    string `json:"articles"`
	NextPageKey string `json:"nextPageKey"`
}

var lists = []list{
	{ID: "life", NextPageKey: "life_1"},
	{ID: "love", NextPageKey: "love_1"},
	{ID: "work", NextPageKey: "work_1"},
}

var pages = []page{
	{ID: "life_1", Articles: "life_content_001", NextPageKey: "life_2"},
	{ID: "life_2", Articles: "life_content_002", NextPageKey: "life_3"},
	{ID: "life_3", Articles: "life_content_003", NextPageKey: ""},
	{ID: "love_1", Articles: "love_content_001", NextPageKey: "love_2"},
	{ID: "love_2", Articles: "love_content_002", NextPageKey: "love_3"},
	{ID: "love_3", Articles: "love_content_003", NextPageKey: ""},
	{ID: "work_1", Articles: "work_content_001", NextPageKey: "work_2"},
	{ID: "work_2", Articles: "work_content_002", NextPageKey: "work_3"},
	{ID: "work_3", Articles: "work_content_003", NextPageKey: ""},
}

func main() {
	router := gin.Default()
	router.GET("/getHead/:id", getHead)
	router.GET("/getPage/:id", getPage)

	router.Run("localhost:8080")
}

// func postAlbums(c *gin.Context) {
// 	var newAlbum album
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

func getHead(c *gin.Context) {
	id := c.Param("id")

	for _, a := range lists {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
}

func getPage(c *gin.Context) {
	id := c.Param("id")

	for _, a := range pages {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pagea not found"})
}
