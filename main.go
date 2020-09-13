package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh/terminal"
	"net/http"
	"time"
)
type Screen struct {
	lines int
	cols int
}


func main() {

	screen := DetermineWinSize()
	fmt.Println("These are the screen dimensions")
	fmt.Println(screen)

	fmt.Println("This will be the server")
	server := gin.Default()
	server.GET("/jack-in", RenderIntro)
	server.GET("/jack-out", TBA)
	server.GET("/observer", AlertRend)
	server.POST("/")


	s := &http.Server{
		Addr:           ":443",
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServeTLS("fullchain.pem", "privkey.pem")
	for {
		time.Sleep(100*time.Millisecond)
	}
}

func DetermineWinSize() (Screen) {
	//Initilialize the screen
	var newScreen Screen
	cols, lines, err := terminal.GetSize(0)
	if err != nil {
		panic(err)
	}
	newScreen.cols = cols
	newScreen.lines = lines

	return newScreen
}


func TBA(c *gin.Context) {

}

func AlertRend(c *gin.Context) {
	//This will spawn a file that
	//indicates to the python front end
	//That it needs to update the screen
	//
	//When that happens the alert token
	//will be briefly replaced with a
	//"processing" bit
	//
	//Then, once the python has drawn all
	//requested, a final token is created
	//and the golang displays it

	//the tokens will probaly yaml, unless toml
	//or json are more viable

}


func RenderIntro(c *gin.Context) {
	//Put call to python3 crunching here
	c.String(http.StatusOK, "Whale oil beef hook")
	//Put golang serving said silliness here.

}
