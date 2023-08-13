<script>
	// import { EventsOn } from '../lib/wailsjs/runtime/runtime.js';

	import { GetAllClipboardItems } from '../lib/wailsjs/go/main/App.js';

	let clipboardHistory = [];

	GetAllClipboardItems().then((data) => {
		clipboardHistory = data;
		console.log(data);
	});

	// EventsOn('newTextData', (data) => {
	// 	clipboardHistory.push(data);

	// 	clipboardHistory = clipboardHistory;
	// });
</script>

<ul class="fixed flex flex-col gap-3 -translate-x-1/2 w-min-max left-1/2 top-1/2">
	{#each clipboardHistory as entry}
		<li class="flex gap-5 p-2 text-gray-100 bg-gray-800">
			<span class="font-bold text-green-500 uppercase">Copied text: </span>{entry.content}
			<div class={entry.status === 'ok' ? 'bg-lime-500' : 'bg-amber-500'}>
				{entry.createdAt}
			</div>
		</li>
	{:else}
		<li>No clipboard history yet</li>
	{/each}
</ul>
