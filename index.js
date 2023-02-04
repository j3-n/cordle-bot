const { Client, Events, GatewayIntentBits } = require("discord.js");
const { token } = require("./config.json");

// Setup the discord client
const client = new Client({ intents: [GatewayIntentBits.Guilds] });
// Trigger the main function once the client has connected
client.once(Events.ClientReady, main);
client.login(token);

function main(c){
    console.log(c.user.tag);
}