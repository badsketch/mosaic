<script lang="ts">
</script>

<main>
	<h1>Mosaic</h1>
	<form method="POST" use:enhance enctype="multipart/form-data">
		<label>
			Source image 1
			<input name="first" type="file" accept=".jpg, .jpeg, .png, .webp" required />
		</label>
		<label>
			Source image 2
			<input name="second" type="file" accept=".jpg, .jpeg, .png, .webp" required />
		</label>
		<label>
			Target image
			<input name="target" type="file" accept=".jpg, .jpeg, .png, .webp" required />
		</label>

		<button>Generate my image!</button>
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
