const { start,end} = require('./time.js');
const compMultiplier = 0.5;
var ogComp = 8; //comp multi
var compTimeScore =  0;

function compScoresCalc() { //calcs 
    if (end() < 20){
      for(let i = 0; i < end(); i++){
        ogComp = ogComp*compMultiplier;
     
      }
      compTimeScore += ogComp
    }
    return Math.round(compTimeScore);
} 

function compWin(turnWin){
    //boiler-plate
    
    const turns = [25,22,18,15,12,9];
    let winScore = turns[turnWin-1] 
    console.log(winScore+compScoresCalc());
    //return winScore + total comptime calc
   //calc elo rating 

}

function compLose(User){
   //I NEED THE DATABASE FOR THIS
}
  
//testers
/* for(let i = 0; i<2; i++){
    start();
    compScoresCalc();
    for(let j = 0; j<30000; j++){
        console.log(j);
    }
  
    end();
}

compWin(2);
   */


//exports
module.exports = {
    compWin,
    compLose,
    compMultiplier
} 