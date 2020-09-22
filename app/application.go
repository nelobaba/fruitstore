package app

import (
	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/gin-gonic/gin"
)

// we are initializing the gin router here, same package as url_mapping
// with logger and recovery middleware setup
var (
	router = gin.Default()
)

// the only layer where we use the http server is in the StartApplication
// as well as your controller
func StartApplication() {
	// because this is defined in the same package as application.go, it is fine
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8082")
}