const { SlashCommandBuilder } = require("discord.js");
const { WordleGame } = require("../src/wordle.js");
const gameManager = require("../game-manager.js");

function startGame(interaction){
    gameManager.game = new WordleGame();
    interaction.reply({content: gameManager.game.word});
}

module.exports = {
    data: new SlashCommandBuilder()
        .setName("start")
        .setDescription("Begin a new game of wordle"),
    async execute(interaction){
        startGame(interaction);
    },
};