<script>
	export let clip;
	import { createEventDispatcher } from 'svelte';
	import { DeleteOneByDateCreated } from '../lib/wailsjs/go/main/App.js';

	const dispatch = createEventDispatcher();

	function recopyFromHistory() {
		console.log('INITIATE RECOPY FROM HISTORY: ');
		dispatch('recopyFromHistory', { createdAt: clip.createdAt, content: clip.content });
	}

	function deleteClipByDate() {
		console.log('deleting clip by date: ', clip.createdAt);

		DeleteOneByDateCreated(clip.createdAt).then((data) => {
			if (data === true) {
				dispatch('clipDeleted', { createdAt: clip.createdAt });
			}
		});
	}
</script>

<!-- <button > -->
<button
	on:click={recopyFromHistory(clip.content)}
	class="flex flex-col items-start justify-between gap-2 p-4 group text-dark rounded-xl hover:bg-light"
>
	<div class="flex items-center h-6 text-xs">
		<span class="pr-2 font-bold"
			>{new Date(clip.createdAt).toLocaleTimeString('en-US', {
				hour: 'numeric',
				minute: 'numeric'
			})}</span
		>
		<button
			on:click={(event) => {
				event.stopPropagation();
				deleteClipByDate();
			}}
			class="hidden h-full px-2 font-medium rounded-md hover:bg-highlight group-hover:inline"
			>delete clip</button
		>
	</div>
	<p class="text-overflow">
		{clip.content}
	</p>
</button>

<!-- </button> -->

<style>
	.text-overflow {
		overflow: hidden;
		display: -webkit-box;
		-webkit-line-clamp: 4; /* number of lines to show */
		line-clamp: 4;
		-webkit-box-orient: vertical;
	}
</style>
