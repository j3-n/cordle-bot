const { FirebaseConnect } = require('./firebase-interface');

class FirebaseFunctions {
    constructor() {
        this.fbConnection = undefined;
    }

    initialize() {
        this.fbConnection = new FirebaseConnect();
        this.fbConnection.initialize();
    }

    async printUsers() {
        let users = await this.fbConnection.getCollection('users');
        
        await users.forEach((user) => {
            console.log(
                user.id, " ",
                user.data().name, " ",
                user.data().games_won, " ",
                user.data().games_lost, " ",
                user.data().games_played, " ",
                user.data().elo, " ",
                user.data().score, " ",
            );
        });
    }

    async printUser(userId) {
        let user = await this.fbConnection.getDocument('users', userId);
        
        await console.log(
            user.id, " ",
            user.data().name, " ",
            user.data().games_won, " ",
            user.data().games_lost, " ",
            user.data().games_played, " ",
            user.data().elo, " ",
            user.data().score, " ",
        );
    }
    
    async getUser(id, collection) {
        const userData = await this.fbConnection.getDocument(collection, id);

        return {
            id: userData.id,
            name: userData.data().name,
            gamesWon: userData.data().games_won,
            gamesLost: userData.data().games_lost,
            gamesPlayed: userData.data().games_played,
            elo: userData.data().elo,
            score: userData.data().score
        };
    }

    async getUserStats(id, collection) {
        const userData = await this.fbConnection.getDocument(collection, id);
        
        return {
            gamesWon: userData.data().games_won,
            gamesLost: userData.data().games_lost,
            gamesPlayed: userData.data().games_played,
            elo: userData.data().elo,
            score: userData.data().score
        };
    }

    async addUser(id, name, gamesWon, gamesLost, gamesPlayed, elo, score, collection) {
        let user = {
            name: name,
            games_won: gamesWon,
            games_lost: gamesLost,
            games_played: gamesPlayed,
            elo: elo,
            score: score
        };

        await this.fbConnection.addDocument(collection, id, user);
    }

    async updateUser(id, name, gamesWon, gamesLost, gamesPlayed, elo, score, collection) {
        let user = {
            name: name,
            games_won: gamesWon,
            games_lost: gamesLost,
            games_played: gamesPlayed,
            elo: elo,
            score: score
        };

        await this.fbConnection.updateDocument(collection, id, user);
    }

    async checkUserExists(id, collection) {
        return await this.fbConnection.checkDocument(collection, id);
    }

    async createUserIfNotExists(id) {
        if (!this.checkUserExists(id, "users")) {
            await this.fbConnection.addDocument(
                "users",
                id,
                {
                    name: "temp-name",
                    games_won: 0,
                    games_lost: 0,
                    games_played: 0,
                    elo: 500,
                    score: 0
                }
            );
        }
    }
}

module.exports.FirebaseFunctions = FirebaseFunctions;