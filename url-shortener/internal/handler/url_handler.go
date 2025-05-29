package handler

import (
	"math/rand"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Хранилище для URL
var (
	urlStore = make(map[string]string)
	mutex    sync.RWMutex
)

// Генерация случайного кода
func generateCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func HandleShorten(c *gin.Context) {
	var input struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерируем уникальный код
	code := generateCode()

	// Сохраняем URL
	mutex.Lock()
	urlStore[code] = input.URL
	mutex.Unlock()

	// Возвращаем сокращенный URL
	shortenedURL := "http://localhost:8080/" + code
	c.JSON(http.StatusOK, gin.H{
		"shortenedUrl": shortenedURL,
		"originalUrl":  input.URL,
		"code":         code,
	})
}

func HandleRedirect(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}

	// Получаем оригинальный URL
	mutex.RLock()
	originalURL, exists := urlStore[code]
	mutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Перенаправляем на оригинальный URL
	c.Redirect(http.StatusMovedPermanently, originalURL)
}
