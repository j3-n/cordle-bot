const { SlashCommandBuilder } = require("discord.js");
const { findGame } = require("../game-manager.js");

module.exports = {
    data: new SlashCommandBuilder()
        .setName("guess")
        .setDescription("Submit a guess to an existing game")
        .addStringOption(option => 
            option.setName("guess")
                .setDescription("Your five-letter guess")
                .setRequired(true),
        ),
    async execute(interaction){
        // Find the game for this channel
        let game = findGame(interaction.channelId);
        if(game){
            console.log(interaction.options.getString("guess"));
        } else
            interaction.reply({content: "There is no active game in this channel! Use `/start` to start one."});
    },
};