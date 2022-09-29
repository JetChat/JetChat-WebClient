document.addEventListener('ready', async () => {
	const cookies = document.cookie.split(';');
	const error = cookies.find(cookie => cookie.startsWith('error'));

	if (error) {
		document.querySelector('#error').innerText = error.split('=')[1];
	}
});
