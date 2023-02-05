const { SlashCommandBuilder } = require('discord.js');
const { Leaderboard } = require('../src/leaderboard-maker');

module.exports = {
    data: new SlashCommandBuilder()
        .setName("leaderboard")
        .setDescription("Get a view of the current leaderboard"),
    async execute(interaction){
        // create leaderboard
        const leaderboard = new Leaderboard();
        leaderboard.initialize();
        const topTen = leaderboard.makeTopTen();

        var leaderboardStr = "";

        for (let i = 0; i < topTen.length; i++) {
            leaderboardStr += (`${i+1}. ${topTen[i].name}:${topTen[i].elo}\n`);
        }

        interaction.reply({
            content: leaderboardStr,
            ephemeral: true
        });
    },
};