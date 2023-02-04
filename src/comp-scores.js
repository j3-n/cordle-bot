const compMultiplier = 0.8;
var startTime, endTime;
var ogComp = 8; //comp multi

function start() { //starts the timer
  startTime = new Date();
}

function end() { //calculates the total time between a guess/answer and end of time
   endTime = new Date();
   var diffTime = endTime - startTime;
   diffTime /= 1000; 
   var seconds = Math.round(diffTime);
   return seconds;
}

function compScoresCalc() { //calcs 
    var timeScore =  0.5;
    if (end() < 20){
      for(let i = 0; i < end(); i++){
        ogComp = ogComp*compMultiplier;
     
      }
    }
    timeScore = ogComp;
    timeScore = Math.round(timeScore);
    return timeScore;
} 

function compWin(turnWin){
    //boiler-plate
    
    const turns = [25,22,18,15,12,9];
    let winScore = turns[turnWin-1] 
    console.log(winScore);
    //return winScore + total comptime calc
    //calc elo rating 

    
}

function compLose(User){
   //I NEED THE SERVER FOR THIS
}
  
//testers
/* start();
  
for(let i = 0; i<2000; i++){
      console.log("yes");
}
   
end();
compWin(6); */
  
  