const { SlashCommandBuilder } = require('discord.js');
const { Leaderboard } = require('../src/leaderboard-maker');

const medals = ["🥇", "🥈", "🥉"];

module.exports = {
    data: new SlashCommandBuilder()
        .setName("leaderboard")
        .setDescription("Get a view of the current leaderboard"),
    async execute(interaction){
        // create leaderboard
        const leaderboard = new Leaderboard();
        leaderboard.initialize();
        const topTen = await leaderboard.makeTopTen();

        var leaderboardStr = "```[Leaderboard]\n";

        for (let i = 0; i < topTen.length; i++) {
            if(i < medals.length)
                leaderboardStr += medals[i];
            else
                leaderboardStr += `#${i+1}`
            leaderboardStr += (` : ${topTen[i].name} 🏆: ${topTen[i].elo}\n`);
        }

        leaderboardStr += "```";

        interaction.reply({
            content: leaderboardStr,
        });
    },
};