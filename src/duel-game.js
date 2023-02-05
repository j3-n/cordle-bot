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
    constructor(playerId){
        super();
        this.playerId = playerId;
        this.guessing = true;
        setTimeout(playerGuessingOutOfTime() ,30000);
    }

} 

// Time -> if player runs out of time they loose
// 

class DuelWordle{
    constructor(player1, player2){
        
        this.player1= new DuelGame(player1);
        this.player2 = new DuelGame(player2);
        this.player2.word = this.player1.word;
    }

    submitGuess(playerId, guess)
    {
        if(Object.is(playerId, this.player1.playerId)){
            if(!this.player1.hasRemainingAttempts)
                return Conditions.OUT_OF_GUESSES;

            let result = this.player1.submitGuess(guess);

            if(result.correct)
                return Conditions.PLAYER_ONE_WIN;

            return result;
        }
        else if(Object.is(playerId, this.player2.playerId)){
            if(!this.player1.hasRemainingAttempts)
                return Conditions.OUT_OF_GUESSES;
            let result = this.player2.submitGuess(guess);

            if(result.correct)
                return Conditions.PLAYER_TWO_WIN;

            return result;
        }
        else
            return Conditions.INVALID_ID;
    }

    getNumberOfAttempts(playerId)
    {
        if(Object.is(playerId, this.player1.playerId))
            return this.player1.guesses.length;
        else if(Object.is(playerId, this.player1.playerId))
            return this.player2.guesses.length;
        return Conditions.INVALID_ID;
    }

    isPlayerOutOftime()
    {

    }

    playerGuessingOutOfTime()
    {
        
    }
}


function testShit()
{
    
    const dw = new DuelWordle(1, 2);
    console.log(dw.player1.word);
    console.log(dw.player2.word);

    console.log(dw.submitGuess(1, "aoiue"));

}

testShit();

module.exports = {
    DuelWordle,
    Conditions,
}