//Author : Daire Ni Chathain
//Go exercise 2 : Adapt the ElizaResponse function to, if the input does not contain the word "father",
//check the input begins with "I am ". If it does, use a regular expression to capture the rest of the user input.
//Return the string "How do you know you are _?", replacing the underscore with the captured part of the input.
//Add, along with the five previous test inputs, the following test inputs.
//Resources:https://golang.org/pkg/regexp/

package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Create Resposne array filled with strings that rep responses
var responses = []string{"I'm not sure what you're trying to say. Could you explain it to me?",
	"How does that make you feel",
	"Why do you say that?"}

func elizaResponse(input string) string {
	//Regular expression to match on "father"
	//return father response
	reg := regexp.MustCompile("(?i)\\bfather\\b")
	if reg.MatchString(input) {
		return "Why don’t you tell me more about your father?"
	}

	//reg expression to check for "I am"
	regIam := regexp.MustCompile("(?i)\\bI am\\b")

	if regIam.MatchString(input) {
		return "How do you know you are ?"
	}

	// if the input string contains the word 'father' return the specific string
	if reg.MatchString(input) {
		return "Why don’t you tell me more about your father?"
	} else {
		//random num generator
		rand.Seed(time.Now().UTC().UnixNano())
		//Radnomly pick a response from the response array
		random := (rand.Intn(3))
		//return a random response
		return (responses[random])
	}
}
func main() {

	// Print a greeting to the user.
	//fmt.Println("Hello, I'm Eliza. How are you feeling today?")

	inputs := []string{"People say I look like both my mother and father",
		"Father was a teacher.",
		"I am my father's favourite.",
		"I'm looking forward to the weekend.",
		"My grandfather was French!"}

	for i := range inputs {
		//get  user input from sample inpout array
		userinput := inputs[i]
		fmt.Println(userinput) //Print from sample input array
		// Generate and print Eliza's response.
		output := elizaResponse(strings.ToLower(userinput))

		fmt.Println(output) //print output of elizaresponse

	}
}
