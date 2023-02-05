const { SlashCommandBuilder, GuildMember } = require('discord.js');
const { Stats } = require('../src/stats-maker');
const { FirebaseFunctions } = require('../src/firebase/firebase-functions');

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
        if(player instanceof GuildMember){
            let id = player ? player.id : interaction.user.id;

            if (!FirebaseFunctions.checkUserExists(id)) {
                interaction.reply({
                    content: "This user has never played before!",
                    ephemeral: true
                });
                return;
            }

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
        } else
            interaction.reply({
                content: "Please query a valid user!", 
                ephemeral: true
            });
    },
};