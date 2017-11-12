//Author : Daire Ni Chathain
//Go exercise 7 : Guessing game with user input of guess & added comparison of cookie value
//Resources:

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Struct that holds the resonse array of strings
type Response struct {
	text []string
}

func elizaResponse(input string) string {

	// Create Resposne array filled with strings that rep responses

	responses := []string{"I'm not sure what you're trying to say. Could you explain it to me?", "How does that make you feel", "Why do you say that?"}
	//random num generator
	rand.Seed(time.Now().UTC().UnixNano())
	//Radnomly pick a response from the response array
	random := (rand.Intn(3))
	//return a random response
	output := (responses[random])
	return output
}
func main() {

	// Print a greeting to the user.
	fmt.Println("Hello, I'm Eliza. How are you feeling today?")

	inputs := []string{"People say I look like both my mother and father",
		"Father was a teacher.",
		"I was my father's favourite.",
		"I'm looking forward to the weekend.",
		"My grandfather was French!"}

	for i := range inputs {
		//get  user input from sample inpout array
		userinput := inputs[i]
		//Print user inout from sample input array
		fmt.Println(inputs[i])
		// Generate and print Eliza's response.
		output := elizaResponse(userinput)
		fmt.Println(output)

	}
	/*for user input
	// Keep reading user input and printing Eliza's response until the user types 'quit'.	for reader := bufio.NewReader(os.Stdin); ; {
	for reader := bufio.NewReader(os.Stdin); ; {
		// Print user prompt.
		fmt.Print("> ")
		// Read user input.
		userinput, _ := reader.ReadString('\n')
		// Trim the user input's end of line characters.
		userinput = strings.Trim(userinput, "\r\n")

		// Generate and print Eliza's response.
		//output := elizaResponse(userinput)
		fmt.Println(elizaResponse(userinput))
		//fmt.Println(eliza.analyse(userinput))

		// If the user input was quit, then quit.
		// Note that Eliza gets to respond to quit before this happens.
		if strings.Compare(strings.ToLower(strings.TrimSpace(userinput)), "quit") == 0 {
			break
		}
	}
	*/
}
