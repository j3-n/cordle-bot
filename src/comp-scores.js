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
        gamesPlayed: user.gamesWon,
        gamesLost: user.gamesLost,
        gamesPlayed: user.gamesPlayed,
        elo: eloScores,
        score: winScore,
        collection: "users"
    };
}

async function compLose(userID){    
    let eloScore = -18;

    const user = await FirebaseFunctions.getUser(userID, "users");
    var elo = user.elo;

    if( elo + eloScore < 0){
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
        eloScore = elo + eloScore;
        return {
            id: user.id,
            name: user.name,
            gamesWon: user.gamesWon,
            gamesLost: user.gamesLost,
            gamesPlayed: user.gamesPlayed,
            elo: eloScore,
            score: user.score,
            collection: "users"
        };
    }
}

//exports
module.exports = {
    compWin,
    compLose
} 