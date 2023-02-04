var fs = require("fs")
var path = require("path")

// For testing user input
// const prompt = require("prompt-sync")({sigint: true});

// TODO:
// KEEP TRACK OF USED CHARACTERS and create getters for them
// CREATE GETTER FOR PAST ATTEMPTS

var WORD;
const guesses = [];

var char_values = {
    "correct_character": "correct_character",
    "incorrect_position":"incorrect_position",
    "incorrect_character":"incorrect_character",
    "correct_guess":"correct_guess"
}

function createGame()
{
    const wordList = importWords();
    WORD = wordList[Math.floor(Math.random()*wordList.length)]
    // submitGuess(WORD);

}

// function testPlay()
// {
//     createGame();

//     let shouldPlay = true;
//     let i = 0;
//     while(shouldPlay)
//     {
//         var userGuess = prompt("Enter guess: ");
//         var result = submitGuess(userGuess);

//         // If null then maximum guesses acheived
//         if(guesses.length > 5)
//         {
//             shouldPlay = false;
//             console.log("Ran out of guesses!")
            
//         }// check for correct guess
//         else if(result.STATE == char_values.correct_guess)
//         {
//             shouldPlay = false;
//             console.log("Correct answer!");
//         }
        
//     }

// }

function importWords(){
    const data = fs.readFileSync("./words.json", {encoding:"utf8", flag:"r"});
    const words = JSON.parse(data);
    return words;
}

 
function submitGuess(guess)
{
    //console.log(WORD);

    let toAdd = [];

    // The following stores each character of the guess under the key 'char' with the correctness of the character under "STATE"
    // The state is any value from char_values
    // the char is the individual characters from the guess
    // Structure example: {CHAR: 'KEY', STATE: 'STATE FROM char_values' }
    // NOTE: CHAR: null WHEN STATE IS EQUAL TO correct_guess. THIS IS WHEN THE GUESS WAS CORRECT
    let isCorrect = true;
    for(let i = 0; i < guess.length; i++)
    {
        let key = guess.charAt(i);
        if(guess.charAt(i) == WORD.charAt(i))
        {
            //console.log("Match! "+guess.charAt(i)+", at position: "+i);
            toAdd.push({CHAR:key, STATE:char_values.correct_character});
        }
        else if(WORD.includes(guess.charAt(i)))
        {
            //console.log("Incorrect position at index: "+i);
            toAdd.push({CHAR:key, STATE:char_values.incorrect_position});
            isCorrect = false;
        }
        else
        {
            //console.log("Character '"+guess.charAt(i)+"' Not in word")
            toAdd.push({CHAR:key, STATE:char_values.incorrect_character});
            isCorrect = false;
        }

    }

    if(isCorrect)
    {
        toAdd = [];
        toAdd.push({CHAR:null, STATE:char_values.correct_guess});
    }
    
    // ENABLE THESE PRINT STATEMENTS TO SEE CORE DATA
    // for(let i = 0; i < toAdd.length; i++)
    // {
    //     console.log(toAdd[i]);
    // } 
    
    guesses.push({CORRECT:isCorrect, DATA:toAdd});

    // ENABLE THESE PRINT STATEMENTS TO SEE DATA INSIDE GUESSES
    // for(let i = 0; i < guesses.length; i++)
    // {
    //     console.log(guesses[i]);
    // } 

    return toAdd;
}

// get remaining guesses

module.exports = {
    createGame,
    submitGuess
};

// Just for testing
// testPlay();