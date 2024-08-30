<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<main>
	<h1>Mosaic</h1>
	<div>Data from GO server: {data.data}</div>
	<form>
		<label>
			Source image
			<input name="src" type="file" />
		</label>
		<label>
			Target image
			<input name="target" type="file" />
		</label>

		<button on:click={() => alert('hello')}>Generate my image!</button>
	</form>
</main>

<style>
	main {
		align-items: center;
		display: flex;
		flex-direction: column;
		padding: 24px;
	}

	form {
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	label {
		display: flex;
		flex-direction: column;
		gap: 4px;
	}
</style>
