<script>
	import XIcon from '../icons/x-icon.svelte';
	import CloseIcon from '../icons/close-icon.svelte';

	import { WindowHide, WindowMinimise } from '../../lib/wailsjs/runtime/runtime.js';
	function useWindowMinimize() {
		WindowMinimise();
	}

	function useWindowClose() {
		WindowHide();
	}
</script>

<div class="flex items-center justify-between">
	<div class="flex items-center gap-2">
		<!-- <Logo class="w-10 h-10" /> -->
		<div>🐠 Goldfish</div>
	</div>
	<div class="flex" style="--wails-draggable:no-drag">
		<button
			on:click={useWindowMinimize}
			class="flex items-center justify-center w-8 h-8 hover:bg-light"><CloseIcon /></button
		>
		<button
			on:click={useWindowClose}
			class="flex items-center justify-center w-8 h-8 hover:bg-light"
			><XIcon classList="text-dark" /></button
		>
	</div>
</div>
