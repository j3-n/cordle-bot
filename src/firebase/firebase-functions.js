const { FirebaseConnect } = require('./firebase-interface');

class FirebaseFunctions {
    constructor() {
        this.firebaseConnection = undefined;
    }

    initialize() {
        this.firebaseConnection = new FirebaseConnect();
        this.firebaseConnection.initialize();
    }

    async printUsers() {
        let users = await this.firebaseConnection.getCollection('users');
        
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
        let user = await this.firebaseConnection.getDocument('users', userId);
        
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
    
    async getUser(id, collection="users") {
        const userData = await this.firebaseConnection.getDocument(collection, id);

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

    async getUsers(collection="users") {
        const usersData = await this.firebaseConnection.getCollection(collection);
        const users = [];
        
        usersData.forEach((user) => {
            users.push({
                id: user.id,
                name: user.data().name,
                gamesWon: user.data().games_won,
                gamesLost: user.data().games_lost,
                gamesPlayed: user.data().games_played,
                elo: user.data().elo,
                score: user.data().score
            })
        });

        return users;
    }

    async getUserStats(id, collection="users") {
        const userData = await this.firebaseConnection.getDocument(collection, id);
        
        return {
            gamesWon: userData.data().games_won,
            gamesLost: userData.data().games_lost,
            gamesPlayed: userData.data().games_played,
            elo: userData.data().elo,
            score: userData.data().score
        };
    }

    async getUserElo(id, collection="users") {
        const userData = await this.firebaseConnection.getDocument(collection, id);
        return userData.data().elo;
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

        await this.firebaseConnection.addDocument(collection, id, user);
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

        await this.firebaseConnection.updateDocument(collection, id, user);
    }

    async checkUserExists(id, collection="users") {
        return await this.firebaseConnection.checkDocument(collection, id);
    }

    async createUserIfNotExists(id, name) {
        if (await this.checkUserExists(id, "users")) {
            await this.firebaseConnection.addDocument(
                "users",
                id,
                {
                    name: name,
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

let functions = new FirebaseFunctions();
functions.initialize();
module.exports.FirebaseFunctions = functions;