package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleShorten(c *gin.Context) {
	var input struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shortenedUrl": "http://localhost:8080/shorten",
	})
}

func HandleRedirect(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "https://example.com")
}
