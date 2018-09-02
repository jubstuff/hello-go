package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jubstuff/hello-go/pkg/utils"
)

func hello() {
	var name string
	fmt.Print("What's your name? ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s, nice to meet you\n", name)
}

func countingCharacters() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the input string? ")
	var text string
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%s contains %d characters\n", text, len(text))
}

func printingQuotes() {
	var quote, author string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the quote?")
	quote, _ = reader.ReadString('\n')
	quote = strings.TrimSpace(quote)
	fmt.Print("Who said it?")
	author, _ = reader.ReadString('\n')
	author = strings.TrimSpace(author)

	sentence := author + " says, \"" + quote + "\""
	fmt.Println(sentence)
}

func madLib() {
	noun := utils.GetLine("Enter a noun")
	verb := utils.GetLine("Enter a verb")
	adjective := utils.GetLine("Enter an adjective")
	adverb := utils.GetLine("Enter an adverb")

	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adjective, noun, adverb)
}

func simpleMath() {
	firstNumber := utils.GetNumber("What's the first number?")
	secondNumber := utils.GetNumber("What's the second number?")

	fmt.Printf("%d + %d = %d\n", firstNumber, secondNumber, firstNumber+secondNumber)
	fmt.Printf("%d - %d = %d\n", firstNumber, secondNumber, firstNumber-secondNumber)
	fmt.Printf("%d * %d = %d\n", firstNumber, secondNumber, firstNumber*secondNumber)
	fmt.Printf("%d / %d = %d\n", firstNumber, secondNumber, firstNumber/secondNumber)
}

func areaRectangularRoom() {
	const ConversionFactor = 0.092903

	length := utils.GetNumber("What is the length of the room in feet?")
	width := utils.GetNumber("What is the width of the room in feet?")

	areaFeet := float64(length) * float64(width)
	areaMeters := float64(areaFeet) * ConversionFactor

	fmt.Printf("You entered dimension of %d feet by %d feet\n", length, width)
	fmt.Printf("The area is %f square feet\n", areaFeet)
	fmt.Printf("That are %f square meters\n", areaMeters)
}

func menu() {
	fmt.Println("Learning Go")
	fmt.Println("0. Exit")
	fmt.Println("1. Hello")
	fmt.Println("2. Counting the number of characters")
	fmt.Println("3. Printing quotes")
	fmt.Println("4. Mad Lib")
	fmt.Println("5. Simple Math")

}

func main() {
	var choice int
	menu()
	fmt.Print("?> ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		hello()
	case 2:
		countingCharacters()
	case 3:
		printingQuotes()
	case 4:
		madLib()
	case 5:
		simpleMath()
	case 6:
		areaRectangularRoom()
	default:
		fmt.Println("Bye Bye")
	}

}
