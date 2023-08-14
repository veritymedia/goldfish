<script>
	import Cliplist from '../components/cliplist.svelte';
	import SettingsIcon from '../components/icons/settings-icon.svelte';
	import SettingsModal from '../components/settings-modal.svelte';
	import { GetAllClipboardItems, GetLatestClipboardItem } from '../lib/wailsjs/go/main/App.js';
	import { DeleteOneByDateCreated } from '../lib/wailsjs/go/main/App.js';

	import { ClipboardSetText, EventsOn } from '../lib/wailsjs/runtime/runtime.js';

	let itemsRaw = [];

	GetAllClipboardItems().then((data) => {
		if (data && data.length === 0) {
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

	function handleClipboardNuke() {
		console.log('clearing clips');
		itemsRaw = [];
		// itemsRaw = itemsRaw;
	}

	function handleClipDeleted(payload) {
		const index = itemsRaw.findIndex((item) => item.createdAt === payload.detail.createdAt);
		console.log('clip deleted: ', index, payload.detail.createdAt);
		itemsRaw.splice(index, 1);
		itemsRaw = itemsRaw;
	}

	function handleRecopyFromHistory(payload) {
		console.log('recopying from history: ', payload.detail);
		console.log('itemsRaw: FIRST ', itemsRaw[0]);
		const { content, createdAt } = payload.detail;
		if (createdAt === itemsRaw[0].createdAt) {
			// ClipboardSetText(content);
			return;
		} else {
			DeleteOneByDateCreated(createdAt).then((data) => {
				if (data === true) {
					handleClipDeleted({ detail: { createdAt: createdAt } });
					// dispatch('clipDeleted', { createdAt: createdAt });
					ClipboardSetText(content);
				}
			});
		}
	}

	// Listen for events from the backend
	EventsOn('clipboard-update', () => {
		console.log('clipboard changed');

		GetLatestClipboardItem().then((data) => {
			console.log('latest clipboard item: ', data);
			if (itemsRaw.length === 0) {
				console.log('update clip DATA ZERO: ', data, 'itemsRaw: ', itemsRaw);
				itemsRaw = [...data];
			} else {
				console.log('update clip  ', data, 'itemsRaw: ', itemsRaw);
				itemsRaw = [data[0], ...itemsRaw];
			}
		});
	});

	let showSettingsModal = false;
</script>

<div
	class="flex overflow-auto flex-col w-full min-h-[100vh] h-full px-2 pt-40 pb-20 bg-background rounded-3xl"
>
	<div class="flex items-center justify-between w-full mb-20">
		<h1 class="text-5xl text-dark">clipboard history</h1>

		<button on:click={() => (showSettingsModal = true)}
			><SettingsIcon classList="text-dark w-8 h-8" /></button
		>
	</div>
	{#if Object.keys(itemsGroupedByDate).length > 0}
		{#each Object.keys(itemsGroupedByDate) as date (date)}
			<Cliplist
				on:clipDeleted={handleClipDeleted}
				on:recopyFromHistory={handleRecopyFromHistory}
				clipListByDay={itemsGroupedByDate[date]}
				{date}
			/>
		{/each}
	{:else}
		<div
			class="flex flex-col w-full min-h-[100vh] h-full px-2 pt-20 pb-20 bg-background rounded-3xl"
		>
			You have no clips yet. Copy something to get started. <br />
		</div>
	{/if}
</div>

<SettingsModal on:clipboardHistoryNuke={handleClipboardNuke} bind:showSettingsModal />
