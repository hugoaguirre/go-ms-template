package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	config "github.com/hugoaguirre/go-ms-template/internal"
)

var AllowedHeaders = []string{"Origin", "Content-Type"}

func StartServer(config config.Config) error {
	e := gin.New()
	e.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/"),
		gin.Recovery(),
	)

	// init new clients here and attach them to the server

	/**
	  * e.Use(func(c *gin.Context))  {
	*     attach clients
	*   }
	  * */

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "404 page not found"})
	})

	addr := fmt.Sprintf("%s:%s", config.Server.Address, config.Server.Port)
	log.Printf("Serving on %s...", addr)

	return e.Run(addr)
}
