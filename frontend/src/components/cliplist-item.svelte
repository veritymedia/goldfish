<script>
	export let clip;
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
	import { ClipboardSetText } from '../lib/wailsjs/runtime/runtime.js';
	import { DeleteOneByDateCreated } from '../lib/wailsjs/go/main/App.js';

	function setClipboardText(text) {
		ClipboardSetText(text);
	}

	function deleteClipByDate() {
		console.log('deleting clip by date: ', clip.createdAt);
		let result;
		DeleteOneByDateCreated(clip.createdAt).then((data) => {
			result = data;
		});

		if (result === true) {
			dispatch('clipDeleted', { createdAt: clip.createdAt });
		}

		console.log('deleting clip by date result: ', result);
	}
</script>

<!-- <button > -->
<button
	on:click={setClipboardText(clip.content)}
	class="flex flex-col items-start justify-between gap-2 p-4 group text-dark rounded-xl hover:bg-light"
>
	<div class="text-xs">
		<span class="pr-4 font-bold"
			>{new Date(clip.createdAt).toLocaleTimeString('en-US', {
				hour: 'numeric',
				minute: 'numeric'
			})}</span
		>
		<button on:click={deleteClipByDate} class="hidden group-hover:inline">delete clip</button>
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
