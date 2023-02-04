var startTime, endTime, tempScore;
const coopAddScore = 6;
const coopMultiplier = 0.3;


function start() {
  startTime = new Date();
  return startTime;
}

function end() { //calculates the total time between a guess/answer
   endTime = new Date();
   var diffTime = endTime - start();
   return diffTime;
}

function scoresClac() {
    var timeScore;
    console.log(end());
    console.log(timeScore+coopAddScore);
  
} 
start();

for(let i = 0; i<30430; i++){
    console.log();
}
end();
scoresClac();