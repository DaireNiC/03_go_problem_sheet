//Author : Daire Ni Chathain
//Go exercise 2 : Adapt the ElizaResponse function to, if the input does not contain the word "father",
//check the input begins with "I am ". If it does, use a regular expression to capture the rest of the user input.
//Return the string "How do you know you are _?", replacing the underscore with the captured part of the input.
//Add, along with the five previous test inputs, the following test inputs.
//Resources:https://golang.org/pkg/regexp/
//http://blog.kamilkisiel.net/blog/2012/07/05/using-the-go-regexp-package/
//https://stackoverflow.com/questions/30957615/regex-to-match-variations-of-i-am-im-im-iam-i-am

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
	input = "i'm happy"
	//Regular expression to match on "father"
	regFather := regexp.MustCompile("(?i)\\bfather\\b")
	if regFather.MatchString(input) {
		return "Why don’t you tell me more about your father?"
	}
	//reg expression to check for "I am and variants e.g I'm "
	regIam := regexp.MustCompile(`\b(?i)I'?\s*a?m(.*)`)
	feeling := regIam.FindStringSubmatch(input)
	//return everything after the regEx match
	if len(feeling) > 1 {
		response := "How do you know you are%s?"
		return fmt.Sprintf(response, feeling[len(feeling)-1])
	} else {
		//no match will return a random response
		//random num generator
		rand.Seed(time.Now().UTC().UnixNano())
		//Radnomly pick a response from the response array
		random := (rand.Intn(3))
		//return a random response
		return (responses[random])
	}

}
func main() {

	//test inputs provided
	inputs := []string{"People say I look like both my mother and father",
		"Father was a teacher.",
		"I am my father's favourite.",
		"I'm looking forward to the weekend.",
		"My grandfather was French!",
		"I am happy.",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you're saying at face value"}
	//test each input
	for i := range inputs {
		//get  user input from sample inpout array
		userinput := inputs[i]
		fmt.Println(userinput) //Print from sample input array
		// Generate and print Eliza's response.
		output := elizaResponse(strings.ToLower(userinput))
		fmt.Println(output) //print output of elizaresponse
	}
}
