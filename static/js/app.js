/**
 * @typedef {Object} AppLocations
 * @property {string} [guildId]
 * @property {string} [channelId]
 * @property {Record<string, string>} [channelLastMessageIds]
 */

document.addEventListener('DOMContentLoaded', async () => {
	let appLocations = localStorage.getItem("app-locations");
	if (appLocations) {
		const parsedAppLocations = JSON.parse(appLocations);
		const guilds = [...document.querySelectorAll('#guilds')];
		/**
		 * @type {HTMLDivElement|null}
		 */
		const selectedGuild = guilds.find(guild => guild.id === parsedAppLocations.guildId)
		if (!selectedGuild) return;
		selectedGuild.classList.add('selected');
	} else {
		localStorage.setItem("app-locations", JSON.stringify({}));
	}
});
