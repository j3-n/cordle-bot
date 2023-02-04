const { WordleGame } = require("./src/wordle");

// Array of games
var games = []

function newGame(channelId){
    // If the channel does not have an active game, create it
    if(!games.find(game => game.channelId == channelId)){
        games.push({channelId: channelId, game: new WordleGame()});
        return true;
    }
    return false;
}

module.exports = {
    games,
    newGame,
}