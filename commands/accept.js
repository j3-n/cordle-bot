const { SlashCommandBuilder } = require("discord.js");
const { newGame, findChallenge, completeChallenge, gameExists } = require("../game-manager.js");
const { FirebaseFunctions } = require("../src/firebase/firebase-functions.js");

module.exports = {
    data: new SlashCommandBuilder()
        .setName("accept")
        .setDescription("Accept a duel invitation against you"),
    async execute(interaction){
        // Check if a challenge exists for this user
        let challenge = findChallenge(interaction.user.id);
        if(challenge){
            let threadName = `${interaction.user.username}'s game`;
            if(!gameExists(challenge.player1)){
                // Create a thread for the challenge
                let thread = await interaction.channel.threads.create({
                    name: threadName,
                    autoArchiveDuration: 60,
                    reason: `${interaction.user.username} is duelling!`,                    
                });
                newGame(thread.id, challenge.player1, challenge.player2);
                completeChallenge(challenge.player1, challenge.player2);
                interaction.reply({content: "Challenge accepted! Good luck.", ephemeral: true});

                // Create users if they don't exist in the database
                console.log("calling firebase module");
                FirebaseFunctions.createUserIfNotExists(challenge.player1, challenge.player1.user.username);
                FirebaseFunctions.createUserIfNotExists(challenge.player2, challenge.player2.user.username);
            } else
                interaction.reply({content: "You are already in a game!", ephemeral: true});
        } else
            interaction.reply({content: "There is no challenge against you!", ephemeral: true});
    },
};