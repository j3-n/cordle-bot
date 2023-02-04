var fs = require("fs")
var path = require("path")

// For testing user input
const prompt = require("prompt-sync")({sigint: true});

// TODO:
// KEEP TRACK OF USED CHARACTERS and create getters for them
// CREATE GETTER FOR PAST ATTEMPTS
// INPUT FILTERING = DONE

const Result = {
    CORRECT_CHARACTER: "correct_character",
    INCORRECT_POSITION: "incorrect_position",
    INCORRECT_CHARACTER: "incorrect_character",
};

const wordList = importWords();

var incorrectCharacters = new Set();

class WordleGame{
    constructor(){
        // Choose a random word
        this.word = wordList[Math.floor(Math.random() * wordList.length)]
        this.guesses = []
    }

    submitGuess(guess){
        guess.toLowerCase();
        // Input filtering
        if(!this.checkInput(guess)){
            return null; // NULL RETURN WHEN INPUT INVALID
        }
        
        let toAdd = [];
        // The following stores each character of the guess under the key 'char' with the correctness of the character under "STATE"
        // The state is any value from char_values
        // the char is the individual characters from the guess
        // Structure example: {CHAR: 'KEY', STATE: 'STATE FROM char_values' }
        // NOTE: CHAR: null WHEN STATE IS EQUAL TO correct_guess. THIS IS WHEN THE GUESS WAS CORRECT
        for(let i = 0; i < guess.length; i++){
            let key = guess.charAt(i); // unused?
            let result;
            
            if(guess.charAt(i) == this.word.charAt(i))
                result = Result.CORRECT_CHARACTER;
            else if(this.word.includes(guess.charAt(i)))
                result = Result.INCORRECT_POSITION;
            else{
                result = Result.INCORRECT_CHARACTER;
                // Add incorrect characters to used incorrect set
                incorrectCharacters.add(guess.charAt(i));
            }
            
            
            toAdd.push({char: key, result: result});
        }
        
        let result = ({
            correct: toAdd.every(r => r.result == Result.CORRECT_CHARACTER), 
            guess: toAdd,
        });

        this.guesses.push(result);

        return result;
    }

    // checks for input being 5 characters long and no numbers only alpha characters
    checkInput(input){
        if(input.length != 5)
            return false;
        for(let i = 0; i < input.length; i++)
        {
            if(input.charCodeAt(i) < 97 && input.charCodeAt(i) > 122)
                return false;
        }
        return true;
    }

    getNumberOfAttempts()
    {
        return this.guesses.length;
    }

    getIncorrectCharactersSet()
    {
        return this.incorrectCharacters;
    }
}

function importWords(){
    let data = fs.readFileSync("./words.json", {encoding:"utf8", flag:"r"});
    return JSON.parse(data);
}

// TESTING ONLY
function testPlay()
{
    let game = new WordleGame();
    console.log(game.word);

    let shouldPlay = true;
    let i = 0;
    while(shouldPlay)
    {
        console.log("Used characters:"+ incorrectCharacters.size)
        for(const item of incorrectCharacters)
        {
            console.log(item);
        }
        let userGuess = prompt("Enter guess: ");
        let result = game.submitGuess(userGuess);
        console.log(result);

        if(result == null)
        {
            console.log("Invalid input!");
            shouldPlay = false;
            return;
        }
        // If null then maximum guesses acheived
        if(game.guesses.length > 5)
        {
            shouldPlay = false;
            console.log("Ran out of guesses!")
            
        }// check for correct guess
        else if(result.correct)
        {
            shouldPlay = false;
            console.log("Correct answer!");
        }
        
    }

}

module.exports = {
    WordleGame,
    Result,
};

// TESTING ONLY
testPlay();