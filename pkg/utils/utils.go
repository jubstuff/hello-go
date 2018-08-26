package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetLine gets an input from the user, given a prompt
func GetLine(prompt string) string {
	var line string
	fmt.Println(prompt + "\n?> ")
	reader := bufio.NewReader(os.Stdin)
	line, _ = reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// GetNumber returns a number
func GetNumber(prompt string) int {
	var input string
	var number int
	var err error

	isOk := false

	for !isOk {
		fmt.Println(prompt + "\n?> ")
		fmt.Scanln(&input)
		number, err = strconv.Atoi(input)
		if err == nil {
			isOk = true
		} else {
			fmt.Printf("%q doesn't look like a number. Please insert a number.\n", input)
		}

	}

	return number
}
