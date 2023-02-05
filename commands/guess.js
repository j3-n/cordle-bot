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
            if(result.condition == Conditions.INVALID_ID)
                interaction.reply({content: "You are not part of this game!", ephemeral: true});
            else if(result.condition == Conditions.INVALID_INPUT)
                interaction.reply({content: "Invalid input!", ephemeral: true});
            else if(result.condition == Conditions.INVALID_WORD)
                interaction.reply({content: "That is not a word!", ephemeral: true});
            else if(result.result){
                reply = `${guess} \n`;
                result.result.guess.forEach(char => {
                // TODO: replace with switch statement
                emote_name = "🟥";
                if(char.result == Result.CORRECT_CHARACTER)
                        emote_name = "🟩";
                    else if(char.result == Result.INCORRECT_POSITION)
                        emote_name = "🟨";
                    reply += emote_name;
                });
                interaction.reply({content: reply, ephemeral: true});
            }

            if(result.condition == Conditions.BOTH_PLAYERS_OUT){
                interaction.channel.send("TIE! All players are out of guesses!");
                completeGame(interaction.channelId);
            } else if(result.condition == Conditions.OUT_OF_GUESSES){
                if(result.result)
                    interaction.channel.send(`${interaction.user} has run out of guesses!`);
                else 
                    interaction.reply({content: "You have run out of guesses!", ephemeral: true});
            } else if(result.condition == Conditions.WIN){
                interaction.channel.send(`${interaction.user} has WON! The word was \`${game.player1.word}\`.`);
                completeGame(interaction.channelId);
            }
        } else
            interaction.reply({content: "There is no active game in this channel!", ephemeral: true});
    },
};