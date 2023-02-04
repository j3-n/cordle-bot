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
    let game = games.find(game => game.channelId == channelId);
    if(game)
        return game.game;
    else
        return null;
}

function completeGame(channelId){
    // Remove the game from the list
    games.splice(games.findIndex(game => game.channelId == channelId), 1);
}

module.exports = {
    games,
    newGame,
    findGame,
    completeGame,
}