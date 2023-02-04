const { WordleGame } = require("./wordle");

const Conditions = {
    OUT_OF_GUESSES: "OUT_OF_GUESSES",
    INVALID_ID: "INVALID_ID",
}

// Two players with individual guesses
// supply user ID to decide whos game to choose

class DuelGame extends WordleGame{
    constructor(playerID){
        super();
        this.playerID = playerID;
    }

} 

class DuelWordle
{
    constructor(PlayerOneID, PlayerTwoID)
    {
        this.PlayerOne = new DuelGame(PlayerOneID);
        this.PlayerTwo = new DuelGame(PlayerTwoID);
    }

    submitGuess(PlayerID, guess)
    {
        if(Object.is(PlayerID, this.PlayerOne.playerID))
            return this.PlayerOne.submitGuess(guess);
        else if(Object.is(PlayerID, this.PlayerTwo.playerID))
            return this.PlayerTwo.submitGuess(guess);
    }
}


// function testShit()
// {
//     const dw = new DuelWordle(1, 2);
//     console.log(dw.submitGuess(1, "Donke"));

// }

// testShit();

module.exports = {
    DuelWordle,
    Conditions,
}