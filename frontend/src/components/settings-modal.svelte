<script>
	import BaseButton from './base/base-button.svelte';
	import XIcon from './icons/x-icon.svelte';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	import { DeleteAllClipboardItems } from '../lib/wailsjs/go/main/App.js';
	function deleteAllClips() {
		DeleteAllClipboardItems().then((hasDeleted) => {
			console.log('all clips deleted: ', hasDeleted);
			if (hasDeleted) {
				dispatch('clipboardHistoryNuke', {});
			}
		});
	}
	export let showSettingsModal; // boolean

	let dialog; // HTMLDialogElement

	$: if (dialog && showSettingsModal) dialog.showModal();
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
	bind:this={dialog}
	on:close={() => (showSettingsModal = false)}
	on:click|self={() => dialog.close()}
	class="w-2/3 bg-transparent h-2/3"
>
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div
		class="flex flex-col justify-between w-full h-full border-2 rounded-xl bg-background border-dark"
		on:click|stopPropagation
	>
		<div>
			<div class="flex items-baseline justify-between w-full">
				<h1 class="pb-10 text-5xl text-dark">settings</h1>

				<!-- svelte-ignore a11y-autofocus -->
				<button class="h-full text-dark" autofocus on:click={() => dialog.close()}
					><XIcon classList="w-6 h-6" /></button
				>
			</div>
			<section class="flex flex-col gap-2 text-dark">
				<h2 class="text-xl">delete all clipboard history</h2>
				<BaseButton on:click={deleteAllClips}>wipe history</BaseButton>
				<p><span class="font-bold text-red-500">warning:</span> this is not reversible</p>
			</section>
		</div>
	</div>
</dialog>

<style>
	dialog {
		max-width: 32em;
		border-radius: 0.2em;
		border: none;
		padding: 0;
	}
	dialog::backdrop {
		background: rgba(0, 0, 0, 0.1);
		backdrop-filter: blur(4px);
		animation: fade 0.5s ease;
	}
	dialog > div {
		padding: 1em;
	}
	dialog[open] {
		animation: zoom 0.5s ease;
	}
	@keyframes zoom {
		from {
			transform: scale(0.95);
		}
		to {
			transform: scale(1);
		}
	}
	dialog[open]::backdrop {
		animation: fade 0.5s ease;
	}

	@keyframes fade {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}
	button {
		display: block;
	}
</style>
