package wordle

import (
	"os"
	"encoding/json"

	"cordle/util"
)

// LoadWords reads the json files containing allowed words and answers and returns them (answers, guesses)
func LoadWords() ([]string, []string){
	afile, err := os.ReadFile("wordle/answers.json")
	util.CheckError(err, "Failed to load answers")
	gfile, err := os.ReadFile("wordle/guesses.json")
	util.CheckError(err, "Failed to load guesses")

	// Decode JSON
	answers := []string{}
	err = json.Unmarshal(afile, &answers)
	util.CheckError(err, "Failed to decode answers")

	guesses := []string{}
	err = json.Unmarshal(gfile, &guesses)
	util.CheckError(err, "Failed to load guesses")

	return answers, guesses
}