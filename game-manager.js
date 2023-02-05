const { DuelWordle } = require("./src/duel-game");

// Array of games
var games = [];
// Array of pending challenges
var challenges = [];

function newGame(threadId, player1, player2){
    games.push({player1Id: player1, player2Id: player2, game: new DuelWordle(player1, player2), threadId: threadId});
}

function findGame(player){
    let game = games.find(game => game.player1Id == player || game.player2Id == player);
    if(game)
        return game.game;
    else
        return null;
}

function findGameByChannelId(channelId){
    return games.find(game => game.threadId == channelId);
}

function gameExists(player){
    return findGame(player) ? true : false;
}

function completeGame(threadId){
    // Remove the game from the list
    games.splice(games.findIndex(game => game.threadId == threadId), 1);
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
    findGameByChannelId,
    gameExists,
    completeGame,
    newChallenge,
    findChallenge,
    completeChallenge,
}