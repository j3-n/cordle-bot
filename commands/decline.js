const { SlashCommandBuilder } = require("discord.js");
const { findChallenge, completeChallenge } = require("../game-manager.js");
const { WordleGame } = require("../src/wordle.js");

module.exports = {
    data: new SlashCommandBuilder()
        .setName("decline")
        .setDescription("Decline a duel invitation against you"),
    async execute(interaction){
        // Create a new game for this channel if it doesnt already exist
        let challenge = findChallenge(interaction.user.id);
        if(challenge){
            completeChallenge(interaction.user.id, challenge.player2);
            interaction.reply({content: "Challenge declined!", ephemeral: true});
        } else
            interaction.reply({content: "There is no challenge against you!", ephemeral: true});
    },
};