const { start,end} = require('./time.js');
const coopMultiplier = 0.9; //co-op multi
var ogCoop = 5; //co-op multi
var compTimeScore =0 ;

function coopScoresCalc() { //calcs score by how long an input takes
    if (end() < 20){
      for(let i = 0; i < end(); i++){
        ogCoop = ogCoop*coopMultiplier;
     
      }
      compTimeScore += ogCoop 
    }
    compTimeScore = Math.round(compTimeScore);
    return compTimeScore;
} 
function coopWin(turnWin){ 
  //boiler-plate
  const turns = [6,5,4,3,2,1];
  let winScore = turns[turnWin-1] 
    
  console.log(coopScoresCalc()+winScore);
  //return winScore + total comptime calc
  //server input
}
function coopLoss(){ //needs to input stuff about user too!
  console.log(-coopScoresCalc());
  //server input
}

//tests

/* for(let i = 0; i<6; i++){
    start();
    console.log("Yes");
    end();
}

coopWin(win);
coopLoss();  */ 

//exports
module.exports = {
  ogCoop,
  coopScoresCalc,
  coopWin,
  coopLoss
}