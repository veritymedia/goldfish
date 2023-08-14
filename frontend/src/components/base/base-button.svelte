<script>
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	function sendClickEvent() {
		dispatch('click', {});
	}
</script>

<button
	on:click={sendClickEvent}
	class="px-2 py-1 text-lg lowercase border rounded-lg bg-light text-dark border-dark"
	><slot /></button
>
