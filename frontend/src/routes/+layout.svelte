<script>
	import { onMount, onDestroy } from 'svelte';
	import '../app.css';
	import Header from '../components/header.svelte';

	function stopContextMenu(event) {
		event.preventDefault();
	}

	onMount(() => {
		document.addEventListener('contextmenu', stopContextMenu);
	});

	onDestroy(() => {
		document.removeEventListener('contextmenu', stopContextMenu);
	});
</script>

<div class=" bg-background">
	<Header />
	<div class="px-2">
		<slot />
	</div>
</div>
