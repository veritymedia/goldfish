<script>
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
</script>

<div>
	<h2 class="text-sm font-light text-dark">{formattedDate}</h2>

	<div class="mt-2 mb-5">
		<ul class="flex flex-col gap-3 p-2 border-2 rounded-2xl border-dark">
			<!-- {clipListByDay} -->
			{#each clipListByDay as clip}
				<!-- {clip.createdAt} -->
				<CliplistItem {clip} />
			{/each}
		</ul>
	</div>
</div>
