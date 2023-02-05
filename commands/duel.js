const { SlashCommandBuilder, GuildMember } = require("discord.js");
const { newChallenge } = require("../game-manager.js");

module.exports = {
    data: new SlashCommandBuilder()
        .setName("duel")
        .setDescription("Begin a new game of wordle")
        .addMentionableOption(option => 
            option.setName("opponent")
                .setDescription("Your noble opponent")
                .setRequired(true),
        ),
    async execute(interaction){
        let opponent = interaction.options.getMentionable("opponent");
        if(!isValidOpponent(opponent))
            interaction.reply({content: "Please duel a valid user!", ephemeral: true});
        else {
            // A valid duel has been started
            if(newChallenge(interaction.user.id, opponent.user.id)){
                // Send the duel invite to the other player
                interaction.reply({content: "Challenge sent! Good luck.", ephemeral: true});
                interaction.channel.send(`${opponent.user}, ${interaction.user} has challenged you to a duel! Type \`/accept\` or \`/decline\` to respond!`);
            } else
                interaction.reply({content: "Each player can only have one active challenge!", ephemeral: true});
        }
    },
};

function isValidOpponent(opponent){
    return opponent instanceof GuildMember;
}