const { start,end} = require('./time.js');
const { FirebaseFunctions } = require('./firebase/firebase-functions');
const compMultiplier = 0.5;
// var ogComp = 8; //comp multi
// var compTimeScore =  0;
const fbFunc = new FirebaseFunctions();
fbFunc.initialize();

// function compScoresCalc() { //calcs 
//     if (end() < 20){
//       for(let i = 0; i < end(); i++){
//         ogComp = ogComp*compMultiplier;
     
//       }
//       compTimeScore += ogComp
//     }
//     return Math.round(compTimeScore);
// } 

async function compWin(turnWin, userID){
    //boiler-plate
    const turns = [25,22,18,15,12,9];
    let winScore = turns[turnWin-1] 
    let eloScore = turns[turnWin-1] 

    const user = await fbFunc.getUser(userID, "users");
    let score = user.score;
    winScore += score;
    let eloScores = user.elo;
    eloScores += eloScore;
    fbFunc.updateUser(
        user.id,
        user.name,
        user.gamesWon,
        user.gamesLost,
        user.gamesPlayed,
        eloScores,
        winScore,
        "users"
    );
}

async function compLose(userID){
    
    let eloScore = -18;
    

    const user = await fbFunc.getUser(userID, "users");
    let elo= user.elo;
    if( elo + eloScore < 0){
        fbFunc.updateUser(
            user.id,
            user.name,
            user.gamesWon,
            user.gamesLost,
            user.gamesPlayed,
            0,
            user.score,
            "users"
        );
    }
    else{
    eloScore = elo + eloScore;
    fbFunc.updateUser(
        user.id,
        user.name,
        user.gamesWon,
        user.gamesLost,
        user.gamesPlayed,
        eloScore,
        user.score,
        "users"
    );
    }
}
//testers


//compWin(2,"natan#71912");
compLose("natan#71912");


//exports
module.exports = {
    compWin,
    compLose,
    compMultiplier
} 