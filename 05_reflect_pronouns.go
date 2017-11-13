//Author : Daire Ni Chathain
//Go exercise 5 :Adapt the function to reflect the pronouns in the captured groups, where necessary
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

// substitute pronouns in response
var substitutions = map[string]string{
	"you":    "I",
	"your":   "my",
	"you're": "I am",
	"me":     "you",
}

func elizaResponse(input string) string {

	//input = "i'm happy and I know it"
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
		//split every time there is a white space
		inputWords := strings.Split(feeling[1], " ")
		inputWords = strings.Split(feeling[1], " ")
		//check evrey word for a pronoun & swap on a match
		for i := range inputWords {
			if _, found := substitutions[inputWords[i]]; found {
				inputWords[i] = substitutions[inputWords[i]]
			}
		}
		// reassemble the sentence and adding space between each word
		feeling[1] = strings.Join(inputWords, " ")
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

		//remove full stops
		if last := len(userinput) - 1; last >= 0 && userinput[last] == '.' {
			userinput = userinput[:last]
		}
		// Generate and print Eliza's response.
		output := elizaResponse(strings.ToLower(userinput))
		fmt.Println(output) //print output of elizaresponse
	}
}
