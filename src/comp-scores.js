const { FirebaseFunctions } = require('./firebase/firebase-functions');

async function compWin(turnWin, userID){
    //boiler-plate
    const turns = [25,22,18,15,12,9];
    var winScore = turns[turnWin-1] 
    var eloScore = turns[turnWin-1] 

    const user = await FirebaseFunctions.getUser(userID, "users");
    var score = user.score;
    winScore += score;
    var eloScores = user.elo;
    eloScores += eloScore;

    return {
        id: user.id,
        name: user.name,
        gamesWon: user.gamesWon+1,
        gamesLost: user.gamesLost,
        gamesPlayed: user.gamesPlayed+1,
        elo: eloScores,
        score: winScore,
        collection: "users"
    };
}

async function compLose(turnLose, userID){    
    const turns = [9,12,15,18,22,25];
    var loseScore = turns[turnWin-1] 

    const user = await FirebaseFunctions.getUser(userID, "users");
    var elo = user.elo;
   
    if( elo -= loseScore < 0){
        return {
            id: user.id,
            name: user.name,
            gamesWon: user.gamesWon,
            gamesLost: user.gamesLost,
            gamesPlayed: user.gamesPlayed,
            elo: 0,
            score: user.score,
            collection: "users"
        };
    } else {
        elo -= loseScore;
        return {
            id: user.id,
            name: user.name,
            gamesWon: user.gamesWon,
            gamesLost: user.gamesLost+1,
            gamesPlayed: user.gamesPlayed+1,
            elo: elo,
            score: user.score,
            collection: "users"
        };
    }
}
async function compDraw(userID){
    var drawScore = 8;
    var drawScores = user.score;
    drawScores += drawScore;

    const user = await FirebaseFunctions.getUser(userID, "users");   

    return {
        id: user.id,
        name: user.name,
        gamesWon: user.gamesWon,
        gamesLost: user.gamesLost,
        gamesPlayed: user.gamesPlayed+1,
        elo: user.elo,
        score: drawScores,
        collection: "users"
    };
}
//exports
module.exports = {
    compWin,
    compLose,
    compDraw
} 