const { start,end} = require('./time.js');
const coopMultiplier = 0.6; //co-op multi
var ogCoop = 5; //co-op multi
var win = 1; //for test


function coopScoresCalc() { //calcs score by how long an input takes
  var timeScore = 0.5;
    if (end() < 20){
      for(let i = 0; i < end(); i++){
        ogCoop = ogCoop*coopMultiplier;
     
      }
    }
    timeScore = timeScore + ogCoop 
    timeScore = Math.round(timeScore);
    return timeScore;
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
start();
for(let i = 0; i<300; i++){
    console.log("Yes");
}
end();
coopWin(win);
coopLoss();  

//exports
module.exports = {
  ogCoop,
  coopScoresCalc,
  coopWin,
  coopLoss
}