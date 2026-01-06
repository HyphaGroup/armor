<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api';

	let name = $state('');
	let description = $state('');
	let loading = $state(false);
	let error = $state('');

	async function create() {
		if (!name.trim()) {
			error = 'Name is required';
			return;
		}

		loading = true;
		error = '';

		try {
			const profile = await api.createProfile(name.trim(), description.trim());
			goto(`/profiles/${profile.id}`);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to create profile';
		} finally {
			loading = false;
		}
	}
</script>

<div class="max-w-xl mx-auto">
	<h1 class="text-2xl font-bold mb-6">New Profile</h1>

	<form onsubmit={(e) => { e.preventDefault(); create(); }} class="bg-white rounded-lg shadow p-6 space-y-4">
		<div>
			<label for="name" class="block text-sm font-medium text-gray-700 mb-1">
				Name <span class="text-red-500">*</span>
			</label>
			<input
				id="name"
				type="text"
				bind:value={name}
				placeholder="e.g., Acme Organization"
				class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
			/>
		</div>

		<div>
			<label for="description" class="block text-sm font-medium text-gray-700 mb-1">
				Description
			</label>
			<textarea
				id="description"
				bind:value={description}
				placeholder="Optional description of this threat model"
				rows="3"
				class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
			></textarea>
		</div>

		{#if error}
			<p class="text-red-500 text-sm">{error}</p>
		{/if}

		<div class="flex gap-4">
			<a
				href="/profiles"
				class="flex-1 text-center px-4 py-2 border rounded-lg hover:bg-gray-50 transition"
			>
				Cancel
			</a>
			<button
				type="submit"
				disabled={loading}
				class="flex-1 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition disabled:opacity-50"
			>
				{loading ? 'Creating...' : 'Create Profile'}
			</button>
		</div>
	</form>
</div>
