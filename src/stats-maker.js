const { FirebaseFunctions } = require('./firebase/firebase-functions');

class Stats {
    constructor(userId) {
        this.fbFunctions = undefined;
        this.userId = userId;
        this.userObj = undefined;
    }

    initialize() {
        this.fbFunctions = new FirebaseFunctions();
        this.fbFunctions.initialize();
    }

    async makeStats() {
        this.userObj = await this.fbFunctions.getUser(this.userId, "users");

        return {
            id: this.userId,
            name: this.userObj.name,
            gamesWon: this.userObj.gamesWon,
            gamesLost: this.userObj.gamesLost,
            gamesPlayed: this.userObj.gamesPlayed,
            elo: this.userObj.elo,
            score: this.userObj.score
        }
    }
}

module.exports.Stats = Stats;