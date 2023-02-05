const { FirebaseFunctions } = require('./firebase/firebase-functions');
const { compWin, compLose } = require('./comp-scores');

class ResultHandler {
    constructor(userIdOne, userIdTwo, winner, attempts) {
        this.uIdOne = userIdOne;
        this.uIdTwo = userIdTwo;
        this.winner = winner;
        this.attempts = attempts;
    }

    async postResult() {
        if (this.uIdOne === this.winner) {
            var userOne = await compWin(this.attempts, this.uIdOne);
            var userTwo = await compLose(this.attempts, this.uIdTwo);
        } else if (this.uIdTwo === this.winner) {
            var userTwo = await compWin(this.attempts, this.uIdTwo);
            var userOne = await compLose(this.attempts, this.uIdOne);
        } else {
            return;
        }

        await FirebaseFunctions.updateUser(
            userOne.id,
            userOne.name,
            userOne.gamesWon,
            userOne.gamesLost,
            userOne.gamesPlayed,
            userOne.elo,
            userOne.score,
            userOne.collection
        );
            
        await FirebaseFunctions.updateUser(
            userTwo.id,
            userTwo.name,
            userTwo.gamesWon,
            userTwo.gamesLost,
            userTwo.gamesPlayed,
            userTwo.elo,
            userTwo.score,
            userTwo.collection
        );
    }
}

module.exports.ResultHandler = ResultHandler;