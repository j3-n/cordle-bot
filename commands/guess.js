const { SlashCommandBuilder } = require("discord.js");
const { findGameByChannelId, completeGame } = require("../game-manager.js");
const { Conditions } = require("../src/duel-game.js");
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
        let game = findGameByChannelId(interaction.channelId);
        if(game){
            let guess = interaction.options.getString("guess");
            game = game.game;
            let result = game.submitGuess(interaction.user.id, guess);
            switch(result){
                case Conditions.INVALID_ID:
                    interaction.reply({content: "You are not part of this game!", ephemeral: true});
                    break;
                default:
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
                    interaction.reply({content: reply});
                    break;
            }
        } else
            interaction.reply({content: "There is no active game in this channel! Use `/start` to start one.", ephemeral: true});
    },
};