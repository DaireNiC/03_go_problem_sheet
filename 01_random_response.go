//Author : Daire Ni Chathain
//Go exercise 7 : Guessing game with user input of guess & added comparison of cookie value
//Resources:

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Struct that holds the resonse array of strings
type Response struct {
	text []string
}

func main() {

	// Create Resposne array filled with strings that rep responses
	arr := []Response{
		{
			[]string{"I'm not sure what you're trying to say. Could you explain it to me?", "How does that make you feel", "Why do you say that?"},
		},
	}

	fmt.Println(arr)

	// Print a greeting to the user.
	fmt.Println("Hello, I'm Eliza. How are you feeling today?")

	// Keep reading user input and printing Eliza's response until the user types 'quit'.	for reader := bufio.NewReader(os.Stdin); ; {
	for reader := bufio.NewReader(os.Stdin); ; {
		// Print user prompt.
		fmt.Print("> ")
		// Read user input.
		userinput, _ := reader.ReadString('\n')
		// Trim the user input's end of line characters.
		userinput = strings.Trim(userinput, "\r\n")

		// Generate and print Eliza's response.
		//fmt.Println(eliza.analyse(userinput))
		fmt.Println(userinput)

		// If the user input was quit, then quit.
		// Note that Eliza gets to respond to quit before this happens.
		if strings.Compare(strings.ToLower(strings.TrimSpace(userinput)), "quit") == 0 {
			break
		}
	}
}

/*
func elizaResponse(string input) {
	fmt.Println(input)
}
*/
