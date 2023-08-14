<script>
	import { onMount } from 'svelte';
	import Cliplist from '../components/cliplist.svelte';
	import SettingsIcon from '../components/icons/settings-icon.svelte';
	import SettingsModal from '../components/settings-modal.svelte';
	import { GetAllClipboardItems, GetLatestClipboardItem } from '../lib/wailsjs/go/main/App.js';
	import { DeleteOneByDateCreated } from '../lib/wailsjs/go/main/App.js';

	import { ClipboardSetText, EventsOn } from '../lib/wailsjs/runtime/runtime.js';

	let itemsRaw = [];

	onMount(() => {
		GetAllClipboardItems().then((data) => {
			console.log('remote list: ', data, 'itemsRaw: ', itemsRaw);

			if (data && data.length > 0) {
				itemsRaw = [...data];
			}
		});

		// console.log('clipsGroupedByDate: ', itemsGroupedByDate);
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
		if (itemsRaw.length === 0) {
			return;
		}
		console.log('clearing clips');
		itemsRaw = [];
	}

	function deleteLocalClipByDate(payload) {
		if (itemsRaw.length === 0) {
			return;
		}
		const index = itemsRaw.findIndex((item) => item.createdAt === payload.detail.createdAt);
		console.log('clip deleted: ', index, payload.detail.createdAt);
		itemsRaw.splice(index, 1);
		itemsRaw = itemsRaw;
	}

	function handleRecopyFromHistory(payload) {
		const { content, createdAt } = payload.detail;
		if (createdAt === itemsRaw[0].createdAt) {
			console.log('tried to recopy. But it was the most recent');
			// ClipboardSetText(content);
			return;
		} else {
			DeleteOneByDateCreated(createdAt).then((data) => {
				if (data === true) {
					deleteLocalClipByDate({ detail: { createdAt: createdAt } });
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
				itemsRaw = [data[0]];
			} else {
				console.log('update clip  ', data, 'itemsRaw: ', itemsRaw);
				itemsRaw = [data[0], ...itemsRaw];
			}
		});
	});

	let showSettingsModal = false;
	function openSettingsModal() {
		console.log('opening settings modal');
		showSettingsModal = true;
	}
</script>

<div
	class="flex overflow-scroll hide-scrollbars flex-col w-full h-screen xmin-h-[100vh] px-2 pt-40 xpb-20 bg-background rounded-3xl"
>
	<div class="sticky flex items-center justify-between w-full pb-4 mb-20 -top-28 bg-background">
		<h1 class="text-5xl text-dark">clipboard</h1>

		<button on:click={openSettingsModal}><SettingsIcon classList="text-dark w-7 h-7" /></button>
	</div>
	{#if Object.keys(itemsGroupedByDate).length > 0}
		{#each Object.keys(itemsGroupedByDate) as date (date)}
			<Cliplist
				on:clipDeleted={deleteLocalClipByDate}
				on:recopyFromHistory={handleRecopyFromHistory}
				clipListByDay={itemsGroupedByDate[date]}
				{date}
			/>
		{/each}
	{:else}
		<div class="flex flex-col text-dark w-full min-h-[100vh] h-full px-2 bg-background rounded-3xl">
			You have no clips yet. <br />
			Copy something to get started. <br />
		</div>
	{/if}
</div>

<SettingsModal on:clipboardHistoryNuke={handleClipboardNuke} bind:showSettingsModal />

<style>
	.hide-scrollbars {
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}

	.hide-scrollbars::-webkit-scrollbar {
		display: none;
	}
</style>
