<script lang="ts">
	import { EventsOn } from '../lib/wailsjs/runtime/runtime.js';
	import type { main } from '../lib/wailsjs/go/models.js';

	import { EventsOnce } from '$lib/wailsjs/runtime/runtime.js';
	import { goto } from '$app/navigation';
	const firstTimeListener = EventsOnce('checkIsFirstLaunch', (isFirstLaunch) => {
		console.log('isFirstLaunch: ', isFirstLaunch);
		goto('/onboarding');
	});

	// import { StartClipboardWatcher } from "../lib/wailsjs/go/main/App.js"

	let clipboardHistory: main.ClipboardMessage[] = [];

	EventsOn('newTextData', (data: main.ClipboardMessage) => {
		clipboardHistory.push(data);
		// I hate this. But the variable needs to be reassigned for svelte to update UI
		// Alt: clipboardHistory = [...clipboardHistory, data]
		clipboardHistory = clipboardHistory;
	});
</script>

<ul class="fixed flex flex-col gap-3 -translate-x-1/2 w-min-max left-1/2 top-1/2">
	{#each clipboardHistory as entry}
		<li class="flex gap-5 p-2 text-gray-100 bg-gray-800">
			<span class="font-bold text-green-500 uppercase">Copied text: </span>{entry.text}
			<div class={entry.status === 'ok' ? 'bg-lime-500' : 'bg-amber-500'}>
				{entry.status}
			</div>
		</li>
	{:else}
		<li>No clipboard history yet</li>
	{/each}
</ul>
