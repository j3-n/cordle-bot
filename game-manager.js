const { WordleGame } = require("./src/wordle");

// Array of games
var games = []
// Array of pending challenges
var challenges = []

function newGame(channelId, newGame){
    // If the channel does not have an active game, create it
    if(!games.find(game => game.channelId == channelId)){
        games.push({channelId: channelId, game: newGame});
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

function newChallenge(player1, player2){
    if(!challenges.find(challenge => challenge.player1 == player1 || challenge.player2 == player2)){
        challenges.push({player1: player1, player2: player2});
        return true;
    }
    return false;
}

module.exports = {
    games,
    newGame,
    findGame,
    completeGame,
    newChallenge,
}