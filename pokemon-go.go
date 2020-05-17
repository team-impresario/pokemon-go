package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	parser "pokemon-go/parser"
)

var usage = `
______   ____ |  | __ ____   _____   ____   ____             ____   ____
\____ \ /  _ \|  |/ // __ \ /     \ /  _ \ /    \   ______  / ___\ /  _ \
|  |_> >  <_> )    <\  ___/|  Y Y  (  <_> )   |  \ /_____/ / /_/  >  <_> )
|   __/ \____/|__|_ \\___  >__|_|  /\____/|___|  /         \___  / \____/
|__|               \/    \/      \/            \/         /_____/


pokemon-go is a great new programming language -- one of the best, in fact.

Usage:

    pokemon-go exeggcute <path to a .pkgo script>


There are no other commands.

Thank you for using pokemon-go.
`

func exeggcute(filePath string) {
	reader, _ := os.Open(filePath)
	parser := parser.NewParser(reader)
	result := parser.Parse()

	// fmt.Println(result)
	file, _ := ioutil.TempFile("", "script*.go")
	// fmt.Println(file.Name())
	defer os.Remove(file.Name())

	err := ioutil.WriteFile(file.Name(), []byte(result), 0644)

	cmnd := exec.Command("go", "run", file.Name())
	output, err := cmnd.Output()
	if err != nil {
		fmt.Println("Your code did not work. Have you considered quitting programming?")
		fmt.Println("Thank you for using pokemon-go!")
	}
	fmt.Printf("%s", output)
}

func handleExeggcute() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a path to a valid .pkgo script.")
		return
	} else {
		_, err := os.Stat(os.Args[2])
		if os.IsNotExist(err) {
			fmt.Printf("Could not find .pkgo script named \"%s\"\n", os.Args[2])
			return
		}
		exeggcute(os.Args[2])

	}
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println(usage)
	} else {
		switch os.Args[1] {
		case "exeggcute":
			handleExeggcute()
		default:
			fmt.Println(usage)

		}
	}
}
