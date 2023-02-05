const { SlashCommandBuilder } = require('discord.js');
const { Stats } = require('../src/stats-maker');

module.exports = {
    data: new SlashCommandBuilder()
        .setName("stats")
        .setDescription("Get a view of your own stats or someone elses"),
    async execute(interaction){
        // create leaderboard
        const statsHandler = new Stats("natan#71912");
        const stats = await statsHandler.makeStats();

        const statsStr = stats.id + "'s stats:\n" + 
            "Games won: " + stats.gamesWon + "\n" + 
            "Games lost: " + stats.gamesLost + "\n" +
            "Games played: " + stats.gamesPlayed + "\n" +
            "Player elo: " + stats.elo + "\n" +
            "Player score: " + stats.score + "\n";

        interaction.reply({
            content: statsStr,
            ephemeral: true
        });
    },
};