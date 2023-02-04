const { WordleGame } = require("./wordle");

const Conditions = {
    OUT_OF_GUESSES: "OUT_OF_GUESSES",
    INVALID_ID: "INVALID_ID",
    PLAYER_ONE_WIN: "PLAYER_ONE_WIN",
    PLAYER_TWO_WIN: "PLAYER_TWO_WIN",
}

// Two players with individual guesses
// supply user ID to decide whos game to choose

class DuelGame extends WordleGame{
    constructor(playerID){
        super();
        this.playerID = playerID;
    }

} 

// Time -> if player runs out of time they loose
// 

class DuelWordle{
    constructor(PlayerOneID, PlayerTwoID){
        
        this.PlayerOne = new DuelGame(PlayerOneID);
        this.PlayerTwo = new DuelGame(PlayerTwoID);
        this.PlayerTwo.word = this.PlayerOne.word;
    }

    submitGuess(PlayerID, guess)
    {
        if(Object.is(PlayerID, this.PlayerOne.playerID)){
            let result = this.PlayerOne.submitGuess(guess);

            if(result.correct)
                return Conditions.PLAYER_ONE_WIN;

            return result;
        }
        else if(Object.is(PlayerID, this.PlayerTwo.playerID)){
            let result = this.PlayerTwo.submitGuess(guess);

            if(result.correct)
                return Conditions.PLAYER_TWO_WIN;


            return result;
        }
    }

    getNumberOfAttempts(PlayerID)
    {
        if(Object.is(PlayerID, this.PlayerOne.playerID))
            return this.PlayerOne.guesses.length;
        else if(Object.is(PlayerID, this.PlayerOne.playerID))
            return this.PlayerTwo.guesses.length;
        return Conditions.INVALID_ID;
    }
}


// function testShit()
// {
    
//     const dw = new DuelWordle(1, 2);
//     console.log(dw.PlayerOne.word);
//     console.log(dw.PlayerTwo.word);

//     console.log(dw.submitGuess(1, "aoiue"));

// }

// testShit();

module.exports = {
    DuelWordle,
    Conditions,
}