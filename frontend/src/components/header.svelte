<script>
	import { Environment } from '../lib/wailsjs/runtime/runtime.js';
	import HeaderLinux from './header/header-linux.svelte';
	import HeaderDarwin from './header/header-darwin.svelte';
	import HeaderWin from './header/header-win.svelte';

	let env;
	let platform;
	Environment().then((data) => {
		env = data;

		platform = env.platform;
	});

	const headerMap = {
		darwin: HeaderDarwin,
		windows: HeaderWin,
		linux: HeaderLinux
	};
</script>

<div style="--wails-draggable:drag" class="fixed top-0 w-full p-2 bg-background text-dark">
	<!-- <svelte:component this={headerMap[platform]} /> -->
	<svelte:component this={HeaderLinux} />
</div>
