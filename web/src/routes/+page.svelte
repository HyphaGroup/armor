<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { hasPassword } from '$lib/api';

	onMount(() => {
		if (hasPassword()) {
			goto('/profiles');
		}
	});
</script>

<div class="text-center">
	<p class="text-gray-500">Redirecting...</p>
</div>
