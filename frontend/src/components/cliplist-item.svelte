<script>
	import { ClipboardSetText } from '../lib/wailsjs/runtime/runtime.js';

	function setClipboardText(text) {
		ClipboardSetText(text);
	}

	function deleteClipByDate(date) {
		console.log('deleting clip by date: ', date);
	}

	export let clip;
</script>

<button on:click={setClipboardText(clip.content)}>
	<li
		class="flex flex-col items-start justify-between gap-2 p-4 group text-dark rounded-xl hover:bg-light"
	>
		<div class="text-xs">
			<span class="pr-4 font-bold"
				>{new Date(clip.createdAt).toLocaleTimeString('en-US', {
					hour: 'numeric',
					minute: 'numeric'
				})}</span
			>
			<button on:click={deleteClipByDate(clip.createdAt)} class="hidden group-hover:inline"
				>delete clip</button
			>
		</div>
		<p class="text-overflow">
			{clip.content}
		</p>
	</li>
</button>

<style>
	.text-overflow {
		overflow: hidden;
		display: -webkit-box;
		-webkit-line-clamp: 4; /* number of lines to show */
		line-clamp: 4;
		-webkit-box-orient: vertical;
	}
</style>
