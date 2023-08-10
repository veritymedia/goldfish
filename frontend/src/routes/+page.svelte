<script lang="ts">
	import { EventsOn } from '../lib/wailsjs/runtime/runtime.js';
	import type { main } from '../lib/wailsjs/go/models.js';
	// import { StartClipboardWatcher } from "../lib/wailsjs/go/main/App.js"

	let clipboardHistory: main.ClipboardMessage[] = [];

	EventsOn('newTextData', (data: main.ClipboardMessage) => {
		clipboardHistory.push(data);
		// I hate this. But the variable needs to be reassigned for svelte to update UI
		// Alt: clipboardHistory = [...clipboardHistory, data]
		clipboardHistory = clipboardHistory;
	});
</script>

<ul class="flex flex-col gap-3 fixed w-min-max left-1/2 top-1/2 -translate-x-1/2">
	{#each clipboardHistory as entry}
		<li class="flex gap-5">
			<span class="font-bold text-green-500 uppercase">Copied text: </span>{entry.text}
			<div class={entry.status === 'ok' ? 'bg-lime-500' : 'bg-amber-500'}>
				STATUS: {entry.status}
			</div>
		</li>
	{:else}
		<li>No clipboard history yet</li>
	{/each}
</ul>
