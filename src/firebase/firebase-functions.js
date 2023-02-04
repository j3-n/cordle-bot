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
                user.data().games_played,
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
            user.data().games_played,
        );
    }

    async addUser(id, name, games_won, games_lost, games_played, collection) {
        let user = {
            name: name,
            games_won: games_won,
            games_lost: games_lost,
            games_played: games_played
        };

        await this.fbConnection.addDocument(collection, id, user);
    }

    async updateUser(id, name, games_won, games_lost, games_played, collection) {
        let user = {
            name: name,
            games_won: games_won,
            games_lost: games_lost,
            games_played: games_played
        };

        await this.fbConnection.updateDocument(collection, id, user);
    }

    async checkUserExists(id, name, collection) {
        let user = {
            name: name,
        }

        return await this.fbConnection.checkDocument(collection, id, user);
    }
}

module.exports.FirebaseFunctions = FirebaseFunctions;