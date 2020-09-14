package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh/terminal"
	"net/http"
	"os"
	"strconv"
	"time"
)
type Screen struct {
	initLines int
	initCols  int
	Points map[Cell]string
	Display string
}

type Cell struct {
	X int
	Y int
	Value string
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
	l := make(map[Cell]string, cols * lines)

	s.Points = l

	s.Display = fmt.Sprintf("\033[0:0H")
	return s
}

func (s Screen) Fill(val string) Screen {
	for i := 0;i < s.initLines;i++ {
		for c := 0;c < s.initCols;c++ {
			pos := Cell{c, i, "boot"}
			s.Points[pos] = "#"
			s.Display += fmt.Sprint(s.Points[pos])
		}
		s.Display += fmt.Sprint("\n")
	}
	return s
}

func (s Screen) EditCell() Screen {
	var Print Cell
	skan := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter the X coodinate that you wish to change")
	skan.Scan()
	val, err := strconv.Atoi(skan.Text())
	if err == nil {
		if val < s.initCols {
			fmt.Println("Passes all tests")
			Print.X = val
		}
	}

	fmt.Printf("Enter the Y coordinate that you wish to change")
	skan.Scan()
	val, err = strconv.Atoi(skan.Text())
	if err == nil {
		if val < s.initLines {
			fmt.Println("Passes all tests")
			Print.Y = val
		}
	}

	//This is kinda dumb and will change with create rectangle, triangle, hyperlink etc
	fmt.Println("Now enter your message")
	skan.Scan()
	Print.Value = skan.Text()
	return s
}

func main() {

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
	//This ends server stuff and starts display stuff

	//The following is all code needed to make a square the size of the terminal
	var scr Screen
	scr = scr.SpawnScreen()
	fmt.Print(scr.Display)


	for {
		time.Sleep(100*time.Millisecond)
	}
}

func (s Screen) SpawnScreen() Screen {

	var screen Screen
	screen = screen.Init()
	screen = screen.Fill("#")
	fmt.Println("These are the screen dimensions")
	//Init calculates the size of the current terminal, ie if you're
	//connected over ssh, you can get the dimensions
	//of the screen you're currently looking at
	return screen
}

func TBA(c *gin.Context) {

}

func AlertRend(c *gin.Context) {
	_, err := os.Stat(".blit")
	if err == nil {
		fmt.Println("File exists, continue with blitting.")
		//Remembering that ddos attacks exists
		//And that to ustilize gin, we need to pass a value
	}else {
		file, err := os.Create(".blit")
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	//This will spawn a file that
	//indicates to the front end
	//That it needs to update the screen
	//
	_, err = os.Stat(".proc")
	if err == nil {
		fmt.Println("Processing is in progress")
	}else {
		file, err := os.Create(".proc")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		//This is where the main logic of compiling the "View"
		//will go

		//WHEN COMPLETE//
		//Clean all tokens
		_, err = os.Stat(".blit")
		if err == nil {
			os.Remove(".blit")
		}
		_, err = os.Stat(".proc")
		if err == nil {
			os.Remove(".proc")
		}

	}


	//When that happens the alert token
	//will be briefly replaced with a
	//"processing" bit
	//
	//Then, once the mechanism has drawn all
	//requested, a final token is created
	//and the server serves it it

	//the tokens will probably yaml, unless toml
	//or json are more viable

}

func Render(s Screen) {

}


func RenderIntro(c *gin.Context) {
	//Put call to python3 crunching here
	c.String(http.StatusOK, "Whale oil beef hook")
	//Put golang serving said silliness here.

}
