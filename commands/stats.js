const { SlashCommandBuilder } = require('discord.js');
const { Stats } = require('../src/stats-maker');

module.exports = {
    data: new SlashCommandBuilder()
        .setName("stats")
        .setDescription("Get a view of your own stats or someone elses"),
    async execute(interaction){
        // create leaderboard
        const stats = new Stats("natan#71912");
        const statsStr = await stats.makeStats();

        interaction.reply({
            content: statsStr,
            ephemeral: true
        });
    },
};