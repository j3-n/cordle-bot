const { start,end} = require('./time.js');
const { FirebaseFunctions } = require('./firebase/firebase-functions');
const compMultiplier = 0.5;
var ogComp = 8; //comp multi
var compTimeScore =  0;
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
  

    const user = await fbFunc.getUser(userID, "users");
    let score = user.score;
    winScore = score + winScore;
    fbFunc.updateUser(
        user.id,
        user.name,
        user.gamesWon,
        user.gamesLost,
        user.gamesPlayed,
        user.elo,
        winScore,
        "users"
    );
}

async function compLose(userID){
    
    let loseScore = -18;
    

    const user = await fbFunc.getUser(userID, "users");
    let score = user.score;
    if( score + loseScore < 0){
        fbFunc.updateUser(
            user.id,
            user.name,
            user.gamesWon,
            user.gamesLost,
            user.gamesPlayed,
            user.elo,
            0,
            "users"
        );
    }
    else{
    loseScore = score + loseScore;
    fbFunc.updateUser(
        user.id,
        user.name,
        user.gamesWon,
        user.gamesLost,
        user.gamesPlayed,
        user.elo,
        loseScore,
        "users"
    );
    }
}
//testers


//compWin(2,"natan#71912");
//compLose("natan#71912");


//exports
module.exports = {
    compWin,
    compLose,
    compMultiplier
} 