<script>
	import BaseButton from './base/base-button.svelte';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	// import { DeleteAllClips } from '../lib/wailsjs/go/main/App.js';
	function deleteAllClips() {
		console.log('deleting all clips');
		dispatch('clipsDeleted', {});
		// Send event to parent to clear list if delete successful
	}
</script>

<div class="fixed bottom-0 w-full p-4 border-t-2 border-dark bg-background text-dark text">
	<div class="w-1/2 lowercase">
		<h3>delete all of clips. This is irreversible!</h3>
		<!-- <BaseButton onclick={deleteAllClips}>delete all clips</BaseButton> -->
	</div>
</div>
