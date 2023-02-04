var startTime, endTime, tempScore;
const coopAddScore = 6;
const coopMultiplier = 0.3;


function start() {
  startTime = new Date();
}

function end() { //calculates the total time between a guess/answer
   endTime = new Date();
   var diffTime = endTime - startTime;
   return diffTime;
}

function scoresClac() {
    var timeScore;
    console.log(end());
    console.log(timeScore+coopAddScore);
  
} 
start();

for(let i = 0; i<30430; i++){
    console.log(i);
}
end();
scoresClac();