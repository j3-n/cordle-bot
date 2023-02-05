const { DuelWordle } = require("./src/duel-game");

// Array of games
var games = []
// Array of pending challenges
var challenges = []

function newGame(player1, player2){
    // If the channel does not have an active game, create it
    if(!findGame(player1, player2)){
        games.push({player1: player1, player2: player2, game: new DuelWordle(player1, player2)});
        return true;
    }
    return false;
}

function findGame(player1, player2){
    let game = games.find(game => game.player1 == player1 && game.player2 == player2);
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

function findChallenge(player2){
    return challenges.find(challenge => challenge.player2 == player2);
}

function completeChallenge(player1, player2){
    challenges.splice(challenges.findIndex(challenge => challenge.player1 == player1 && challenge.player2 == player2), 1);
}

module.exports = {
    games,
    newGame,
    findGame,
    completeGame,
    newChallenge,
    findChallenge,
    completeChallenge,
}