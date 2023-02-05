const { SlashCommandBuilder } = require('discord.js');
const { Stats } = require('../src/stats-maker');

module.exports = {
    data: new SlashCommandBuilder()
        .setName("stats")
        .setDescription("Get a view of your own stats or someone elses")
        .addMentionableOption(option => 
            option.setName("player")
                .setDescription("The player to query")
        ),
    async execute(interaction){
        // create leaderboard
        let player = interaction.options.getMentionable("player");
        let id = player ? player.id : interaction.user.id;

        const statsHandler = new Stats(id);
        const stats = await statsHandler.makeStats();

        const statsStr = `\`\`\` [${stats.name}]
        🏆: ${stats.elo}
        ==================
        Wins: ${stats.gamesWon}
        Losses: ${stats.gamesLost}
        Total: ${stats.gamesPlayed}
        Score gained: ${stats.score}
        \`\`\``;

        interaction.reply({
            content: statsStr,
            ephemeral: true
        });
    },
};