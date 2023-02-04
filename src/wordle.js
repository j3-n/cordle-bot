const { GuildDefaultMessageNotifications } = require("discord.js");
var fs = require("fs")
var path = require("path")

// For testing user input
const prompt = require("prompt-sync")({sigint: true});

// TODO:
// KEEP TRACK OF USED CHARACTERS and create getters for them
// CREATE GETTER FOR PAST ATTEMPTS

const Result = {
    CORRECT_CHARACTER: "correct_character",
    INCORRECT_POSITION: "incorrect_position",
    INCORRECT_CHARACTER: "incorrect_character",
};

const wordList = importWords();

class WordleGame{
    constructor(){
        // Choose a random word
        this.word = wordList[Math.floor(Math.random() * wordList.length)]
        this.guesses = []
    }

    submitGuess(guess){
        let toAdd = [];
        // The following stores each character of the guess under the key 'char' with the correctness of the character under "STATE"
        // The state is any value from char_values
        // the char is the individual characters from the guess
        // Structure example: {CHAR: 'KEY', STATE: 'STATE FROM char_values' }
        // NOTE: CHAR: null WHEN STATE IS EQUAL TO correct_guess. THIS IS WHEN THE GUESS WAS CORRECT
        for(let i = 0; i < guess.length; i++){
            let key = guess.charAt(i);
            let result = Result.INCORRECT_CHARACTER;

            if(guess.charAt(i) == this.word.charAt(i))
                result = Result.CORRECT_CHARACTER;
            else if(this.word.includes(guess.charAt(i)))
                result = Result.INCORRECT_POSITION;
            
            toAdd.push({char: key, result: result});
        }
        
        let result = ({
            correct: toAdd.every(r => r.result == Result.CORRECT_CHARACTER), 
            guess: toAdd,
        });

        this.guesses.push(result);

        return result;
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
        let userGuess = prompt("Enter guess: ");
        let result = game.submitGuess(userGuess);
        console.log(result);

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
//testPlay();