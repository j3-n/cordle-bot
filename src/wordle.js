var fs = require("fs")
var path = require("path")
// Create new game

// Submit new guesses
// Enum with guessed characters
// 

//process.stdout.write("Hello world!")

// Create game -> init function
function createGame()
{
    
}

function importWords()
{
    // Read file
    fs.readFile("./words.json", "utf8", function (err, data){
        if(err)
        {
            process.stdout.write(err.message);
            return;
        }
        try
        {
            // Parse json data into legible words and return
            const words = JSON.parse(data);
            return words;
        }catch (err)
        {
            process.stdout.write(err.message+"\n");
        }
    })
}

function submitGuess()
{

}

module.exports = {
    createGame,
    submitGuess
};

importWords();