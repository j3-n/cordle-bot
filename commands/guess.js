const { SlashCommandBuilder } = require("discord.js");
const { findGame } = require("../game-manager.js");
const { Result } = require("../src/wordle.js");

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
            let guess = interaction.options.getString("guess");
            let result = game.submitGuess(guess);
            if(!result){
                interaction.reply({content: "That guess was invalid!", ephemeral: true});
                return;
            }
            
            reply = `${guess} \n`;
            result.guess.forEach(char => {
                // TODO: replace with switch statement
                emote_name = "🟥";
                if(char.result == Result.CORRECT_CHARACTER)
                    emote_name = "🟩";
                else if(char.result == Result.INCORRECT_POSITION)
                    emote_name = "🟨";
                reply += emote_name;
            });
            console.log(interaction.client.emojis);
            interaction.reply({content: reply});
        } else
            interaction.reply({content: "There is no active game in this channel! Use `/start` to start one.", ephemeral: true});
    },
};