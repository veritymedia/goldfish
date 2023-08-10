<script lang="ts">
	import { EventsOn } from '../lib/wailsjs/runtime/runtime.js';
	import type { main } from '../lib/wailsjs/go/models.js';
	// import { StartClipboardWatcher } from "../lib/wailsjs/go/main/App.js"

	let clipboardHistory: main.ClipboardMessage[] = [];

	EventsOn('newTextData', (data: main.ClipboardMessage) => {
		console.log('newTextData', data);
		clipboardHistory.push(data);
		clipboardHistory = clipboardHistory;
		console.log(clipboardHistory);
	});
</script>

<ul class="flex flex-col gap-3">
	{#each clipboardHistory as entry}
		<li>
			<span class="font-bold text-green-500 uppercase">Copied text: </span>{entry.text}
			STATUS: {entry.status}
		</li>
	{:else}
		<li>No clipboard history yet</li>
	{/each}
</ul>
