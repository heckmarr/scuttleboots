package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	fmt.Println("This will be the server")
	server := gin.Default()
	server.GET("/screen", RenderScreen)
	server.Run(":80")
	for {
		time.Sleep(100*time.Millisecond)
	}
}

func RenderScreen(c *gin.Context) {
	//Put call to python3 crunching here
	c.String(http.StatusOK, "Whale oil beef hook")
}
