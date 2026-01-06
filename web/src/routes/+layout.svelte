<script lang="ts">
	import '../app.css';
	import { hasPassword, setPassword, clearPassword, api } from '$lib/api';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let { children } = $props();

	let authenticated = $state(false);
	let checking = $state(true);
	let password = $state('');
	let error = $state('');

	onMount(() => {
		authenticated = hasPassword();
		checking = false;
	});

	async function login() {
		error = '';
		const valid = await api.checkPassword(password);
		if (valid) {
			setPassword(password);
			authenticated = true;
		} else {
			error = 'Invalid password';
		}
	}

	function logout() {
		clearPassword();
		authenticated = false;
		password = '';
	}
</script>

<svelte:head>
	<title>ARMOR - Threat Modeling</title>
</svelte:head>

{#if checking}
	<div class="min-h-screen flex items-center justify-center">
		<p class="text-gray-500">Loading...</p>
	</div>
{:else if !authenticated}
	<div class="min-h-screen flex items-center justify-center bg-gray-100">
		<div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
			<h1 class="text-2xl font-bold mb-6 text-center">ARMOR</h1>
			<p class="text-gray-600 mb-6 text-center">Enter password to access</p>
			<form onsubmit={(e) => { e.preventDefault(); login(); }}>
				<input
					type="password"
					bind:value={password}
					placeholder="Password"
					class="w-full px-4 py-2 border rounded-lg mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
				/>
				{#if error}
					<p class="text-red-500 text-sm mb-4">{error}</p>
				{/if}
				<button
					type="submit"
					class="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition"
				>
					Enter
				</button>
			</form>
		</div>
	</div>
{:else}
	<div class="min-h-screen">
		<nav class="bg-white shadow-sm border-b">
			<div class="max-w-7xl mx-auto px-4 py-3 flex justify-between items-center">
				<a href="/profiles" class="text-xl font-bold text-gray-900">ARMOR</a>
				<button
					onclick={logout}
					class="text-gray-600 hover:text-gray-900 text-sm"
				>
					Logout
				</button>
			</div>
		</nav>
		<main class="max-w-7xl mx-auto px-4 py-8">
			{@render children()}
		</main>
	</div>
{/if}
