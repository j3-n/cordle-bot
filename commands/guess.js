const { SlashCommandBuilder } = require("discord.js");
const { findGameByChannelId, completeGame } = require("../game-manager.js");
const { Conditions } = require("../src/duel-game.js");
const { Result } = require("../src/wordle.js");
const { ResultHandler } = require("../src/result-handler.js");

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
            let result = game.game.submitGuess(interaction.user.id, guess);
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
                interaction.channel.send(`TIE! All players are out of guesses! The word was \`${game.game.player1.word}\`.`);
                endGame(interaction.channel);
            } else if(result.condition == Conditions.OUT_OF_GUESSES){
                if(result.result)
                    interaction.channel.send(`${interaction.user} has run out of guesses!`);
                else 
                    interaction.reply({content: "You have run out of guesses!", ephemeral: true});
            } else if(result.condition == Conditions.WIN){
                interaction.channel.send(`${interaction.user} has WON! The word was \`${game.game.player1.word}\`.`);
                endGame(interaction.channel);
                // Update the database
                // TODO: number of guesses taken
                let handler = new ResultHandler(game.player1Id, game.player2Id, interaction.user.id, 1);
                handler.postResult();
            }
        } else
            interaction.reply({content: "There is no active game in this channel!", ephemeral: true});
    },
};

function endGame(channel){
    completeGame(channel.id);
    //channel.threads.cache.find(th => th.id == channel.id).setArchive(true);
}