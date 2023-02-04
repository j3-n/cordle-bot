var fs = require("fs")
var path = require("path")

var WORD;
const guesses = [];

var char_values = {
    "Correct_Character": "Correct_Character",
    "Incorrect_Position":"Incorrect_Position",
    "Incorrect_Character":"Incorrect_Character"
}

function createGame()
{
    const wordList = importWords();
    WORD = wordList[Math.floor(Math.random()*wordList.length)]
    submitGuess("Josh");

}



function importWords(){
    const data = fs.readFileSync("./words.json", {encoding:"utf8", flag:"r"});
    const words = JSON.parse(data);
    return words;
}

// Submit guess and send back how good
// Struct contains array full of 
function submitGuess(guess)
{
    if(guesses.length > 6)
    {
        return null;
    }

    const toAdd = [];

    // The following stores each character of the guess under the key 'char' with the correctness of the character under "STATE"
    // The state is any value from char_values
    // the char is the individual characters from the guess
    for(let i = 0; i < guess.length; i++)
    {
        let key = guess.charAt(i);
        if(guess.charAt(i) == WORD.charAt(i))
        {
            console.log("Match! "+guess.charAt(i)+", at position: "+i);
            toAdd.push({char:key, STATE:char_values.Correct_Character});
        }
        else if(WORD.includes(guess.charAt(i)))
        {
            console.log("Incorrect position at index: "+i);
            toAdd.push({char:key, STATE:char_values.Incorrect_Position});
        }
        else
        {
            console.log("Character '"+guess.charAt(i)+"' Not in word")
            toAdd.push({char:key, STATE:char_values.Incorrect_Character});
        }

    }   
    
    // DEBUG PRINT STATEMENT
    // for(let i = 0; i < toAdd.length; i++)
    // {
    //     console.log(toAdd[i]);
    // } 
    guesses.push(toAdd);
    return toAdd;
}

// get remaining guesses

module.exports = {
    createGame,
    submitGuess
};

// Just for testing
// createGame();