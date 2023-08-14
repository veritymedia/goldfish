<script>
	import { createEventDispatcher } from 'svelte';

	// import { EventsOn } from '../lib/wailsjs/runtime/runtime.js';
	import CliplistItem from './cliplist-item.svelte';
	export let clipListByDay;
	export let date;
	function formatDate(date) {
		const dateObject = new Date(date);
		let formattedDate = dateObject.toLocaleDateString('en-US', {
			// weekday: 'long',
			// year: 'numeric',
			month: 'long',
			day: 'numeric'
		});
		return formattedDate;
	}
	let formattedDate = formatDate(date);
	console.log('clipListByDay: ', clipListByDay);

	const dispatch = createEventDispatcher();
	function propagateClipDeleted(payload) {
		dispatch('clipDeleted', { createdAt: payload.detail.createdAt });
	}

	function propagateRecopyFromHistory(payload) {
		dispatch('recopyFromHistory', {
			createdAt: payload.detail.createdAt,
			content: payload.detail.content
		});
	}
</script>

<div>
	<h2 class="text-sm font-light text-dark">{formattedDate}</h2>

	<div class="mt-2 mb-5">
		<ul class="flex flex-col gap-3 p-2 border-2 rounded-2xl border-dark">
			<!-- {clipListByDay} -->
			{#each clipListByDay as clip}
				<!-- {clip.createdAt} -->
				<CliplistItem
					{clip}
					on:clipDeleted={propagateClipDeleted}
					on:recopyFromHistory={propagateRecopyFromHistory}
				/>
			{/each}
		</ul>
	</div>
</div>
