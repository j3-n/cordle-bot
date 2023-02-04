const { SlashCommandBuilder } = require("discord.js");

module.exports = {
    data: new SlashCommandBuilder()
        .setName("test")
        .setDescription("funy comand"),
    async execute(interaction){
        await interaction.reply("cum");
    },
};