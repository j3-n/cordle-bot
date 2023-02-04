const { SlashCommandBuilder, GuildMember } = require("discord.js");
const { newGame } = require("../game-manager.js");

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
        if(!isValidOpponent(interaction.options.getMentionable("opponent"))){
            interaction.reply({content: "Please duel a valid user!", ephemeral: true});
            return;
        }
    },
};

function isValidOpponent(opponent){
    return opponent instanceof GuildMember;
}