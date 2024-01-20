const { Routes, REST } = require('discord.js');
const { CLIENT_ID: clientId, DISCORD_TOKEN: discordToken } = process.env;

const rest = new REST().setToken(discordToken);

rest.put(Routes.applicationCommands(clientId), { body: [] })
	.then(() => console.log('Successfully deleted all application commands.'))
	.catch(console.error);