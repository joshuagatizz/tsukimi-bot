const { Routes, REST } = require('discord.js');
const fs = require('node:fs');
const path = require('node:path');
const { CLIENT_ID: clientId, DISCORD_TOKEN: discordToken } = process.env;

const commands = [];
const commandsPath = path.join(__dirname, 'commands');
const commandFiles = fs.readdirSync(commandsPath).filter(file => file.endsWith('.js'));

for (const file of commandFiles) {
	const filePath = path.join(commandsPath, file);
	const command = require(filePath);
	commands.push(command.data.toJSON());
}
const rest = new REST().setToken(discordToken);

rest.put(Routes.applicationCommands(clientId), { body: commands })
	.then(() => console.log('Successfully registered application commands.'))
	.catch(console.error);