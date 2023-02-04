var startTime, endTime;

function start() { //starts the timer
    startTime = new Date();
}

function end() { //calculates the total time between a guess/answer and end of time
     endTime = new Date();
     var diffTime = endTime - startTime;
     return diffTime;
}

module.exports = {
    start,
    end,
    startTime,
    endTime
}