const { SlashCommandBuilder, EmbedBuilder } = require('discord.js');
const _ = require('lodash');
const gif = 'https://media.tenor.com/c9WptHOa_LMAAAAC/pong.gif';

module.exports = {
	data: new SlashCommandBuilder()
		.setName('ping')
		.setDescription('Replies with Pong!'),
	async execute(interaction) {
		const embedReply1 = new EmbedBuilder()
			.setColor('Green')
			.setTitle('Pong! 🏓');
		const embedReply2 = new EmbedBuilder()
			.setColor('Green')
			.setImage(gif);

		const reply = _.sample([embedReply1, embedReply2]);

		await interaction.reply({ embeds: [reply] });
	},
};