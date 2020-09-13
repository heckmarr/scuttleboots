package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh/terminal"
	"net/http"
	"time"
)
type Screen struct {
	initLines int
	initCols  int
	lines map[int]string
	cols map[int]string
}



//An important side effect of getting this to work is that
//the string variable can be replaced with a custom struct
func (s Screen) Init() Screen {
	cols, lines, err := terminal.GetSize(0)
	if err != nil {
		panic(err)
	}
	s.initLines = lines
	s.initCols = cols
	s.lines = make(map[int]string, s.initLines)

	s.cols = make(map[int]string, s.initCols)
	return s
}

func (s Screen) Fill(val string) Screen {

	for i := 0;i < s.initLines;i++ {

		for c:= 0;c < s.initCols;c++ {
			s.cols[c] = val

			s.lines[i] += s.cols[c]


		}
	}

	return s
}


func (s Screen) EditCell(x int, y int) Screen {

	return s
}

func main() {

	var screen Screen
	screen = screen.Init()
	screen = screen.Fill("#")
	fmt.Println("These are the screen dimensions")
	//They work on the current terminal, ie if you're
	//connected over ssh, you can get the dimensions
	//of the screen you're currently looking at
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
