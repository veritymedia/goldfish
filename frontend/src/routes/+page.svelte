<script>
	import Cliplist from '../components/cliplist.svelte';
	import SettingsBlock from '../components/settings-block.svelte';
	import { GetAllClipboardItems, GetLatestClipboardItem } from '../lib/wailsjs/go/main/App.js';
	import { EventsOn } from '../lib/wailsjs/runtime/runtime.js';

	// let allClips = [];
	let itemsRaw = [];
	// let itemsGroupedByDate = {};

	GetAllClipboardItems().then((data) => {
		// clipboardHistory = data;
		console.log('data: ', data);

		if (data.length === 0) {
			return;
		}
		itemsRaw = data;

		console.log('clipsGroupedByDate: ', itemsGroupedByDate);
	});

	// this needs to be a computed.
	$: itemsGroupedByDate = itemsRaw.reduce((groups, item) => {
		let date = new Date(item.createdAt);

		// Build a date string in the format yyyy-mm-dd
		let dateString = date.toISOString().split('T')[0];

		if (!groups[dateString]) {
			groups[dateString] = [];
		}
		groups[dateString].push(item);
		return groups;
	}, {});

	function clearClips() {
		console.log('clearing clips');
		clipboardHistory = [];
		clipboardHistory = clipboardHistory;
	}

	function handleClipDeleted(payload) {
		const index = itemsRaw.findIndex((item) => item.createdAt === payload.detail.createdAt);
		console.log('clip deleted: ', index, payload.detail.createdAt);
		itemsRaw.splice(index, 1);
		itemsRaw = itemsRaw;
	}

	// Listen for events from the backend
	EventsOn('clipboard-update', () => {
		console.log('clipboard changed');
		GetLatestClipboardItem().then((data) => {
			console.log('latest clipboard item: ', data);
			itemsRaw.unshift(data[0]);
			itemsRaw = itemsRaw;
		});
	});
</script>

{#if Object.keys(itemsGroupedByDate).length > 0}
	<div
		class="flex overflow-auto flex-col w-full min-h-[100vh] h-full px-2 pt-40 pb-20 bg-background rounded-3xl"
	>
		<h1 class="mb-10 text-5xl text-dark">clipboard history</h1>

		{#each Object.keys(itemsGroupedByDate) as date (date)}
			<Cliplist
				on:clipDeleted={handleClipDeleted}
				clipListByDay={itemsGroupedByDate[date]}
				{date}
			/>
		{/each}
		<!-- <Cliplist {clipboardHistory} /> -->
	</div>
{:else}
	<div class="flex flex-col w-full min-h-[100vh] h-full px-2 pt-20 pb-20 bg-background rounded-3xl">
		<h1 class="mb-10 text-4xl text-dark">clips</h1>

		No clips
	</div>
{/if}

<SettingsBlock on:clipsDeleted={clearClips} />
