/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				background: '#F5F1E3',
				light: '#FCE3CC',
				dark: '#4E351F',
				highlight: '#FACF90'
			}
		}
	},
	plugins: []
};
