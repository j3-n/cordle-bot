const { WordleGame } = require("./src/wordle");

// Array of games
var games = []

function newGame(channelId){
    games.push({channelId: channelId, game: new WordleGame()});
}

module.exports = {
    games,
    newGame,
}