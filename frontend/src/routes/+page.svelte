<script>
	import Cliplist from '../components/cliplist.svelte';
	import SettingsBlock from '../components/settings-block.svelte';
	import { createEventDispatcher } from 'svelte';
	import { GetAllClipboardItems } from '../lib/wailsjs/go/main/App.js';

	// let allClips = [];

	let itemsGroupedByDate = {};

	GetAllClipboardItems().then((data) => {
		// clipboardHistory = data;
		console.log('data: ', data);

		if (data.length === 0) {
			return;
		}
		itemsGroupedByDate = data.reduce((groups, item) => {
			let date = new Date(item.createdAt);

			// Build a date string in the format yyyy-mm-dd
			let dateString = date.toISOString().split('T')[0];

			if (!groups[dateString]) {
				groups[dateString] = [];
			}
			groups[dateString].push(item);
			return groups;
		}, {});

		console.log('clipsGroupedByDate: ', itemsGroupedByDate);
	});

	function clearClips() {
		console.log('clearing clips');
		clipboardHistory = [];
		clipboardHistory = clipboardHistory;
	}

	const dispatch = createEventDispatcher();
	function clipDeleted(payload) {
		dispatch('clipDeleted', payload);
	}
</script>

{#if Object.keys(itemsGroupedByDate).length > 0}
	<div class="flex flex-col w-full min-h-[100vh] h-full px-2 pt-40 pb-20 bg-background rounded-3xl">
		<h1 class="mb-10 text-5xl text-dark">clipboard history</h1>

		{#each Object.keys(itemsGroupedByDate) as date (date)}
			<Cliplist on:clipDeleted={clipDeleted} clipListByDay={itemsGroupedByDate[date]} {date} />
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
