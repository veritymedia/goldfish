<script>
	import '../app.css';
	import Header from '../components/header.svelte';
</script>

<div class=" bg-background">
	<Header />
	<div class="px-2">
		<slot />
	</div>
</div>
