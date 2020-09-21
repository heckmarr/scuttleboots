package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)
type Screen struct {
	initLines int
	initCols  int
	Points map[Cell]string
	Display string
}

type Missive struct {
	Header string
	Body string
	Footer string
	Recipient string
	Encrypted bool

}



type Input struct {
	Value string
	Prompt string
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
			lines := " ,.,0,4,%,^,&"
			//line := strings.Split(lines, ",")
			//s.Points[rand.Intn(pos.Y)]
			v := lines[rand.Intn(c)]
			pos.Value = string(v)
			s.Display += fmt.Sprint(v)
		}
		s.Display += fmt.Sprint("\n")
	}
	return s
}

func (s Screen) CreateShape() Screen {
	//This will just require an X and Y
	//dimensions as well as a position
	//on screen
	return s
}

func (s Screen) Scramble() Screen {
	var Print Cell

	lines := " ,.,0,4,%,^,&"
	line := strings.Split(lines, ",")
	//val := strings.Split(lines, ",")
	for c := 0;c < s.initCols;c++ {
		for l := 0;l < s.initLines;l++ {
			Print.Y = rand.Intn(l)
			Print.X = rand.Intn(c)
			value := line[rand.Intn(Print.Y)]
			Print.Value = fmt.Sprint("\x1b["+strconv.Itoa(Print.Y)+";"+strconv.Itoa(Print.X)+"H\x1b[38:2:0:200:0m"+value)

			fmt.Printf(Print.Value)

		}
	}

	fmt.Printf(Print.Value)

	return s
}
func (s Screen) EditCell() Screen {
	var Print Cell
	skan := bufio.NewScanner(os.Stdin)

	fmt.Printf("\x1b[38:2:50:100:25mEnter the X coodinate that you wish to change\x1b[0m")
	skan.Scan()
	val, err := strconv.Atoi(skan.Text())
	if err == nil {
		if val < s.initCols {
			//fmt.Println("Passes all tests")
			Print.X = val
		}
	}

	fmt.Printf("\x1b[38:2:50:100:25mEnter the Y coordinate that you wish to change\x1b[0m")
	skan.Scan()
	val, err = strconv.Atoi(skan.Text())
	if err == nil {
		if val < s.initLines {
			//fmt.Println("Passes all tests")
			Print.Y = val
		}
	}

	//This is kinda dumb and will change with create rectangle, triangle, hyperlink etc
	fmt.Println("\x1b[38:2:55:100:25mNow enter your message\x1b[0m")
	skan.Scan()
	var prompt Input
	prompt.Prompt = fmt.Sprint("\x1b[0:0H\x1b[38:2:200:100:50m<<<<-[")

	Print.Value = fmt.Sprint("\x1b["+strconv.Itoa(Print.Y)+";"+strconv.Itoa(Print.X)+"H\x1b[38:2:0:200:0m"+skan.Text())
	fmt.Printf(prompt.Prompt)
	fmt.Printf(Print.Value)

	return s
}

func (s Screen) FlipCell(val Cell) {

}

func main() {
	var prompt Input
	prompt.Prompt = fmt.Sprint("\x1b[0:0H\x1b[38:2:200:175:50m<<<<-[")
	fmt.Println("This will be the server")
	file, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(file)
	server := gin.Default()


	server.GET("/jack-in", RenderIntro)
	server.GET("/jack-out", TBA)

	server.POST("/observer", AlertRend)
	server.POST("/")


	s := &http.Server{
		Addr:           ":443",
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go s.ListenAndServeTLS("fullchain.pem", "privkey.pem")
	//This ends server stuff and starts display stuff

	//The following is all code needed to make a square the size of the terminal
	var scr Screen
	scr = scr.SpawnScreen()
	fmt.Print(scr.Display)

	box := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(scr.Display)
		fmt.Printf(prompt.Prompt)
		box.Scan()
		//remember to bluemonday this shit
		if box.Text() == "quit" || box.Text() == "exit" {
			os.Exit(1)
		}
		if box.Text() == "zealot" {
			scr.EditCell()
			fmt.Printf("\x1b[0m")
			box.Scan()
			continue
		}
		if box.Text() == "login" {
			for c := 0;c < scr.initCols;c++ {
				for r := 0;r < scr.initLines;r++ {
					scr.Scramble()
				}
			}
		}
		prompt.Value = box.Text()
		fmt.Printf(prompt.Value)

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
type changingCell struct {
	series []string
	progress int
	X int
	Y int
}

func RenderIntro(c *gin.Context) {
	//Put call to python3 crunching here
	c.String(http.StatusOK, "Whale oil beef hook")
	//Put golang serving said silliness here.
	cols, lines, err := terminal.GetSize(0)
	if err != nil {
		panic(err)
	}


	//total := make([]string, cols * lines)
	for i := 0; i < (cols); i++ {

		for c := 0;c < lines;c++ {
			go DoRender()
		}


	}
}
func DoRender() {


	cols, lines, err := terminal.GetSize(0)
	if err != nil {
		panic(err)
	}
	finished := make([]bool, cols)
	started := make([]bool, cols)
			cells := "^,&,*,<,>,$,#,@,!"
			values := strings.Split(cells, ",")
			var currentCell changingCell
			currentCell.series = values
			currentCell.progress = 0
			currentCell.X = rand.Intn(cols)
			currentCell.Y = rand.Intn(lines)



			/*
				var s Screen
				s = s.Init()
				s = s.Fill(" ")
			*/
			for {
				column := rand.Intn(cols)
				if currentCell.progress == len(currentCell.series)-1 {
					finished[column] = true
					/*	for i := 0; i < len(finished); i++ {
						if !finished[column] {
							continue
						} else {
							break
						}
					}*/
				}
				if !started[column] {
					//fmt.Println("HAVE WE STARTED")
					//Do the thing
					if currentCell.progress > 10 {
						finished[column] = true
					}
					if finished[column] == true {
						break
					}
					val := fmt.Sprint("\033[" + strconv.Itoa(currentCell.Y) + ";" + strconv.Itoa(currentCell.X) + "H" + currentCell.series[currentCell.progress])
					currentCell.progress++
					fmt.Printf(val)
					//Then mark that it has been done
					started[column] = true
				} else {
					//pick a new one
					//Continue gracefully exits and restarts the loop
					continue
				}
				time.Sleep(250*time.Millisecond)
			}

}





