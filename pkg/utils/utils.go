package utils

import (
	"bufio"
	"fmt"
	"os"
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
	var number int
	fmt.Println(prompt + "\n?> ")
	fmt.Scanln(&number)
	return number
}
