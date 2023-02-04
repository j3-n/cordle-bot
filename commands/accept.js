const { SlashCommandBuilder } = require("discord.js");
const { newGame, findChallenge, completeChallenge } = require("../game-manager.js");
const { DuelWordle } = require("../src/duel-game.js");

module.exports = {
    data: new SlashCommandBuilder()
        .setName("accept")
        .setDescription("Accept a duel invitation against you"),
    async execute(interaction){
        // Check if a challenge exists for this user
        let challenge = findChallenge(interaction.user.id);
        if(challenge){
            if(newGame(interaction.channelId, new DuelWordle(challenge.player1, challenge.player2))){
                completeChallenge(challenge.player1, challenge.player2);
                interaction.reply({content: "Challenge accepted! Good luck.", ephemeral: true});

                // Create a new thread for the game to take place in
                interaction.channel.threads.create({
                    name: `${interaction.user.username}'s game`,
                    autoArchiveDuration: 60,
                    reason: `${interaction.user.username} is duelling!`,                    
                });
            } else
                interaction.reply({content: "You are already in a game!", ephemeral: true});
        } else
            interaction.reply({content: "There is no challenge against you!", ephemeral: true});
    },
};