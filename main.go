package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	// "musicstore"
)

func main() {
	helloMessage()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		// Read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle execution of input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func helloMessage() {
	helloMsg := "Welcome to the music store!\n"
	fmt.Println(helloMsg)

	instructions := `Please enter one of the following commands:
  1) addSong song artist genre
  2) deleteSong song
  3) listSongs`
	fmt.Println(instructions, "\n")
}

func goodbyeMessage() {
	goodbyeMsg := "Goodbye!\n"
	fmt.Println(goodbyeMsg)
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	fmt.Println("Args:", args)

	switch args[0] {
	case ("exit" || "quit"):
		goodbyeMessage()
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

//func callUserAction() {
//
//}
