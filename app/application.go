package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mohitjindal78/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start application...")
	router.Run(":8080")
}
