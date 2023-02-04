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

function findGame(channelId){
    return games.find(game => game.channelId == channelId);
}

module.exports = {
    games,
    newGame,
    findGame,
}