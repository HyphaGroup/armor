<script lang="ts">
	import { onMount } from 'svelte';
	import { api, type ProfileSummary } from '$lib/api';

	let profiles = $state<ProfileSummary[]>([]);
	let loading = $state(true);
	let error = $state('');

	onMount(async () => {
		try {
			profiles = await api.listProfiles();
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to load profiles';
		} finally {
			loading = false;
		}
	});

	async function deleteProfile(id: string, name: string) {
		if (!confirm(`Delete profile "${name}"?`)) return;
		try {
			await api.deleteProfile(id);
			profiles = profiles.filter(p => p.id !== id);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to delete profile';
		}
	}
</script>

<div class="space-y-6">
	<div class="flex justify-between items-center">
		<h1 class="text-2xl font-bold">Profiles</h1>
		<a
			href="/profiles/new"
			class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
		>
			New Profile
		</a>
	</div>

	{#if loading}
		<p class="text-gray-500">Loading...</p>
	{:else if error}
		<p class="text-red-500">{error}</p>
	{:else if profiles.length === 0}
		<div class="text-center py-12 bg-white rounded-lg shadow">
			<p class="text-gray-500 mb-4">No profiles yet</p>
			<a
				href="/profiles/new"
				class="text-blue-600 hover:text-blue-700"
			>
				Create your first profile
			</a>
		</div>
	{:else}
		<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
			{#each profiles as profile}
				<div class="bg-white rounded-lg shadow p-6 hover:shadow-md transition">
					<div class="flex justify-between items-start mb-4">
						<div>
							<h2 class="font-semibold text-lg">{profile.name}</h2>
							{#if profile.description}
								<p class="text-gray-500 text-sm mt-1">{profile.description}</p>
							{/if}
						</div>
						<button
							onclick={() => deleteProfile(profile.id, profile.name)}
							class="text-gray-400 hover:text-red-500 transition"
							title="Delete profile"
						>
							<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
								<path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
							</svg>
						</button>
					</div>
					
					<div class="mb-4">
						<div class="flex justify-between text-sm mb-1">
							<span class="text-gray-600">Completeness</span>
							<span class="font-medium">{Math.round(profile.completeness)}%</span>
						</div>
						<div class="w-full bg-gray-200 rounded-full h-2">
							<div
								class="bg-blue-600 h-2 rounded-full transition-all"
								style="width: {profile.completeness}%"
							></div>
						</div>
					</div>

					<a
						href="/profiles/{profile.id}"
						class="block text-center bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 transition"
					>
						Open
					</a>
				</div>
			{/each}
		</div>
	{/if}
</div>
