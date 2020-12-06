<script lang='ts'>

	let name: string = ''
	let reply: string = ''

	let err: string = ''

	function submit() {
		fetch('/api/greet', {
			method: 'post',
			body: name,
		})
		.then(response => response.text())
		.then(text => reply = text)
		.catch(e => err = e)
	}

</script>

{#if err}
	<p class='p-4 text-red'>
		<strong>Oops, something went wrong:</strong>
		<blink>{err}</blink>
	</p>
{/if}

<form 
	on:submit|preventDefault={ submit }
>
	<input 
		bind:value={ name }
		type='text' 
		placeholder='Your name' 
	/>
	<button>Greet</button>
</form>

{#if reply}
	<h2 class='text-red-600'>
		<marquee>{reply}</marquee>
	</h2>
{/if}
