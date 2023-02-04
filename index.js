const { Client, Events, GatewayIntentBits, REST, Routes, Collection } = require("discord.js");
const fs = require("fs");
const path = require("path");
const config = require("./config.json");

// Setup the discord client
const client = new Client({ intents: [GatewayIntentBits.Guilds] });
// Trigger the main function once the client has connected
client.once(Events.ClientReady, main);
// Setup event handlers
client.on(Events.InteractionCreate, interactionCreated);

// Load commands
const commandPath = path.join(__dirname, 'commands');
const commandFiles = fs.readdirSync(commandPath).filter(f => f.endsWith('.js'));

client.commands = new Collection();

commandFiles.forEach(file => {
    let command = require(path.join(commandPath, file));
    client.commands.set(command.data.name, command);
});

client.login(config.token);

function main(c){
    // Display a success message
    console.log(`Logged in successfully [${c.user.tag}]`);
    c.user.setActivity("Wordle");
}

async function interactionCreated(interaction){
    let command = interaction.client.commands.get(interaction.commandName);
    await command.execute(interaction);
}