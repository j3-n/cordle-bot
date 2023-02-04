var startTime, endTime;

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

module.exports = {
    start,
    end,
    startTime,
    endTime
}