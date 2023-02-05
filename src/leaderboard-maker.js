const { FirebaseFunctions } = require('./firebase-functions');

class Leaderboard {
    constructor() {
        this.fbFunctions = undefined;
        this.users = undefined;
        this.topTen = undefined;
    }

    initialize() {
        this.fbFunctions = new FirebaseFunctions();
        this.fbFunctions.initialize();
        this.users = [];
        this.topTen = [];
    }

    async makeTopTen() {
        this.users = await this.fbFunctions.getUsers("users");

        if (this.users.length <= 10) {
            var leadboardLength = this.users.length;
        } else {
            var leadboardLength = 10;
        }

        this.users.sort((a, b) => {
            if (a.elo > b.elo) {
                return -1;
            } else {
                return 0;
            }
        });

        for (let i = 0; i < leadboardLength; i++) {
            this.topTen.push(this.users[i]);
        }

        return this.topTen;
    }
}

module.exports.Leaderboard = Leaderboard;